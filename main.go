package main

import (
	"Snap/Database"
	"Snap/Routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func FiberAuth() {
	Database.Connect()

	engine := html.New("./DashboardFiles", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	Routes.Setup(app)
	app.ListenTLS(":3002", "go-server.crt", "go-server.key")
}

func main() {
	FiberAuth()
	fmt.Println("Server Running...")
}
