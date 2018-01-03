package main

import (
	"github.com/aaa59891/go_empty_gin/configs"
	"github.com/aaa59891/go_empty_gin/controllers"
	"github.com/aaa59891/go_empty_gin/db"
	"github.com/aaa59891/go_empty_gin/inits"
	"github.com/aaa59891/go_empty_gin/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.CreateTable()
	inits.RegisterStruct()
}

func main() {
	defer db.DB.Close()
	config := configs.GetConfig()
	r := gin.Default()

	r.GET("/socket.io/", gin.WrapH(inits.SocketServer))
	r.NoRoute(controllers.ErrorHandler404)
	routers.SetRoutes(r)

	r.Run(config.Server.Port)
}
