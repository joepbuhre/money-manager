package routers

import (
	"context"
	"money-manager/database"
	m "money-manager/middlewares"
	u "money-manager/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAccountRouter(rg *gin.RouterGroup, env *u.Env) {
	public := rg.Group("/accounts")

	public.POST("/", m.IsAllowedTo(env, "accounts", m.CanCreate), CreateAccount(env))

}

// Create a user.
//
//	@ID			create-user
//	@Produce	json
//	@Router		/users [post]
func CreateAccount(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var account database.Account
		var newAccount database.CreateAccountParams

		_, err = u.ReadRequestBody(c, &newAccount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"expecting another object": "error"})
			return
		}

		account, err = env.Queries.CreateAccount(context.Background(), newAccount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)

	}
}
