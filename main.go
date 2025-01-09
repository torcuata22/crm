package main

import (
	"crm/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setup routes:
func setupRoutes(app *fiber.App) {
	app.Get("GetLeads")
	app.Get("GetLead")
	app.Post("CreateLead")
	//app.Put("UpdateLead")
	app.Delete("DeleteLead")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		// handle the error, e.g., log it or panic
		log.Fatal(err)
	}
}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	db, err := database.DBConn.DB()
	if err != nil {
		// handle error
	}
	defer db.Close()
}
