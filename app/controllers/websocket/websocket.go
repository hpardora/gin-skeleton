package websocket

import (
	"encoding/json"
	"gin-skeleton/app/websocket"
	"gin-skeleton/helper/auth"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

var hub = websocket.NewHub()

func init() {
	go hub.Run()
}

func Websocket(c *gin.Context) {
	userId := auth.GetUserId(c)

	websocket.ServeWs(hub, c.Writer, c.Request, uint(userId))
}

func SendByUserID(userID int, data Data) {
	jsonData, _ := json.Marshal(data)
	hub.Clients[uint(userID)].Send <- []byte(jsonData)
}
