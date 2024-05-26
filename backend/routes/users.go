package routers

// func AddUserRouter(rg *gin.RouterGroup, db *sql.DB) {
// 	usersPublic := rg.Group("/users")
// 	usersPrivate := rg.Group("/users")

// 	usersPrivate.Use(mw.GetAuthMiddleware(db))
// 	usersPublic.POST("/create", createUser(db))
// 	usersPublic.POST("/login", loginUser(db))

// usersPrivate.GET("/me", func(c *gin.Context) {
// 	var user, err = GetUserFromContext(c)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// })
// }

// Create a request to upload a specific file.
//
//	@ID			create-user
//	@Produce	json
//	@Success	200	{object}	database.User
//	@Router		/create-user [post]
// func createUser(db *sql.DB) gin.HandlerFunc {
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
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
// 			return
// 		}

// 		// Small validation on username password
// 		if user.Username == "" || user.Password == "" {
// 			c.Error(errors.New("either username or password is empty"))
// 			return
// 		}

// 		// var stmt *sql.Stmt
// 		// stmt, err = db.Prepare("insert into users (username, password) values ($1, $2)")
// 		// if err != nil {
// 		// 	c.Error(err)
// 		// 	return
// 		// }

// 		user, err = database.CreateUser(user, db)
// 		if err != nil {
// 			c.Error(err)
// 			return
// 		}

// 		c.JSON(http.StatusOK, user)

// 	}

// }

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
