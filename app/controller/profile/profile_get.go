package profile

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ProfileGet godoc
// @Summary Get List of Profile
// @Description Get List of Profile
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Profile "Success"
// @Router /Profile [get]
// @Tags Profile
func GetAllprofile(c *fiber.Ctx) error {
	db := services.DB
	var profile []model.Profile
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.Profile{}).Find(&profile)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    profile,
	})
}
