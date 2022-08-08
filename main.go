package main

import (
	"personal-project/app/config"
	"personal-project/app/lib"
	"personal-project/app/routes"
	"personal-project/app/services"

	// "api/app/config"
	// "api/app/lib"
	// "api/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	lib.LoadEnvironment(config.Environment)
}

// @title My Project
// @version 1.0.0
// @description API Documentation
// @contact.name Developer
// @host localhost:8000
// @schemes http
// @BasePath /api/v1

func main() {
	services.InitDatabase()

	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
