package main

import (
	"log"

	"github.com/gabriel-leone/go-chat/db"
	"github.com/gabriel-leone/go-chat/internal/router"
	"github.com/gabriel-leone/go-chat/internal/user"
	"github.com/gabriel-leone/go-chat/internal/ws"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	log.Println("Database created successfully")

	userRep := user.NewRepository(dbConn.GetDB())
	userSrv := user.NewService(userRep)
	userHandler := user.NewHandler(userSrv)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	r := router.InitRouter(userHandler, wsHandler)
	router.Start(":8080", r)
}
