package main

import (
	"fmt"
	"log"
	"smp/config"
	"smp/routes"
	"smp/wire"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	fmt.Println("Starting Student Management Portal")

	cfg := config.LoadConfig()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	handlers := wire.InitializeHandlers()

	routes.SetupUserRoutes(app,
		handlers.UserHandler,
		handlers.ClassroomHandler,
		handlers.EventHandler,
		handlers.SalaryHandler,
		handlers.TimetableHandler,
		handlers.MarksHandler,
		handlers.DashboardHandler)

	log.Fatal(app.Listen(cfg.ServerPort))
}
