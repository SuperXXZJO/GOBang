package user

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func InitDB()  {
	db,err:=gorm.Open("mysql", "root:root@/GOBang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("数据库连接失败：",err)
	}
	db.AutoMigrate(&User{})
	DB = db

}
func init() {
	InitDB()
}