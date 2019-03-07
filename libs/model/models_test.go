package model

import (
	"strings"
	"testing"

	//	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/septianw/bara/modules/user"
)

var uname = "pertama"
var upass = "ini password saya"
var umodpass = "ini bukan password saya"

func TestNewUser(t *testing.T) {
	db, _ := Connect()
	var dbu = new(User)

	var u = user.New(UserModel)
	u.User = uname
	u.Pass = upass
	err := u.Save()
	t.Log(string(u.Hash))
	t.Log(u)

	if err != nil {
		t.Errorf("Save user should return nil if user is saved instead of %+v", err)
	}

	db.Where("Username = ?", uname).Find(&dbu)
	t.Log(dbu)
	if (strings.Compare(dbu.Username, u.User) != 0) || (strings.Compare(dbu.Hash, string(u.Hash)) != 0) {
		t.Errorf("Save user fail to stored to database.")
	}

	err = u.Save()
	t.Log(string(u.Hash))

	if err == nil {
		t.Errorf("Save duplicate user should return error instead of %+v", err)
	}
}

func TestRetrieveUser(t *testing.T) {
	var u = user.New(UserModel)

	t.Log(u)

	err := u.Retrieve(uname)

	if err != nil {
		t.Errorf("User retrieve should return nil instead of %+v", err)
	}
}

func TestModifyUser(t *testing.T) {
	var u = user.New(UserModel)
	var err error

	u.Retrieve(uname)
	t.Log(string(u.Hash))
	err = u.Modify(upass, umodpass)
	t.Log(string(u.Hash))

	if err != nil {
		t.Errorf("User modify should return nil instead of %+v", err)
	}
}

func TestRemoveUser(t *testing.T) {
	var u = user.New(UserModel)
	var err error

	err = u.Remove()

	if err == nil {
		t.Errorf("User remove should return error instead of %+v", err)
	}

	err = u.Retrieve(uname)
	t.Log(err)
	t.Log(u)

	err = u.Remove()
	t.Log(err)
	t.Log(u)

	if err != nil {
		t.Errorf("User remove should return nil instead of %+v", err)
	}

	if strings.Compare(u.User, "") != 0 {
		t.Errorf("User remove should also remove object value instead of %+v", u)
	}
}
