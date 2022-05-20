package main

import (
	"fmt"

	"github.com/D4C-lolu/go-fiber-demo/client"
	"github.com/D4C-lolu/go-fiber-demo/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/client/", client.GetClients)
	app.Get("/api/v1/client/:id", client.GetClient)
	app.Post("/api/v1/client/", client.CreateClient)
	app.Delete("/api/v1/client/:id", client.DeleteClient)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "clients.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&client.Client{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
