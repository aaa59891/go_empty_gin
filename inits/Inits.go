package inits

import (
	"github.com/aaa59891/go_empty_gin/db"
	"github.com/googollee/go-socket.io"
)

var SocketServer *socketio.Server

func init() {
	// var err error
	// SocketServer, err = socketio.NewServer(nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// SocketServer.On("connection", func(so socketio.Socket) {
	// 	log.Println("socket connected.")
	// 	so.Join("mosi")
	// 	so.On("disconnection", func() {
	// 		log.Println("on disconnect")
	// 	})
	// })

	// SocketServer.On("error", func(so socketio.Socket, err error) {
	// 	log.Panic("Socket server had an error, ", err)
	// })
}

func CreateTable() {
	modelArr := make([]interface{}, 0)

	for _, model := range modelArr {
		db.DB.Set("gorm:table_options", "CHARACTER SET = utf8").AutoMigrate(model)
	}
}

func RegisterStruct() {
}
