package room

import (
	"GOBang/player"
	"github.com/gorilla/websocket"
)
type RoomHub struct {
	Model RoomModel
	Palyers map[string]*player.Player
	SendChan chan []byte
}


//新建一个Hub
func NewRoomHub() *RoomHub {
	r := &RoomHub{}
	return r
}


//发送消息
func (r *RoomHub) WriteMsg()  {
	for {
		select {
		case msg := <-r.SendChan:
			for _,v :=range r.Palyers{
				v.Conn.WriteMessage(websocket.TextMessage,msg)
			}
		}
	}
}

//查看player是否都准备好了
func (r *RoomHub) CheckPlayerISReady() bool {
	count :=0
	for _,v :=range r.Palyers{

		if  v.IsReady == true {
			count ++
		}
	}
	if count == 2 {
		return true
	}else {
		return  false
	}
}

//