package educations

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostEducation go doc
// @Summary Create new Education
// @Description create new Education
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.EducationAPI "created"
// @Success 400
// @Success 401
// @Param data body model.EducationAPI true "Education data"
// @Router /education [post]
// @Tags Education
func EducationPost(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var education model.EducationAPI

	if err := c.BodyParser(&education); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_education := &model.Education{EducationAPI: education}
	db.Model(&model.Profile{}).Create(create_education)

	return c.Status(201).JSON(create_education)

}
