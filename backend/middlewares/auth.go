package middlewares

import (
	"context"
	"database/sql"
	"log"
	"money-manager/database"
	u "money-manager/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthMethodEnum int

const (
	// Define the error types
	StaticAuthEnum AuthMethodEnum = iota
	SessionAuthEnum
	ReverseProxyAuthEnum
)

// AuthMethod is an interface that both SessionAuth and HeaderAuth will implement
type AuthMethod interface {
	AuthType() AuthMethodEnum
}

// Implement the AuthMethod interface for StaticAuth
type StaticAuth struct {
	headerkey string
	value     string
}

func (s StaticAuth) AuthType() AuthMethodEnum {
	return StaticAuthEnum
}

// Implement reverse proxy auth
type ReverseProxyAuth struct {
	env               *u.Env
	usernameHeaderKey string
	emailHeaderKey    string
	nameHeaderKey     string
}

func (s ReverseProxyAuth) AuthType() AuthMethodEnum {
	return ReverseProxyAuthEnum
}

// Implement the AuthMethod interface for SessionAuth
type SessionAuth struct {
	db        *sql.DB
	query     string
	headerkey string
}

func (s SessionAuth) AuthType() AuthMethodEnum {
	return SessionAuthEnum
}

type AuthMiddleware struct {
	DB          *sql.DB
	errors      u.JsonErrors
	AuthMethods []AuthMethod
}

func (am *AuthMiddleware) AddAuth(method AuthMethod) {
	am.AuthMethods = append(am.AuthMethods, method)
}

func (am *AuthMiddleware) handleStatic(c *gin.Context, auth StaticAuth) bool {
	var headerValue = c.GetHeader(auth.headerkey)
	log.Println("Handle static auth")

	// Only if fetched header value, but only if headervalue has been provided
	if headerValue == auth.value && headerValue != "" {
		return true
	}
	if headerValue != "" {
		return false
	}

	// We are not allowed
	return false
}

func (am *AuthMiddleware) handleSession(c *gin.Context, auth SessionAuth) bool {
	var preparedStatement string = auth.query
	var db = auth.db
	stmt, err := db.Prepare(preparedStatement)
	if err != nil {
		log.Printf("RestGate: Run time database error: %+v", err)
		am.errors.Add(u.ApiErrors[u.ErrAuthNotSetup]) //"Software Developers have not setup authentication correctly"
		return false
	}
	defer stmt.Close()

	var key = c.GetHeader(auth.headerkey)
	key = strings.Replace(key, "Bearer ", "", 1)

	// If key has not been setup exiting session auth
	if key == "" {
		log.Println("Exiting session auth, because session auth has not been setup")
		return false
	}

	// Now we are going to fetch the userid based on supplied token
	var id int

	err = stmt.QueryRow(key).Scan(&id)

	// We also set the userid in the context of gin
	if err == nil && id > 0 {
		//Authentication PASSED
		log.Println("Setting userid to ", id)
		c.Set("UserId", id)
		return true
	} else { //==sql.ErrNoRows or count == 0:: At this point authenticaton has been supplied but no valid rows has been given
		return false
	}

}

func (am *AuthMiddleware) handleReverseProxy(c *gin.Context, auth ReverseProxyAuth) bool {
	var q = auth.env.Queries
	var headerValue = c.GetHeader(auth.usernameHeaderKey)

	if headerValue == "" {
		return false
	}

	var user, err = q.GetUserByUsername(context.Background(), headerValue)

	if err == pgx.ErrNoRows {
		var nameHeader = c.GetHeader(auth.nameHeaderKey)
		var emailHeader = c.GetHeader(auth.emailHeaderKey)

		// Create user because of reverse proxy
		user, err = q.CreateUser(context.Background(), database.CreateUserParams{
			Username: headerValue,
			Name:     nameHeader,
			Email:    pgtype.Text{String: emailHeader},
		})

		if err != nil {
			// TODO
			return false
		}

		c.Set("UserId", int(user.ID))
		return true

	} else if err != nil {
		log.Println(err)
		return false
	} else if int(user.ID) > 0 {
		c.Set("UserId", int(user.ID))
		return true
	}

	return false

}

func (am *AuthMiddleware) ServeHTTP(c *gin.Context, next http.HandlerFunc) {
	// c.Writer, c.Request
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	if am == nil {
		c.AbortWithStatusJSON(500, "Authmiddleware has not been set")
		return
	}

	defer func() {
		if am.errors.HasErrors() {
			c.AbortWithStatusJSON(am.errors.ErrorCode, am.errors.ToString())
			return
		} else {
			next(w, req)
		}
	}()

	//Check key in Header
	// key := req.Header.Get(am.headerKeyLabel)
	if am.AuthMethods == nil {
		am.errors.Add(u.ApiErrors[u.ErrAuthNotSetup])
		return
	}
	var allowed bool = false

	for i := 0; i < len(am.AuthMethods); i++ {
		var auth AuthMethod = am.AuthMethods[i]

		if auth.AuthType() == ReverseProxyAuthEnum {
			log.Println("Handling reverse proxy auth")
			allowed = am.handleReverseProxy(c, auth.(ReverseProxyAuth))
		} else if auth.AuthType() == SessionAuthEnum {
			log.Println("Implemented sessionauth")
			allowed = am.handleSession(c, auth.(SessionAuth))
		} else if auth.AuthType() == StaticAuthEnum {
			log.Println("Implemented StaticAuth")
			allowed = am.handleStatic(c, auth.(StaticAuth))
		}

		// if an allowed auth has been provided exit here
		if allowed {
			next(w, req)
			break
		}
	}

	// Apparently we have no valid authentication here, so we are exiting, but only if no auth has been failed yet
	if !allowed && !am.errors.HasErrors() {
		log.Println("No valid auth has been provided,")
		am.errors.Add(u.ApiErrors[u.ErrAuthNotSupplied])
	}

}
