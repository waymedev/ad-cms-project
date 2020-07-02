package mapper

import (
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"testing"
)

func TestInsertUser(t *testing.T) {
	model.InitGormWithPath("../ad.db")
	InsertUser()
}

func TestSelectUser(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	userInput := vo.UserInput{
		Username: "admin1",
		Password: "admin",
	}

	user, err := SelectUserByCondition(userInput)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(*user)
}

func TestSelectUsers(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	users, err := SelectUsers()
	if err != nil {
		t.Error(err)
	}

	t.Log(users)
}

func TestSelectUser2(t *testing.T) {
	model.InitGormWithPath("../ad.db")



	user, err := SelectUser("1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(*user)


}
