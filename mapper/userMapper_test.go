package mapper

import (
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"testing"
)

func TestInsertUser(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	i := vo.UserInput{
		Username: "test",
		Password: "test",
		Type:     0,
	}

	err := InsertUser(i)
	if err != nil {
		t.Error(err)
	}
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

func TestUpdateUser(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	i := vo.UserInput{
		SystemId: 2,
		Password: "12345",
	}

	rtv, err := UpdateUser(i)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(rtv)
}

func TestDeleteUser(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	err := DeleteUser("2")
	if err != nil {
		t.Error(err)
	}

	t.Log(true)
}

func TestAli(t *testing.T) {

	//i := 9
	nums := []int{1,2,3,4}
	var rtv []int

	for index,_ := range nums {
		sum := 0
		for i:=0;i<=index;i++ {
			sum += nums[i]
		}

		rtv = append(rtv,sum)
		sum = 0
	}

	t.Log(rtv)

}
