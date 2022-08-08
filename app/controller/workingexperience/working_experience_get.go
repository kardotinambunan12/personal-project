package workingexperience

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// WorkingExperienceGet godoc
// @Summary Get List of WorkingExperience
// @Description Get List of WorkingExperience
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.WorkingExperience "Success"
// @Router /working_experience [get]
// @Tags WorkingExperience
func GetAllWorkingExperience(c *fiber.Ctx) error {
	db := services.DB

	var working_exp []model.WorkingExperience
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.WorkingExperience{}).Find(&working_exp)
	return c.JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    working_exp,
	})
}
