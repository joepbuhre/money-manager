package routers

import (
	"context"
	"log"
	"money-manager/database"
	m "money-manager/middlewares"
	u "money-manager/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(rg *gin.RouterGroup, env *u.Env) {
	public := rg.Group("/users")

	public.GET("/", m.IsAllowedTo(env, "users", m.CanRead), GetUsers(env))

	public.POST("/", m.IsAllowedTo(env, "users", m.CanCreate), CreateUser(env))
	public.POST("/login", LoginUser(env))

	public.DELETE("/:userid", m.IsAllowedTo(env, "users", m.CanDelete), DeleteUser(env))

}

// Create a user.
//
//	@ID			create-user
//	@Produce	json
//	@Router		/users [post]
func CreateUser(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var newUser database.CreateUserParams
		_, err = u.ReadRequestBody(c, &newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
			return
		}

		var user database.User
		user, err = env.Queries.CreateUser(c, newUser)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}

// Get all users.
//
//	@ID			get-user
//	@Produce	json
//	@Router		/users [get]
func GetUsers(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		var users, _ = env.Queries.ListUsers(context.Background())
		c.JSON(http.StatusOK, users)
	}
}

// Delete user
//
//	@ID		delete-user
//	@Router	/users/:userid [delete]
func DeleteUser(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q = env.Queries
		var idStr string
		var id int
		var err error
		var errors u.JsonErrors
		idStr = c.Param("userid")

		id, err = strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			errors.Add(u.ApiErrors[u.ErrInternal])
			c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
			return
		}

		// Delete user here
		err = q.DeleteUser(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			errors.Add(u.ApiErrors[u.ErrInternal])
			c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
			return
		}

		// Handle an accepted response
		c.Status(http.StatusAccepted)
	}
}

// Login a user.
//
//	@ID			Login-user
//	@Produce	json
//	@Router		/login [post]
func LoginUser(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var token string
		var errors u.JsonErrors

		// Check if token is still valid...
		// env.Queries.GetUserFromToken(context.Background() )

		token, err = u.GenerateSecureToken(40)
		if err != nil {
			errors.Add(u.ApiErrors[u.ErrForbidden])
			c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
			return
		}

		var returnToken database.UserToken
		returnToken, err = env.Queries.LoginUser(context.Background(), database.LoginUserParams{
			UserID: 1,
			Token:  token,
		})

		if err != nil {
			errors.Add(u.ApiErrors[u.ErrUsersSomethingWrong])
			c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
		}

		c.JSON(http.StatusOK, returnToken)

	}
}
