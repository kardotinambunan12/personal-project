package educations

import (
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// EducationGet godoc
// @Summary Get List of Education
// @Description Get List of Education
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Education "Success"
// @Router /Education [get]
// @Tags Education
func GetAllEducation(c *fiber.Ctx) error {
	db := services.DB

	var educations []model.Education
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.Education{}).Find(&educations)

	// if response.RowsAffected == 0 {
	// 	msg := fmt.Sprintf(`movie%s is empty`, movie_id)
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"message": msg,
	// 	})
	// }

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   educations,
	})
}
