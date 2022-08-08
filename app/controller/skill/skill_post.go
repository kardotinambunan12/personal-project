package skill

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostSkill go doc
// @Summary Create new Skill
// @Description create new Skill
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.SkillAPI "created"
// @Success 400
// @Success 401
// @Param data body model.SkillAPI true "Skill data"
// @Router /skill [post]
// @Tags Skill
func SkillPost(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var skill model.SkillAPI

	if err := c.BodyParser(&skill); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_skill := &model.Skill{SkillAPI: skill}
	db.Model(&model.Skill{}).Create(create_skill)

	return c.Status(201).JSON(create_skill)

}
