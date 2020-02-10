
package main

import (
	"server/database"
	"server/room"
	"server/server"
)




func main() {
	database.Conn()
	room.Init()
	server.Serve()
}
