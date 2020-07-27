package room

import "github.com/gin-gonic/gin"

type RoomBind struct {
	RoomName string		`json:"room_name" binding:"required,gt=4"`
	RoomOwnerName string 	`json:"room_owner_name" binding:"required,gt=4"`
}

func CreateNewRoom(c *gin.Context)  {
	mod :=&RoomBind{}
	if err:=c.BindJSON(mod);err!=nil{
		c.JSON(403,gin.H{
			"message":err.Error(),
		})
		return
	}
	err := CheckRoomOwner(mod)
	if err == nil {
		c.JSON(403,gin.H{
			"message":"您已经创建过房间",
		})
		return
	}
	if err := CreateRoom(mod);err !=nil{
		c.JSON(500,gin.H{
			"message":"创建失败",
			"data":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"创建房间成功",
	})


}