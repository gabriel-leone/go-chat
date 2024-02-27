package router

import (
	"github.com/gabriel-leone/go-chat/internal/user"
	"github.com/gabriel-leone/go-chat/internal/ws"
	"github.com/gin-gonic/gin"
)

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/get-clients/:roomId", wsHandler.GetClients)

	return r
}

func Start(addr string, r *gin.Engine) {
	r.Run(addr)
}
