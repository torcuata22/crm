package lead

import (
	"github.com/gofiber/fiber/v2"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Message string `json:"message"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func CreateLead(c *fiber.Ctx) error {
	db := database.DBConn

	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}

// func UpdateLead(c *fiber.Ctx) error {

// }

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "lead not found",
		})

	}

	db.Delete(&lead)
	c.Send([]byte("Lead sucessfully deleted"))
	return nil
}
