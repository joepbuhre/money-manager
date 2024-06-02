package middlewares

/*
|--------------------------------------------------------------------------
| WARNING
|--------------------------------------------------------------------------
| Never Set HTTPSProtectionOff=true In Production.
| The Key and Password will be exposed and highly unsecure otherwise!
| The database server should also use HTTPS Connection and be hidden away
|
*/

/*
Thanks to Ido Ben-Natan ("IdoBn") for postgres fix.
Thanks to Jeremy Saenz & Brendon Murphy for timing-attack protection
*/

import (
	// "errors"

	"money-manager/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/stdlib"
)

func GetAuthMiddleware(env *utils.Env) func(c *gin.Context) {
	// Create Gin middleware - integrate Restgate with Gin
	return func(c *gin.Context) {
		var am *AuthMiddleware = &AuthMiddleware{}

		am.AddAuth(ReverseProxyAuth{
			env:               env,
			usernameHeaderKey: utils.GetConfig("PROXY_USERNAME"),
			emailHeaderKey:    utils.GetConfig("PROXY_EMAIL"),
			nameHeaderKey:     utils.GetConfig("PROXY_NAME"),
		})

		am.AddAuth(StaticAuth{
			headerkey: utils.GetConfig("HEADER_API_KEY"),
			value:     utils.GetConfig("MASTER_API_KEY"),
		})

		am.AddAuth(SessionAuth{
			db:        stdlib.OpenDBFromPool(env.DB),
			query:     "select user_id from user_token where token = $1",
			headerkey: "Authorization",
		})

		nextAdapter := func(http.ResponseWriter, *http.Request) {
			c.Next()
		}
		am.ServeHTTP(c, nextAdapter)
	}

}
