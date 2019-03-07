package model

import (
	//	"encoding/base64"
	//	"errors"
	//	"fmt"
	"os"
	//	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/septianw/bara/modules/user"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index"`
	Hash     string
}

func Connect() (*gorm.DB, error) {
	dbtype := os.Getenv("DBTYPE")
	dbstring := os.Getenv("DBCONNECTION")

	return gorm.Open(dbtype, interface{}(dbstring))
}

var UserModel = user.Model{
	func(u *user.User) error { // create
		db, err := Connect()
		defer db.Close()

		us := new(User)
		us.Username = u.User
		us.Hash = string(u.Hash)

		if (err == nil) && !db.HasTable(&User{}) {
			err = db.CreateTable(&User{}).Error
		}

		if (err == nil) && db.NewRecord(us) {
			err = db.Create(&us).Error
		}

		return err
	},
	func(username string, u *user.User) error { // read
		db, err := Connect()
		defer db.Close()
		us := new(User)

		if err == nil {
			err = db.Where("Username = ?", username).Find(&us).Error
		}
		if err == nil {
			u.User = us.Username
			u.Hash = []byte(us.Hash)
		}

		return err
	},
	func(u *user.User) error { // update
		db, err := Connect()
		defer db.Close()

		if err == nil {
			err = db.Model(&User{}).Update("Hash", string(u.Hash)).Error
		}

		return err
	},
	func(u *user.User) error { // delete
		db, err := Connect()
		defer db.Close()

		if err == nil {
			err = db.Where("Username = ?", u.User).Delete(&User{}).Error
		}

		return err
	},
}
