package handlers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/types/user"
	views "github.com/kalpio/allsell/src/views/user"
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
	return render(c, views.Login())
}

func (u UserHandler) LoginPost(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, usr := u.userService.Login(c.Request().Context(), username, password)
	if result.Success() {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			Secure:   false,
		}

		sess.Values["username"] = usr.Unwrap().Name
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (u UserHandler) LogoutGet(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (u UserHandler) RegisterGet(c echo.Context) error {
	return render(c, views.Register())
}

func (u UserHandler) RegisterPost(c echo.Context) error {
	ur := user.Register{}
	if err := c.Bind(&ur); err != nil {
		return c.HTML(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	if err := ur.Validate(); err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	usr := user.NewUser(ur.UserName, ur.Email, ur.Password)
	if err := u.userService.Register(c.Request().Context(), usr); err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, usr.Name)
}
