package employment

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// EmploymentGet godoc
// @Summary Get List of Employment
// @Description Get List of Employment
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Employment "Success"
// @Router /Employment [get]
// @Tags Employment
func GetAllEmployment(c *fiber.Ctx) error {
	db := services.DB

	var employment []model.Employment
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.Employment{}).Find(&employment)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    employment,
	})
}
