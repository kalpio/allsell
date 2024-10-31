package handlers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/middleware"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/session"
	"github.com/kalpio/allsell/src/types/user"
	views "github.com/kalpio/allsell/src/views/user"
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
		if err := session.Set(c,
			middleware.DefaultAuthorizationConfig.SessionKey,
			usr.Unwrap().Name,
			session.DefaultSessionOptions); err != nil {
			return err
		}
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (u UserHandler) LogoutGet(c echo.Context) error {
	if err := session.Delete(c); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (u UserHandler) RegisterGet(c echo.Context) error {
	return render(c, views.Register())
}

func (u UserHandler) RegisterPost(c echo.Context) error {
	register := user.Register{}
	if err := c.Bind(&register); err != nil {
		return c.HTML(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	if err := register.Validate(); err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	usr := user.NewUser(register.UserName, register.Email, register.Password)
	if err := u.userService.Register(c.Request().Context(), usr); err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, usr.Name)
}

func (u UserHandler) ChangePasswordGet(c echo.Context) error {
	value, _ := session.Get[string](c, middleware.DefaultAuthorizationConfig.SessionKey)
	usrValue, err := u.userService.Get(c.Request().Context(), value.UnwrapOr(""))
	if err != nil || usrValue.IsNone() {
		return err
	}
	usr := usrValue.Unwrap()

	return render(c, views.ChangePassword(usr.Name, usr.Email))
}

func (u UserHandler) ChangePasswordPost(c echo.Context) error {
	value, _ := session.Get[string](c, middleware.DefaultAuthorizationConfig.SessionKey)
	changePassword := user.ChangePassword{}
	if err := c.Bind(&changePassword); err != nil {
		return c.HTML(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}
	if err := changePassword.Validate(); err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	return u.userService.ChangePassword(c.Request().Context(), value.UnwrapOr(""), changePassword.CurrentPassword, changePassword.NewPassword)
}
