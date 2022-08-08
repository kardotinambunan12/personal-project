package profile

import (
	"fmt"
	"personal-project/app/model"
	"personal-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutProfile Go doc
// @Summary Updated Profile by Id
// @Description udpated Profile by Id
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.ProfileAPI "OK"
// @Success 401
// @Success 404
// @Param id path string true "Profile ID"
// @Param data body model.ProfileAPI true "Profile data"
// @Router /Profile/{id} [put]
// @Tags Profile
func ProfilePut(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")

	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	var profile model.ProfileAPI
	id := c.Params("id")

	if err := c.BodyParser(&profile); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	data_update := &model.Profile{ProfileAPI: profile}
	result := db.Model(&model.Profile{}).Where(`id=?`, id).Updates(data_update)
	data_row := result.RowsAffected
	msg := fmt.Sprintf(`%d row affected`, data_row)

	return c.JSON(fiber.Map{
		"message": msg,
		"status":  201,
		"data":    data_update,
	})

}
