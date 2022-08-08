package photo

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PhotoGet godoc
// @Summary Get List of Photo
// @Description Get List of Photo
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Photo "Success"
// @Router /photo [get]
// @Tags Photo
func GetPhoto(c *fiber.Ctx) error {
	db := services.DB

	var photo []model.Photo
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.Photo{}).Find(&photo)
	return c.JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    photo,
	})

}
