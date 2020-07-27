package user

import (
	"errors"
	"fmt"
)

//验证用户名重复
func CheckUserName(mod *UserItem) error {
	_,err:=FindUserByUserName(mod.Username)
	if err != nil {
		return err
	}
	return nil
}

//验证密码
func CheckPassword(mod *UserItem) (error) {
	res,err:=FindUserByUserName(mod.Username)
	if err != nil {
		return fmt.Errorf("用户名不存在:%s",err.Error())
	}
	if res.Password != mod.Password {
		return errors.New("密码错误")
	}
	return nil
}


//user入库
func NewUser(mod *UserItem) error {
	m :=&User{
		Username:mod.Username,
		Password:mod.Password,
	}
	if err:=CreateNewUser(m);err!=nil{
		return fmt.Errorf("用户注册失败：",err)
	}
	return nil
}