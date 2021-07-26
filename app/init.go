package app

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/revel/modules"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"r_res/app/controllers"
	"r_res/app/models"
	"time"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	revel.OnAppStart(SetupDatabaseConnection)
	revel.OnAppStop(CloseDatabaseConnection)
	revel.OnAppStop(LoadJWTVariables)

	revel.InterceptFunc(checkUser, revel.BEFORE, &controllers.User{})
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

func SetupDatabaseConnection() {

	host := revel.Config.StringDefault("mongo.host", "localhost")
	port := revel.Config.StringDefault("mongo.port", "27017")
	models.Database = revel.Config.StringDefault("mongo.database", "test")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + host + ":" + port))

	if err != nil {
		panic(err)
	}

	models.DB = client
}


func CloseDatabaseConnection() {
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := models.DB.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

type TokenResponse struct {
	Logged bool		`json:"logged"`
}
func checkUser(c *revel.Controller) revel.Result {

	if !isLoggedIn(c) {
		log.Println("USER IS NOT LOGGED IN!!!")
		r := TokenResponse{
			Logged: false,
		}
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(r)
	}
	log.Println("USER IS LOGGED IN!!!")
	return nil
}

func LoadJWTVariables() {
	controllers.JwtKey = []byte(revel.Config.StringDefault("jwt.key", "loodloo"))
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func isLoggedIn(c *revel.Controller) bool {
	token := c.Request.Header.Get("x-token")
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return controllers.JwtKey, nil
	})
	if err != nil {
		return false
	}
	if !tkn.Valid {
		return false
	}
	return true
}