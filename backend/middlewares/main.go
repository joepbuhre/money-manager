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

	"context"
	"crypto/subtle"
	"database/sql"
	"log"
	"money-manager/database"
	"money-manager/utils"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/stdlib"

	"github.com/gin-gonic/gin"
	e "github.com/pjebs/jsonerror"
)

type AuthenticationSource int

const (
	Static   AuthenticationSource = 0
	Database AuthenticationSource = 1
)

// When AuthenticationSource=Static, Key(s)=Actual Key and Secret(s)=Actual Secret.
// When AuthenticationSource=Database, Key[0]=Key_Column and Secret[0]=Secret_Column.
type Config struct {
	*sql.DB
	Key           []string
	Secret        []string //Can be "" but not recommended
	TableName     string
	ErrorMessages map[int]e.JE
	Context       func(r *http.Request, authenticatedKey string)
	Debug         bool
	Postgres      bool
	Logger        ALogger
	QueryString   string
	Env           *utils.Env
}

type RESTGate struct {
	headerKeyLabel    string
	headerSecretLabel string
	source            AuthenticationSource
	config            Config
}

type ALogger interface {
	Printf(format string, v ...interface{})
}

func New(headerKeyLabel string, headerSecretLabel string, as AuthenticationSource, config Config) *RESTGate {
	if config.Logger == nil {
		config.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	t := &RESTGate{headerKeyLabel: headerKeyLabel, headerSecretLabel: headerSecretLabel, source: as, config: config}
	t.config.Logger.Printf("RestGate initializing")

	numberKeys := len(t.config.Key)
	numberSecrets := len(t.config.Secret)

	if numberKeys == 0 { //Key is not set
		if t.config.Debug == true {
			t.config.Logger.Printf("RestGate: Key is not set")
		}
		return nil
	}

	if numberSecrets > numberKeys { //Too many Secret's defined
		if t.config.Debug == true {
			t.config.Logger.Printf("RestGate: Too many Secrets defined. At most there should be 1 secret per key")
		}
		return nil
	}

	if headerKeyLabel == "" { //headerKeyLabel must be defined
		if t.config.Debug == true {
			t.config.Logger.Printf("RestGate: headerKeyLabel is not defined.")
		}
		return nil
	}

	//Default Error Messages
	if t.config.ErrorMessages == nil {
		t.config.ErrorMessages = map[int]e.JE{
			1:  e.New(1, "No Key Or Secret", "", "com.github.pjebs.restgate"),
			2:  e.New(2, "Unauthorized Access", "", "com.github.pjebs.restgate"),
			3:  e.New(3, "Please use HTTPS connection", "", "com.github.pjebs.restgate"),
			99: e.New(99, "Software Developers have not setup authentication correctly", "", "com.github.pjebs.restgate"),
		}

	} else {
		if _, ok := t.config.ErrorMessages[1]; !ok {
			t.config.ErrorMessages[1] = e.New(1, "No Key Or Secret", "", "com.github.pjebs.restgate")
		}

		if _, ok := t.config.ErrorMessages[2]; !ok {
			t.config.ErrorMessages[2] = e.New(2, "Unauthorized Access", "", "com.github.pjebs.restgate")
		}

		if _, ok := t.config.ErrorMessages[3]; !ok {
			t.config.ErrorMessages[3] = e.New(3, "Please use HTTPS connection", "", "com.github.pjebs.restgate")
		}

		if _, ok := t.config.ErrorMessages[99]; !ok {
			t.config.ErrorMessages[99] = e.New(99, "Software Developers have not setup authentication correctly", "", "com.github.pjebs.restgate")
		}
	}

	if as == Database {

		if numberKeys != 1 { //We need exactly 1 Key (it represents field name in database)
			if t.config.Debug == true {
				t.config.Logger.Printf("RestGate: For Database mode, we need exactly 1 Key which represents the field name in the database table")
			}
			return nil
		}

		//Check if database is set.
		//The developer should ensure a database has been selected (i.e. to prevent "No Database selected" error)
		if t.config.DB == nil { //DB is not set
			if t.config.Debug == true {
				t.config.Logger.Printf("RestGate: Database is not set. Be sure that a database name is selected")
			}
			return nil
		}

		//Check if table is set
		if t.config.TableName == "" { //Table name is not set
			if t.config.Debug == true {
				t.config.Logger.Printf("RestGate: For Database mode, a table name is required")
			}
			return nil
		}

	}

	return t
}

func (self *RESTGate) ServeHTTP(c *gin.Context, next http.HandlerFunc) {
	// c.Writer, c.Request
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request

	//Check key in Header
	key := req.Header.Get(self.headerKeyLabel)
	secret := req.Header.Get(self.headerSecretLabel)
	var errors utils.JsonErrors

	defer func() {
		// We expect to have a user here
		var user database.User
		var err error
		user, err = self.config.Env.Queries.GetUserFromToken(context.Background(), key)
		log.Println(errors)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ToString())
			return
		} else {
			c.Set("currentUser", user)
			next(w, req)
		}

	}()

	if key == "" {
		//Authentication Information not included in request
		errors.Add(self.config.ErrorMessages[1]) //"No Key Or Secret"
		return
	}

	if self.source == Static {

		secretDoesntExist := len(self.config.Secret) == 0
		authenticationPassed := false

		//First search through all keys
		for index, element := range self.config.Key {
			if secureCompare(key, element) { //Key matches

				//Now check if secret matches
				if secretDoesntExist {
					//Authentication PASSED
					authenticationPassed = true
					break
				} else if index > (len(self.config.Secret) - 1) { //Out of Range so corresponding secret doesn't exist
					//Authentication PASSED
					authenticationPassed = true
					break
				} else {
					//Corresponding Secret exists
					if secureCompare(secret, self.config.Secret[index]) {
						//Authentication PASSED
						authenticationPassed = true
						break
					} else {
						//Authentication FAILED
						authenticationPassed = false
						break
					}
				}
			}
		}

		//Authentication FAILED - No Key's matched
		if authenticationPassed == false {
			errors.Add(self.config.ErrorMessages[2]) //"Unauthorized Access"
			return
		} else { //Authentication PASSED
			if self.config.Context != nil {
				self.config.Context(req, key)
			}
			return
		}

	} else if self.source == Database {
		db := self.config.DB

		secretDoesntExists := len(self.config.Secret) == 0 || self.config.Secret[0] == ""

		var preparedStatement string = self.config.QueryString

		stmt, err := db.Prepare(preparedStatement)
		if err != nil {
			if self.config.Debug {
				self.config.Logger.Printf("RestGate: Run time database error: %+v", err)
			}
			errors.Add(self.config.ErrorMessages[99]) //"Software Developers have not setup authentication correctly"
			return
		}
		defer stmt.Close()

		var count int //stores query result
		if secretDoesntExists {
			err = stmt.QueryRow(key).Scan(&count)
		} else {
			err = stmt.QueryRow(key, secret).Scan(&count)
		}

		if err == nil && count == 1 {
			//Authentication PASSED
			if self.config.Context != nil {
				self.config.Context(req, key)
			}
			return
		} else { //==sql.ErrNoRows or count == 0
			//Something went wrong
			if self.config.Debug == true && count > 1 {
				self.config.Logger.Printf("RestGate: Database query returned more than 1 identical Key. Make sure the KEY field in the table is set to UNIQUE")
			}
			errors.Add(self.config.ErrorMessages[2]) //"Unauthorized Access"
			return
		}

	} else {

		errors.Add(self.config.ErrorMessages[99]) //"Software Developers have not setup authentication correctly"
		return
	}

}

// secureCompare performs a constant time compare of two strings to limit timing attacks.
func secureCompare(given string, actual string) bool {
	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		return subtle.ConstantTimeCompare([]byte(given), []byte(actual)) == 1
	} else {
		/* Securely compare actual to itself to keep constant time, but always return false */
		return subtle.ConstantTimeCompare([]byte(actual), []byte(actual)) == 1 && false
	}
}

func GetAuthMiddleware(env *utils.Env) func(c *gin.Context) {
	// Initialize Restgate
	var rg = New("api-key", "", Database, Config{
		Key:         []string{"api-key"},
		Secret:      []string{""},
		DB:          stdlib.OpenDBFromPool(env.DB),
		Debug:       true,
		TableName:   "user_sessions",
		Postgres:    true,
		QueryString: "select count(*) from user_sessions where token = $1",
		Env:         env,
	})

	// Create Gin middleware - integrate Restgate with Gin
	return func(c *gin.Context) {
		nextAdapter := func(http.ResponseWriter, *http.Request) {
			c.Next()
		}
		rg.ServeHTTP(c, nextAdapter)
	}

}
