package educations

import (
	"fmt"
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Deleted Education go Doc
// @Summary Delete Education by id
// @Description Delete Education by id
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "Education ID"
// @Router /education/{id} [delete]
// @Tags Education
func EducationDelete(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	id := c.Params("id")
	// regex := regexp.MustCompile(`^[0-9]+$`)
	// if !regex.MatchString(movie_id) {
	// 	msg := "bad request"
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"msg": msg,
	// 	})
	// }
	response := db.Where(`id=?`, id).Delete(&model.Education{})
	data := response.RowsAffected
	msg := fmt.Sprintf(`education  %d has deleted`, data)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": msg,
	})

}
