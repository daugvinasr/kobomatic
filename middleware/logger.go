package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func LogrusLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			c.Error(err)
		}

		switch {
		case c.Response().Status >= http.StatusInternalServerError:
			logrus.Error(fmt.Sprintf("method=%s, status=%d, uri=%s, error=%s",
				c.Request().Method, c.Response().Status, c.Request().RequestURI, err))
		case c.Response().Status == http.StatusNotFound:
			logrus.Warn(fmt.Sprintf("method=%s, status=%d, uri=%s",
				c.Request().Method, c.Response().Status, c.Request().RequestURI))
		default:
			logrus.Info(fmt.Sprintf("method=%s, status=%d, uri=%s",
				c.Request().Method, c.Response().Status, c.Request().RequestURI))
		}

		return err
	}
}
