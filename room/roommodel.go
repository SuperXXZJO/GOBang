package room

import "github.com/jinzhu/gorm"

type RoomModel struct {
	gorm.Model
	RoomName string
	RoomOwner string
}

//新建房间
func CreateNew(mod *RoomModel) error {
	 if err:=DB.Create(mod).Error;err!=nil{
		 return err
	 }
	return nil
}

//根据房间名查询房间
func FindRoomByRoomName(roomname string) (*RoomModel,error) {
	res :=&RoomModel{}
	if err:=DB.Where("room_name = ?",roomname).First(res).Error;err!=nil{
		return &RoomModel{},err
	}
	return  res,nil
}