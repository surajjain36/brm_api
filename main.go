package main

import (
	"brm_api/middlewares"
	"brm_api/migrations"
	"brm_api/routes"
	"brm_api/services/db/postgres"
	"fmt"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
)

var (
	port = flag.String("port", ":4000", "Port to listen on") // go run . -port=:3000
)

func main() {
	glg.Get().SetMode(glg.STD)
	flag.Parse()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .local.env file")
	}
	if err := postgres.NewConnect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}
	migrations.PreAutoMigrate()
	fmt.Println("PreMigrate")

	fmt.Println("Migrate")
	migrations.Migrate()

	fmt.Println("PostMigrate")
	migrations.PostMigrate()

	app := StartServer()
	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000

}

// StartServer Starting server
func StartServer() *fiber.App {
	app := routes.New()

	// Handle not founds
	app.Use(middlewares.NotFound)

	return app
}
