package adapters

import (
	"github.com/alserov/hrs/gateway/internal/adapters/handlers"
	"github.com/alserov/hrs/gateway/internal/clients"
	"github.com/alserov/hrs/gateway/internal/log"
	"github.com/alserov/hrs/gateway/internal/middleware"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *handlers.Auth
	Chat *handlers.Chat
}

func NewController(clients *clients.Clients) *Controller {
	return &Controller{
		Chat: handlers.NewChat(clients.Comm),
	}
}

func Setup(r *echo.Echo, ctrl *Controller, log log.Logger) {
	r.Use(middleware.HandleError)
	r.Use(middleware.SetLogger)

	auth := r.Group("/auth")
	auth.POST("/signup", ctrl.Auth.Signup)
	auth.POST("/login", ctrl.Auth.Login)

	chat := r.Group("/chat")
	chat.Use(middleware.SetCookieToken)
	chat.GET("/history/:recipient_id", ctrl.Chat.History)
	chat.GET("/chats", ctrl.Chat.Chats)

	ws := r.Group("/ws")
	ws.Use(middleware.UpgradeWS)
	ws.Use(middleware.SetHeaderToken)
	ws.GET("/join", ctrl.Chat.Join)
}
