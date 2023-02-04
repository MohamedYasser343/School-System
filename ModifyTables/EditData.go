package ModifyTables

import (
	"Snap/Controllers"
	"Snap/Database"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func EditData(c *fiber.Ctx) error {
	Controllers.User(c)
	if Controllers.CurrentUser.Permission > 0 {
		var data struct {
			Table string `json:"Table"`
			// Json Object with Column Name and Value
			Columns     map[string]string `json:"NewData"`
			Id          string            `json:"Id"`
			IndexColumn string            `json:"IndexColumn"`
		}
		// Unmarshel Json Data
		if err := c.BodyParser(&data); err != nil {
			return c.JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		fmt.Println(data)
		db := Database.DBConnect()
		// Create Query
		query := fmt.Sprintf("UPDATE `%s` SET ", data.Table)
		// Delete Id From Map
		data.Id = data.Columns["Id"]
		delete(data.Columns, "Id")
		// Loop Through Json Object
		for key, value := range data.Columns {
			query += fmt.Sprintf("`%s` = '%s', ", key, value)
		}
		_ = db

		query = strings.TrimSuffix(query, ", ")
		query += fmt.Sprintf(" WHERE `%s` = '%s';", data.IndexColumn, data.Id)
		// Execute Query
		fmt.Println(query)
		update, err := db.Query(query)
		if err != nil {
			log.Println(err)
		}
		defer update.Close()

		return c.JSON(fiber.Map{
			"Message": "Success",
		})
	} else {
		return c.JSON(fiber.Map{
			"Message": "Not Enough Permission, Or Not Logged In",
		})
	}
}
