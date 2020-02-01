package server

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"server/room"
	"server/user"
	"strings"
)


func Serve(){
	http.Handle("/chat", websocket.Handler(chat))
	http.Handle("/picture", websocket.Handler(picture))
	http.HandleFunc("/room/create", room.Create)
	http.HandleFunc("/room/join", room.Join)
	http.HandleFunc("/user/login", user.Login)
	http.HandleFunc("/user/info", user.GetInfo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func chat(ws *websocket.Conn) {
	var err error
	//var userInfo user.Info
	for {
		var body string
		var result []byte

		if err = websocket.Message.Receive(ws, &body); err != nil {
			fmt.Println(err)
			break
		}

		m := make(map[string]string)
		if err = json.Unmarshal([]byte(body), &m); err != nil {
			fmt.Println("28行错误：" + err.Error())
			continue
		}

		switch m["method"] {
		case "login":
			//result = user.Login(m, &userInfo)
			break
		default:
			result = []byte("{'code': 500;'msg' : 'wrong msg format'}")
		}


		if err = websocket.Message.Send(ws, strings.ToLower(string(result))); err != nil {
			fmt.Println("43行错误：" + err.Error())
			continue
		}
	}
}
func picture(ws *websocket.Conn) {
	var err error
	//var userInfo user.Info
	for {
		var body string
		var result []byte

		if err = websocket.Message.Receive(ws, &body); err != nil {
			fmt.Println(err)
			break
		}

		m := make(map[string]string)
		if err = json.Unmarshal([]byte(body), &m); err != nil {
			fmt.Println("28行错误：" + err.Error())
			continue
		}

		switch m["method"] {
		case "login":
			//result = user.Login(m, &userInfo)
			break
		default:
			result = []byte("{'code': 500;'msg' : 'wrong msg format'}")
		}


		if err = websocket.Message.Send(ws, strings.ToLower(string(result))); err != nil {
			fmt.Println("43行错误：" + err.Error())
			continue
		}
	}
}