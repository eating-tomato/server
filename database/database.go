package database

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func Conn(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	conn = c
}

func Set(key string, value interface{}){
	valueString, _ := json.Marshal(value)
	if _, e := conn.Do("SET", key, string(valueString));e != nil {
		fmt.Println(e)
	}
}

func Get(key string) interface{}{
	reply, e := redis.String(conn.Do("GET", key))
	if e != nil {
		fmt.Println(e)
	}
	var replyJson interface{}
	json.Unmarshal([]byte(reply), &replyJson)
	return replyJson
}
