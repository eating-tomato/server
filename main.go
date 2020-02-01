
package main

import (
	"server/database"
	"server/server"
)




func main() {
	database.Conn()
	server.Serve()
}
