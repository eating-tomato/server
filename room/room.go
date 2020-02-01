package room

import (
	"net/http"
	"server/user"
)

type Info struct {
	Id string
	Status int
	User map[string]user.Info
}

func Create(w http.ResponseWriter, r *http.Request){

}

func Join(w http.ResponseWriter, r *http.Request){

}
