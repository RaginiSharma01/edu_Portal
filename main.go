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

	userHandler := wire.InitializeUserHandler()

	routes.SetupUserRoutes(app, userHandler)

	log.Fatal(app.Listen(cfg.ServerPort))
}