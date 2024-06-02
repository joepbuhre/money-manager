package middlewares

import (
	"context"
	"log"
	"money-manager/database"
	u "money-manager/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/pjebs/jsonerror"
)

type PermissionType int

const (
	// Define the error types
	CanRead PermissionType = iota
	CanCreate
	CanUpdate
	CanDelete
	IsSuperadmin
)

func IsAllowedTo(env *u.Env, table string, permtype PermissionType) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors u.JsonErrors

		defer func() {
			if errors.HasErrors() {
				c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
			}
		}()

		var uidStr, ok = c.Get("UserId")
		var uid int
		uid, ok = uidStr.(int)
		// Exit when no userid has been supplied
		if !ok || uid == 0 {
			log.Println("No valid user has been supplied")
			errors.Add(u.ApiErrors[u.ErrForbidden])
			return
		}

		var permission database.GetUserPermissionRow
		var err error

		// Fetch required permissions
		permission, err = env.Queries.GetUserPermission(
			context.Background(),
			database.GetUserPermissionParams{
				TableName: "users",
				ID:        int32(uid),
			},
		)
		if err != nil && err != pgx.ErrNoRows {
			log.Println("DB error:", err)
			errors.Add(u.ApiErrors[u.ErrInternal])
		}
		// Check if is superadmin
		if permission.IsSuperadmin.Bool {
			log.Println("User is superadmin, skipping additional auth checks")
			c.Set("ObjectId", 0)
			return
		}

		// Setting obj
		c.Set("ObjectId", permission.ObjectID.Int32)

		// Check create permissions
		if permtype == CanCreate {
			if permission.CanCreate.Bool {

				return
			} else {
				errors.AddWoError(jsonerror.New(http.StatusUnauthorized, "Unauh", "not allowed to CanCreate"))
				return
			}
		}

		// Check read permissions
		if permtype == CanRead {
			if permission.CanRead.Bool {

				return
			} else {
				errors.AddWoError(jsonerror.New(http.StatusUnauthorized, "Unauh", "not allowed to CanRead"))
				return
			}
		}

		// Check Delete permissions
		if permtype == CanDelete {
			if permission.CanDelete.Bool {

				return
			} else {
				errors.AddWoError(jsonerror.New(http.StatusUnauthorized, "Unauh", "not allowed to CanDelete"))
				return
			}
		}

		// If we're still here exit with forbidden
		errors.Add(u.ApiErrors[u.ErrForbidden])

	}
}
