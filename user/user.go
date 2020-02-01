package user

import (
	"crypto/md5"
	"encoding/hex"
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
	info := database.Get(r.Form["id"][0])
	if info == nil {
		helper.HttpRespond(w, helper.ErrUserNotFound, nil)
		return
	}
	helper.HttpRespond(w, helper.Success, info)
}

func createId(name string) string{
	h := md5.New()
	h.Write([]byte(name + time.Now().String()))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}