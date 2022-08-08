package workingexperience

import (
	"fmt"
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutWorkingExperience Go doc
// @Summary Updated WorkingExperience by Id
// @Description udpated WorkingExperience by Id
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.WorkingExperienceAPI "OK"
// @Success 401
// @Success 404
// @Param id path string true "WorkingExperience ID"
// @Param data body model.WorkingExperienceAPI true "WorkingExperience data"
// @Router /working_experience/{id} [put]
// @Tags WorkingExperience
func PutWorkingexperience(c *fiber.Ctx) error {
	db := services.DB
	accept := c.Get("accept")

	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	var working_experience model.WorkingExperienceAPI
	id := c.Params("id")

	if err := c.BodyParser(&working_experience); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	data_update := &model.WorkingExperience{WorkingExperienceAPI: working_experience}
	result := db.Model(&model.WorkingExperience{}).Where(`id=?`, id).Updates(data_update)
	data_row := result.RowsAffected
	msg := fmt.Sprintf(`%d row affected`, data_row)

	return c.JSON(fiber.Map{
		"message": msg,
		"status":  201,
		"data":    data_update,
	})

}
