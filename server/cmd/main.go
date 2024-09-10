package main

import (
	"gopractice2/server/database"
	"gopractice2/server/internal/message"
	"gopractice2/server/internal/user"
	"gopractice2/server/internal/ws"
	"gopractice2/server/router"
	"log"
	// "github.com/redis/go-redis/v9"
)

func main() {
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })
	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}
	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	//initialize new with the correct connection
	messageRep := message.NewRepository(dbConn.GetDB())
	wsHandler := ws.NewWebsocketServer(messageRep)
	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}
