package handlers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/models"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/types"
	"github.com/kalpio/allsell/src/views/user"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	db          *sqlx.DB
	userService services.UserService
}

func NewUserHandler(db *sqlx.DB) UserHandler {
	return UserHandler{db, services.NewUserService(db)}
}

func (u UserHandler) LoginGet(c echo.Context) error {
	return render(c, user.Login())
}

func (u UserHandler) LoginPost(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "admin" && password == "admin" {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		sess.Values["username"] = username
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (u UserHandler) RegisterGet(c echo.Context) error {
	return render(c, user.Register())
}

func (u UserHandler) RegisterPost(c echo.Context) error {
	ur := types.UserRegister{}
	if err := c.Bind(&ur); err != nil {
		return c.HTML(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	if err := ur.Validate(); err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	usr := models.NewUser(ur.UserName, ur.Email, ur.Password)
	if err := u.userService.Register(c.Request().Context(), usr); err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, usr.Name)
}
