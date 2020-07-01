package mapper

import (
	"cwm.wiki/ad-CMS/initStep"
	"testing"
)

func TestInsertUser(t *testing.T) {
	initStep.InitGormWithPath("../ad.db")
	InsertUser()
}