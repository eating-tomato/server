package room

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"
	"server/database"
	"server/helper"
	"server/user"
	"strings"
	"time"
)

var (
	statusReady = 0
	statusPlaying = 1
)

type Info struct {
	Id string
	Status int
	User map[string]user.Info
}

func Create(w http.ResponseWriter, r *http.Request){
	m := helper.BodyToStringMap(r)
	if _ ,ok := m["user_id"]; !ok {
		helper.Respond(helper.ErrInvalidRequest, nil);return
	}
	userInfoByte := database.Get(m["user_id"])
	userInfo := user.Info{}
	if userInfoByte == nil {
		helper.Respond(helper.ErrUserNotFound, nil);return
	}
	json.Unmarshal(userInfoByte, &userInfo)
	room := Info{createId(), statusReady, make(map[string]user.Info)}
	room.User[userInfo.Id] = userInfo
	database.Set(room.Id, room)
	helper.HttpRespond(w, helper.Success, room)
}

func Join(w http.ResponseWriter, r *http.Request){
	m := helper.BodyToStringMap(r)
	if _ ,ok := m["user_id"]; !ok {
		helper.Respond(helper.ErrInvalidRequest, nil);return
	}
	if _ ,ok := m["room_id"]; !ok {
		helper.Respond(helper.ErrInvalidRequest, nil);return
	}

	roomByte := database.Get(m["room_id"])
	room := Info{}
	if roomByte == nil {
		helper.Respond(helper.ErrRoomNotFound, nil);return
	}
	json.Unmarshal(roomByte, &room)
	if room.Status != statusReady{
		helper.Respond(helper.ErrRoomUnavailable, nil);return
	}

	userInfoByte := database.Get(m["user_id"])
	userInfo := user.Info{}
	if userInfoByte == nil {
		helper.Respond(helper.ErrUserNotFound, nil);return
	}

	room.User[userInfo.Id] = userInfo
	database.Set(room.Id, room)
	helper.Respond(helper.Success, room)
}

func createId() string{
	h := md5.New()
	h.Write([]byte(string(rand.Int()) + time.Now().String()))
	return "ROOM_" + strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
