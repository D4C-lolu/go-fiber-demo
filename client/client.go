package client

import (
	"github.com/D4C-lolu/go-fiber-demo/database"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Client struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetClients(c *fiber.Ctx) {
	db := database.DBConn
	var clients []Client
	db.Find(&clients)
	c.JSON(clients)

}

func GetClient(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var client Client
	db.Find(&client, id)
	c.JSON(client)

}

func CreateClient(c *fiber.Ctx) {
	db := database.DBConn
	client := new(Client)
	if err := c.BodyParser(client); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&client)
	c.JSON(client)

}

func DeleteClient(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var client Client
	db.First(&client, id)
	if client.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&client)
	c.Send("Client successfully deleted!")

}
