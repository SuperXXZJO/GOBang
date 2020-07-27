package player

import (
	"GOBang/Message"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type Player struct {

	sync.Mutex

	PlayerName string
	Conn *websocket.Conn
	OutChan chan []byte
	IsReady bool



}

//新建player
func NewPlayer(conn *websocket.Conn,playername string) *Player {
	p :=&Player{
		PlayerName: playername,
		Conn:       conn,
		OutChan:    make(chan []byte),
		IsReady:    false,
	}
	return p

}


//写消息
func (p *Player) ReadMsg () error {

	p.Lock()
	defer p.Unlock()

	_,msg,err:=p.Conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("写消息时出现的错误：%s",err.Error())
	}

	p.OutChan <- msg
	return nil
}


//切换准备状态
func (p *Player) GetReady(data string)  {
	if data == "ready"{
		p.IsReady = true
	}else if data == "cancel_ready" {
		p.IsReady = false
	}

}

//解析消息
func (p *Player) UnmarshalMsg(msg []byte)  {

	var res  = & Message.Message{}
	json.Unmarshal(msg,res)

	switch res.MessageType {
	case Message.MSGTYPE_READY:
		p.GetReady(res.Data)
	case Message.MSGTYPE_CANCELREADY:
		p.GetReady(res.Data)
	case Message.MSGTYPE_POINT:
		//TODO
	case Message.MSGTYPE_WINNER:
		//TODO
	}

}

//读消息
func (p *Player)SendMsg() []byte {
	msg :=<-p.OutChan
	return msg
}