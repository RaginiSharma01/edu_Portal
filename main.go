package main

import (
	"fmt"
	"log"
	"smp/config"
	"smp/db"
	"smp/handler"
	"smp/repository"
	"smp/routes"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

func main() {

	fmt.Println("Starting Student Management Portal")

	cfg := config.LoadConfig() // load config

	database, err := db.ConnectDb(cfg) // connect postgres
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	// err = db.SeedAdmin(database.Pool)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	redisClient := db.InitRedis(cfg) // connect redis

	userRepo := repository.NewUserRepo(database.Pool) // repositories

	userService := service.NewUserService(userRepo, redisClient) // services

	userHandler := handler.NewUserHandler(userService) // handlers

	app := fiber.New()

	routes.SetupUserRoutes(app, userHandler) // routes
	log.Fatal(app.Listen(cfg.ServerPort))
}
