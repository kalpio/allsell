package main

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/handlers"
	"github.com/kalpio/allsell/src/migrations"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"strings"
)

func main() {
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	app.Use(Authorize())

	dsn := "file:./tmp/data.db?cache=shared&mode=rwc"
	migrationConfig := migrations.MigrationConfig{
		Driver:           "sqlite",
		ConnectionString: dsn,
	}
	if err := migrations.Migrate(context.Background(), migrationConfig); err != nil {
		log.Fatal(err)
	}

	db := sqlx.MustOpen("sqlite", dsn)

	userHandler := handlers.NewUserHandler(db)
	homeHandler := handlers.HomeHandler{}

	app.GET("/user/login", userHandler.LoginGet)
	app.POST("/user/login", userHandler.LoginPost)
	app.GET("/user/register", userHandler.RegisterGet)
	app.POST("/user/register", userHandler.RegisterPost)
	app.GET("/", homeHandler.IndexGet)

	app.Logger.Fatal(app.Start(":1234"))
}

type AuthorizationConfig struct {
	Skipper middleware.Skipper
}

var DefaultAuthorizationConfig = AuthorizationConfig{
	Skipper: middleware.DefaultSkipper,
}

func Authorize() echo.MiddlewareFunc {
	return AuthorizeWithConfig(DefaultAuthorizationConfig)
}

func AuthorizeWithConfig(config AuthorizationConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultAuthorizationConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			url := req.URL
			path := url.Path

			if strings.HasPrefix(path, "/user") {
				return next(c)
			}

			var login = func(c echo.Context) error {
				return c.Redirect(http.StatusMovedPermanently, "/user/login")
			}
			sess, err := session.Get("session", c)
			if err != nil {
				return login(c)
			}
			value, ok := sess.Values["username"]
			if !ok {
				return login(c)
			}
			v := value.(string)
			if len(v) == 0 {
				return login(c)
			}
			return next(c)
		}
	}
}
