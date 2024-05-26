package routers

// import (
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"log"
// 	"money-manager/database"

// 	"github.com/gin-gonic/gin"
// )

// func ReadRequestBody(c *gin.Context, targetInterface interface{}) (interface{}, error) {
// 	var bd []byte

// 	// Read the request body
// 	bd, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		return nil, errors.New("failed to read request body")
// 	}
// 	defer c.Request.Body.Close()

// 	// Parse JSON data into SnappicFile struct
// 	err = json.Unmarshal(bd, &targetInterface)
// 	if err != nil {
// 		log.Println((err))
// 		return nil, errors.New("failed to parse JSON data")
// 	}

// 	return targetInterface, nil

// }

// func GetUserFromContext(c *gin.Context) (database.User, error) {
// 	var userRaw any
// 	var user database.User
// 	var exists bool
// 	userRaw, exists = c.Get("currentUser")

// 	if !exists {
// 		// Handle the case where the object doesn't exist
// 		return user, errors.New("user not found in context")
// 	}

// 	user, exists = userRaw.(database.User)
// 	if !exists {
// 		return user, errors.New("user not found in context")
// 	}

// 	return user, nil
// }
