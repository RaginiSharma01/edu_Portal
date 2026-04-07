package main

import (
	"fmt"
	"log"
	"smp/config"
	"smp/routes"
	"smp/wire"

	"github.com/gofiber/fiber/v3"
)

func main() {

	fmt.Println("Starting Student Management Portal")

	cfg := config.LoadConfig()

	app := fiber.New()

	handlers := wire.InitializeHandlers()

	routes.SetupUserRoutes(app, handlers.UserHandler, handlers.ClassroomHandler, handlers.EventHandler)

	log.Fatal(app.Listen(cfg.ServerPort))
}
