package photo

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostPhoto go doc
// @Summary Create new Photo
// @Description create new Photo
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.PhotoAPI "created"
// @Success 400
// @Success 401
// @Param data body model.PhotoAPI true "Photo data"
// @Router /photo [post]
// @Tags Photo
func PostPhoto(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var photo model.PhotoAPI

	if err := c.BodyParser(&photo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_photo := &model.Photo{PhotoAPI: photo}
	db.Model(&model.Photo{}).Create(create_photo)

	return c.Status(201).JSON(create_photo)
}
