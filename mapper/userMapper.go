package mapper

import (
	"cwm.wiki/ad-CMS/initStep"
	"cwm.wiki/ad-CMS/model"
	"time"
)

func InsertUser() {

	user := model.Users{
		Username: "admin",
		Password: "admin",
		Type: 1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),

	}

	initStep.DB.Create(&user)

}