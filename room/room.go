package room

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"math/rand"
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
	UserConn map[string]*websocket.Conn
}

//todo 创建一个全局的map map[房间id][websocket.Conn的map]

type ConnMap map[string]roomConns

type roomConns map[string]*websocket.Conn

var connMap ConnMap

func Init(){
	connMap = make(ConnMap)
}

func Chat(ws *websocket.Conn){
	var err error
	var info Info
	for {
		var body string
		var result []byte

		if err = websocket.Message.Receive(ws, &body); err != nil {
			//TODO 离开逻辑
			fmt.Println(err)
			break
		}

		m := make(map[string]string)
		if err = json.Unmarshal([]byte(body), &m); err != nil {
			fmt.Println("json解析错误：" + err.Error())
			continue
		}

		switch m["method"] {
		case "create":
			result = create(m, &info, ws)
			break
		case "join":
			result = join(m, &info, ws)
			break
		case "send":
			send(m, info)
			result = nil
			break
		default:
			result = []byte("{'code': 500;'msg' : 'wrong msg format'}")
		}

		if result == nil {
			continue
		}

		if err = websocket.Message.Send(ws, string(result)); err != nil {
			fmt.Println("下发数据时错误：" + err.Error())
			continue
		}
	}
}

func send(m map[string]string, info Info){
	for _, ws := range info.UserConn {
		if err := websocket.Message.Send(ws, string(result)); err != nil {
			fmt.Println("下发数据时错误：" + err.Error())
			continue
		}
	}
}

func create(m map[string]string, info *Info, ws *websocket.Conn) (result []byte){
	if _ ,ok := m["user_id"]; !ok {
		return helper.Respond(helper.ErrInvalidRequest, nil)
	}
	userInfoByte := database.Get(m["user_id"])
	userInfo := user.Info{}
	if userInfoByte == nil {
		return helper.Respond(helper.ErrUserNotFound, nil)
	}
	json.Unmarshal(userInfoByte, &userInfo)
	info = &Info{createId(), statusReady, make(map[string]user.Info), make(map[string]*websocket.Conn)}
	info.User[userInfo.Id] = userInfo
	info.UserConn[userInfo.Id] = ws
	database.Set(info.Id, info)
	connMap[info.Id] = map[string]*websocket.Conn{userInfo.Id : ws}
	return helper.Respond(helper.Success, info)
}

func join(m map[string]string, info *Info, ws *websocket.Conn) (result []byte){
	if _ ,ok := m["user_id"]; !ok {
		return helper.Respond(helper.ErrInvalidRequest, nil)
	}
	if _ ,ok := m["room_id"]; !ok {
		return helper.Respond(helper.ErrInvalidRequest, nil)
	}

	roomByte := database.Get(m["room_id"])
	room := Info{}
	if roomByte == nil {
		return helper.Respond(helper.ErrRoomNotFound, nil)
	}
	json.Unmarshal(roomByte, &room)
	if room.Status != statusReady{
		return helper.Respond(helper.ErrRoomUnavailable, nil)
	}

	userInfoByte := database.Get(m["user_id"])
	userInfo := user.Info{}
	if userInfoByte == nil {
		return helper.Respond(helper.ErrUserNotFound, nil)
	}

	room.User[userInfo.Id] = userInfo
	room.UserConn[userInfo.Id] = ws
	database.Set(room.Id, room)
	info = &room
	connMap[info.Id][userInfo.Id] = ws
	return helper.Respond(helper.Success, room)
}

func createId() string{
	h := md5.New()
	h.Write([]byte(string(rand.Int()) + time.Now().String()))
	return "ROOM_" + strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
