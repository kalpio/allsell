package session

import (
	"github.com/gorilla/sessions"
	"github.com/kalpio/option"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const sessionKey = "session"

var DefaultSessionOptions = &sessions.Options{
	Path:     "/",
	MaxAge:   86400 * 7, // 7 days
	HttpOnly: true,
	Secure:   false,
}

func Get[T any](c echo.Context, key string) option.Option[T] {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return option.None[T](err)
	}

	value, ok := sess.Values[key]
	if !ok {
		return option.None[T](err)
	}

	v := value.(T)

	return option.Some(v)
}

func Set(c echo.Context, key string, value any, options *sessions.Options) error {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return err
	}

	sess.Options = options
	sess.Values[key] = value

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}

func Delete(c echo.Context) error {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}
