package game

import (
	"GOBang/player"
	"GOBang/room"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type item struct {
	RoomName string `json:"room_name" binding:"required"`
	PlayerName string `json:"player_name" binding:"required"`
}



func Start(c *gin.Context)  {
	mod := &item{}
	if err:= c.BindJSON(mod);err!=nil{
		c.JSON(403,gin.H{
			"message":err.Error(),
		})
		return
	}

	//查询房间是否存在
	res,err:=room.FindRoomByRoomName(mod.RoomName)
	if err != nil {
		c.JSON(410,gin.H{
			"message":"房间不存在",
			"data":err,
		})
		return
	}


	//新建房间Hub
	hub := room.NewRoomHub()

	hub.Model.RoomName = mod.RoomName

	//ws升级
	upgrader := websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws,err := upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		c.JSON(500,gin.H{
			"message":"升级失败",
			"data":err,
		})
		return
	}

	//新建Player
	p :=player.NewPlayer(ws,mod.PlayerName)

	//判断是否是房主
	if res.RoomOwner == mod.PlayerName {
		hub.Model.RoomOwner=mod.PlayerName
		hub.Palyers[mod.PlayerName] = p
	}else {
		hub.Palyers[mod.PlayerName] = p
	}

	//TODO 一个go 写消息
	//TODO 一个 go room监听
	//TODO  开始游戏

}