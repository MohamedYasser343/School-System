package ModifyTables

import (
	"Snap/Controllers"
	"Snap/Database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RemoveData(c *fiber.Ctx) error {
	Controllers.User(c)
	if Controllers.CurrentUser.Permission > 0 {
		var data struct {
			Table      string `json:"Table"`
			ColumnName string `json:"ColumnName"`
			Value      int    `json:"Value"`
		}

		if err := c.BodyParser(&data); err != nil {
			return c.JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		db := Database.DBConnect()
		query := fmt.Sprintf("DELETE FROM `%s` WHERE `%s` = '%v';", data.Table, data.ColumnName, data.Value)
		delete, err := db.Query(query)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"Message": "An Error Has Occured (Invalid Credintials)",
			})
		}
		defer delete.Close()
		return c.JSON(fiber.Map{
			"Message": "Success",
		})
	} else {
		return c.JSON(fiber.Map{
			"Message": "Not Enough Permission, Or Not Logged In",
		})
	}
}
