package profile

import (
	"bytes"
	"net/http"
	"personal-project/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestProfilePost(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()
	app.Post("/profile", ProfilePost)

	payload := bytes.NewReader([]byte(`
	{
		"address": "Test address",
		"city": "Test City",
		"country": "Test Country",
		"date_of_birth": "2022-08-07",
		"driving_license": "087878787",
		"email": "Admin@mail.com",
		"first_name": "Jhon",
		"last_name": "Doe",
		"nationality": "Test nationality",
		"phone": "0897979",
		"place_of_birth": "Konoha",
		"postal_code": 25362,
		"wanted_job_title": "test Job title"
	}
	`))

	//positive case
	req, _ := http.NewRequest("POST", "/profile", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	// //negative case
	req, _ = http.NewRequest("POST", "/profile", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	// negative case
	req, _ = http.NewRequest("POST", "/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	//negative case
	req, _ = http.NewRequest("POST", "/profile", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	req, _ = http.NewRequest("POST", "/profile", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "post response body parser")

}
