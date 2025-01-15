package main

import (
	"crm/database"
	"crm/lead"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setup routes:
func setupRoutes(app *fiber.App) {
	app.Get("api/1/leads", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.CreateLead)
	//app.Put("UpdateLead")
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		// handle the error, e.g., log it or panic
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
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
