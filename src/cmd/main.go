package main

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/handlers"
	"github.com/kalpio/allsell/src/middleware"
	"github.com/kalpio/allsell/src/migrations"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	app.Use(middleware.Authorize())

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
	app.GET("/user/logout", userHandler.LogoutGet)
	app.GET("/user/change-password", userHandler.ChangePasswordGet)
	app.POST("/user/change-password", userHandler.ChangePasswordPost)

	app.GET("/", homeHandler.IndexGet)

	app.Logger.Fatal(app.Start(":1234"))
}
