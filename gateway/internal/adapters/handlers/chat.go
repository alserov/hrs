package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/hrs/gateway/internal/clients/grpc"
	"github.com/alserov/hrs/gateway/internal/middleware"
	"github.com/alserov/hrs/gateway/internal/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"sync"
)

type Chat struct {
	cl grpc.CommunicationServiceClient

	conns sync.Map

	//convert utils.Converter
}

func NewChat(cl grpc.CommunicationServiceClient) *Chat {
	return &Chat{
		cl:    cl,
		conns: sync.Map{},
	}
}

func (ch *Chat) Join(c echo.Context) error {
	ws, ok := c.Get(string(middleware.CtxWebsocket)).(*websocket.Conn)
	if !ok {
		return nil
	}

	token := c.Get(string(middleware.CtxToken))

	ch.conns.Store(token, ws)
	defer ch.conns.Delete(token)

	for {
		code, msg, err := ws.ReadMessage()
		if code == websocket.CloseMessage || err != nil {
			return nil
		}

		var message models.Message
		if err = json.Unmarshal(msg, &message); err != nil {
			return nil
		}

		//_, err = ch.cl.CreateMessage(c.Request().Context(), ch.convert.ToMessage(message))
		//if err != nil {
		//	return nil
		//}

		fmt.Println(message)
		//err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		//if err != nil {
		//	c.Logger().Error(err)
		//}
	}
}

func (ch *Chat) History(c echo.Context) error {
	return nil
}

func (ch *Chat) Chats(c echo.Context) error {
	return nil
}
