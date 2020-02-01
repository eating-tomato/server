package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)
const jsonContentType = "application/json"

var (
	Success = map[string]string{"code" : "200", "msg" : "success"}
	ErrUserNotFound = map[string]string{"code" : "20001", "msg" : "user not found"}
	ErrInvalidRequest = map[string]string{"code" : "10001", "msg" : "invalid request"}
	SystemErr = map[string]string{"code" : "500", "msg" : "system error"}
)

func HttpRespond(w http.ResponseWriter, code map[string]string, data interface{}){
	w.Header().Set("content-type", jsonContentType)
	httpStatusCode := 200
	if reflect.DeepEqual(code, SystemErr) {
		httpStatusCode = 500
	}
	w.WriteHeader(httpStatusCode)
	fmt.Fprint(w, string(Respond(code, data)))
}

func Respond(code map[string]string, data interface{}) []byte {
	responseMap := make(map[string]interface{})
	for key, value := range code {
		responseMap[key] = value
	}
	responseMap["data"] = data
	b, _ := json.Marshal(responseMap)
	return b
}

func BodyToStringMap(r *http.Request)map[string]string{
	body, _ := ioutil.ReadAll(r.Body)
	m := JSONToStringMap(string(body))
	return m
}