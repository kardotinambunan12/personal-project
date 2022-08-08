package profile

import (
	"net/http"
	"personal-project/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetAllProfile(t *testing.T) {

	app := fiber.New()

	services.InitDatabaseForTest()
	app.Get("/profile", GetAllprofile)

	// positive case
	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response success")

	//negative case
	req, _ = http.NewRequest("GET", "/profile", nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response success")
}
