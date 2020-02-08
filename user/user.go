package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"server/database"
	"server/helper"
	"strings"
	"time"
)

type Info struct {
	Username string
	Id string
	Score int
	RoomId string
}

func Login(w http.ResponseWriter, r *http.Request) {
	m := helper.BodyToStringMap(r)
	info := Info{m["username"], createId(m["username"]), 0, ""}
	database.Set(info.Id, info)
	helper.HttpRespond(w, helper.Success, info)
}

func GetInfo(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if len(r.Form["id"]) <= 0{
		helper.HttpRespond(w, helper.ErrInvalidRequest, nil)
		return
	}
	infoByte := database.Get(r.Form["id"][0])
	var info Info
	if infoByte == nil {
		helper.HttpRespond(w, helper.ErrUserNotFound, nil)
		return
	}
	json.Unmarshal(infoByte, &info)
	helper.HttpRespond(w, helper.Success, info)
}

func createId(name string) string{
	h := md5.New()
	h.Write([]byte(name + time.Now().String()))
	return "USER_" + strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}