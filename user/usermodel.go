package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}


//创建新的user
func CreateNewUser(mod *User)error  {
	if err:=DB.Create(&mod).Error;err!=nil{
		return err
	}
	return nil
}

//根据用户名查找User
func FindUserByUserName(username string) (res *User,err error) {
	if err = DB.Where("user_name = ?",username).First(&res).Error;err!=nil{
		return &User{},err
	}
	return res,nil
}

