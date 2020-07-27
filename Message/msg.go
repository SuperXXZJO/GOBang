package Message

import "encoding/json"

const (

	MSGTYPE_READY = "ready"  //准备
	MSGTYPE_CANCELREADY = "cancel_ready"  //取消准备
	MSGTYPE_POINT = "point"   //落子消息
	MSGTYPE_WINNER = "winner"  //胜利消息

	MSGFROM_SYS = "system" //系统消息
	
)

type Message struct {
	MessageType string		`json:"message_type"` //消息类型
	MessageFrom string		`json:"message_from"` //消息来自
	Data string		`json:"data"` //消息内容
}

//解析消息
func UnmarshalMsg(msg []byte) *Message {
	m :=&Message{}
	json.Unmarshal(msg,m)
	return m
}

//重新拼装消息类型
func NewMsg(data string,msgtype string,msgfrom string) *Message {
	m :=&Message{
		MessageType: msgtype,
		MessageFrom: msgfrom,
		Data:        data,
	}
	return m
}