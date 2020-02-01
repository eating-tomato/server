package helper

import (
	"encoding/json"
	"fmt"
)

func JSONToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		fmt.Println(err)
		return make(map[string]interface{})
	}

	return tempMap
}

func JSONToStringMap(str string) map[string]string {

	var tempMap map[string]string

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		fmt.Println(err)
		return make(map[string]string)
	}

	return tempMap
}
