package profile

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostProfile go doc
// @Summary Create new Profile
// @Description create new Profile
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.ProfileAPI "created"
// @Success 400
// @Success 401
// @Param data body model.ProfileAPI true "Profile data"
// @Router /Profile [post]
// @Tags Profile
func ProfilePost(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var profile model.ProfileAPI

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_profile := &model.Profile{ProfileAPI: profile}
	db.Model(&model.Profile{}).Create(create_profile)

	return c.Status(201).JSON(create_profile)

}
