package employment

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostEmployment go doc
// @Summary Create new Employment
// @Description create new Employment
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.EmploymentAPI "created"
// @Success 400
// @Success 401
// @Param data body model.EmploymentAPI true "Employment data"
// @Router /employment [post]
// @Tags Employment
func EmploymentPost(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var employment model.EmploymentAPI

	if err := c.BodyParser(&employment); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_employment := &model.Employment{EmploymentAPI: employment}
	db.Model(&model.Employment{}).Create(create_employment)

	return c.Status(201).JSON(create_employment)

}
