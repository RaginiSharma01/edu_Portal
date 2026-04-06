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

	cfg := config.LoadConfig()

	database, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	redisClient := db.InitRedis(cfg)

	// USER MODULE
	userRepo := repository.NewUserRepo(database.Pool)
	userService := service.NewUserService(userRepo, redisClient)
	userHandler := handler.NewUserHandler(userService)

	// CLASSROOM MODULE
	classroomRepo := repository.NewClassroomRepo(database.Pool)
	classroomService := service.NewClassroomService(classroomRepo)
	classroomHandler := handler.NewClassroomHandler(classroomService)

	app := fiber.New()

	// ROUTES
	routes.SetupUserRoutes(app, userHandler, classroomHandler)

	log.Fatal(app.Listen(cfg.ServerPort))
}
