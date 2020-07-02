package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"time"
)

func InsertUser(user vo.UserInput) {

	i := model.Users{
		Username: user.Username,
		Password: user.Password,
		Type: user.Type,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	model.DB.Create(&i)

}

func SelectUserByCondition(userInput vo.UserInput) (*model.Users, error) {
	var user model.Users

	if err := model.DB.Where("username = ? and password = ?", userInput.Username, userInput.Password).First(&user).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &user, nil
}

func SelectUsers() (*[]model.Users, error) {

	var users []model.Users
	if err := model.DB.Find(&users).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &users, nil

}

func SelectUser(id string) (*model.Users,error) {
	var user model.Users
	if err := model.DB.Where("system_id = ?", id).First(&user).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &user,nil
}


