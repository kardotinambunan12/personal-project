package skill

import (
	"fmt"
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Deleted Skill go Doc
// @Summary Delete Skill by id
// @Description Delete Skill by id
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "Skill ID"
// @Router /skill/{id} [delete]
// @Tags Skill
func SkillDelete(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	id := c.Params("id")
	response := db.Where(`id=?`, id).Delete(&model.Skill{})
	data := response.RowsAffected
	msg := fmt.Sprintf(`education  %d has deleted`, data)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": msg,
	})

}
