package room

import (
	"errors"
	"fmt"
)

//判断用户是否已经创建过房间
func CheckRoomOwner(mod *RoomBind) error {

	_,err:=FindRoomByRoomName(mod.RoomOwnerName)
	if err != nil {
		return errors.New("没有创建过房间")
	}
	return nil
}

//创建新房间
func CreateRoom(mod *RoomBind)error {
	m :=&RoomModel{
		RoomName:mod.RoomName,
		RoomOwner: mod.RoomOwnerName,
	}
	err:=CreateNew(m)
	if err != nil {
		return fmt.Errorf("'创建失败：%s",err.Error())
	}
	return nil
}

