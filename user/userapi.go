package user

import (
	"github.com/gin-gonic/gin"
)
type UserItem struct {
	Username string		`json:"username" binding:"required,min=4,min=10"`
	Password string		`json:"password" binding:"required,min=6,max=16"`
}

//注册
func Signup(c *gin.Context)  {
	item := &UserItem{}
	if err:=c.BindJSON(item);err!=nil{
		c.JSON(401,err.Error())
	}
	err:=CheckUserName(item)
	if err == nil {  //用户名重复
		c.JSON(403,gin.H{
			"message":"用户名重复",
		})
		return
	}
	if err=NewUser(item);err!=nil{
		c.JSON(500,gin.H{
			"message":"用户注册失败",
			"data":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"注册成功",
	})
}

//登录
func Login(c *gin.Context)  {
	item :=&UserItem{}
	if err:=c.BindJSON(item);err!=nil{
		c.JSON(401,err.Error())
		return
	}
	err :=CheckPassword(item)
	if err != nil {
		c.JSON(403,gin.H{
			"message":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"登录成功",
	})


}