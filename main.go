package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7845"
	}
	app := fiber.New()
	app.Use((cors.New()))
	app.Static("/", "./customer_frontend_app/dist")

	fmt.Println("Starting server on port " + port)
	app.Listen(":" + port)

}
