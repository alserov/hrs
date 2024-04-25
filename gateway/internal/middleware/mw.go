package middleware

import (
	"fmt"
	"github.com/alserov/hrs/gateway/internal/log"
	"github.com/alserov/hrs/gateway/internal/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleError(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := fn(c); err != nil {
			st, msg := utils.FromError(err)
			_ = c.JSON(st, echo.Map{"error": msg})
		}

		return nil
	}
}

var (
	upgrader = websocket.Upgrader{}

	CtxWebsocket CtxValue = "websocket"
	CtxLogger    CtxValue = "logger"
	CtxToken     CtxValue = "token"
)

type (
	CtxValue string
)

func UpgradeWS(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer func() {
			_ = ws.Close()
		}()

		c.Set(string(CtxWebsocket), ws)

		return fn(c)
	}
}

func SetLogger(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(string(CtxLogger), log.GetLogger())
		return fn(c)
	}
}

func SetCookieToken(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie(string(CtxToken))
		if err != nil {
			fmt.Println(token)
			_ = c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		c.Set(string(CtxToken), token)
		return fn(c)
	}
}

func SetHeaderToken(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(string(CtxToken))
		if token == "" {
			_ = c.JSON(http.StatusBadRequest, echo.Map{"error": "token not provided"})
		}

		c.Set(string(CtxToken), token)
		return fn(c)
	}
}
