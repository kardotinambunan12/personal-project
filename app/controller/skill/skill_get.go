package skill

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// SkillGet godoc
// @Summary Get List of Skill
// @Description Get List of Skill
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Skill "Success"
// @Router /Skill [get]
// @Tags Skill
func GetAllSkill(c *fiber.Ctx) error {
	db := services.DB

	var skils []model.Skill
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.Skill{}).Find(&skils)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    skils,
	})
}
