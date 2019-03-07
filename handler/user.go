package handler

import (
	"encoding/json"
	"io"
	//	"io/ioutil"
	//	"log"
	"net/http"

	"github.com/septianw/bara/libs/model"
	"github.com/septianw/bara/modules/user"
)

type User struct {
	Id       int8
	Username string
	Hash     string
	salt     string
}

var userModel = model.UserModel

func PostUser(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	var usr = user.New(userModel)
	Writelog("debug", "user Object", "object", usr)

	//	var body interface{}
	var dec = json.NewDecoder(r.Body)

	if err := dec.Decode(&body); err != io.EOF {
		Writelog("debug", "Ini isi error", "err", err)
		if err == io.EOF {
			Writelog("debug", "EOF", "body", body)
		} else {
			Writelog("debug", "Ini body", "body", body)
			Writelog("warn", "Fail to decode JSON: ", "URL", r.URL, "IP", r.RemoteAddr, "err", err)
			Fail.Message = "Fail to decode JSON"
			Send(w, UNSUPPORTED, JSON, PackToHuman(Fail))
			return
		}
		//		break
	} else {
		Writelog("debug", "Ini isi error", "err", err)

		Fail.Message = "Body empty, body post user cannot be empty."
		Send(w, UNSUPPORTED, JSON, PackToHuman(Fail))
		return
	}
}
