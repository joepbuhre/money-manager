package routers

import (
	"context"
	"log"
	"money-manager/database"
	"money-manager/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pjebs/jsonerror"
)

func AddUserRouter(rg *gin.RouterGroup, env *utils.Env) {
	public := rg.Group("/users")

	public.GET("/", GetUser(env))

	public.POST("/", CreateUser(env))
	public.POST("/login", LoginUser(env))

}

// Create a request to upload a specific file.
//
//	@ID			create-user
//	@Produce	json
//	@Router		/users [get]
func GetUser(env *utils.Env) gin.HandlerFunc {

	return func(c *gin.Context) {
		var users, _ = env.Queries.ListUsers(context.Background())
		log.Println(users)

		c.JSON(http.StatusOK, users)
	}
}

// Create a user.
//
//	@ID			create-user
//	@Produce	json
//	@Router		/users [post]
func CreateUser(env *utils.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var newUser database.CreateUserParams
		_, err = utils.ReadRequestBody(c, &newUser)
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

// Login a user.
//
//	@ID			Login-user
//	@Produce	json
//	@Router		/login [post]
func LoginUser(env *utils.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var token string
		var errors utils.JsonErrors

		// Check if token is still valid...
		// env.Queries.GetUserFromToken(context.Background() )

		token, err = utils.GenerateSecureToken(40)
		if err != nil {
			errors.Add(jsonerror.New(1, "Something went wrong with token", ""))
			c.AbortWithStatusJSON(http.StatusBadRequest, errors.ToString())
			return
		}

		var returnToken database.UserToken
		returnToken, err = env.Queries.LoginUser(context.Background(), database.LoginUserParams{
			UserID: 1,
			Token:  token,
		})

		if err != nil {
			errors.Add(jsonerror.New(1, "Something went wrong with logging in user", ""))
			c.AbortWithStatusJSON(http.StatusBadRequest, errors.ToString())
		}

		c.JSON(http.StatusOK, returnToken)

	}
}

// // Create a request to login a user
// //
// //	@ID			login-user
// //	@Produce	json
// //	@Param		string	body		database.User	false	"string valid"
// //	@Success	200		{object}	database.UserToken
// //	@Router		/users/login [post]
// func loginUser(db *sql.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var err error

// 		defer func() {
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			} else if len(c.Errors) > 0 {
// 				c.JSON(http.StatusBadRequest, c.Errors[0])
// 			}
// 		}()

// 		var user database.User
// 		_, err = routers.ReadRequestBody(c, &user)
// 		if err != nil {
// 			c.Error(err)
// 			return
// 		}

// 		var res string
// 		res, err = database.LoginUser(user, db)
// 		if err != nil {
// 			c.Error(err)
// 			return
// 		}

// 		c.JSON(http.StatusOK, database.UserToken{Token: res})

// 	}
// }
