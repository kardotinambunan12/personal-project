package educations

import (
	"bytes"
	"net/http"
	"personal-project/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestEducationPost(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()
	app.Post("/education", EducationPost)

	payload := bytes.NewReader([]byte(`
	{
		"city": "Test City",
		"degree": "Test degree",
		"description": "Test Description",
		"end_date": "2022-08-07",
		"school": "Test school",
		"start_date": "2022-08-07"
	  }
	`))

	//positive case
	req, _ := http.NewRequest("POST", "/education", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	// //negative case
	req, _ = http.NewRequest("POST", "/education", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	// negative case
	req, _ = http.NewRequest("POST", "/education", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	//negative case
	req, _ = http.NewRequest("POST", "/education", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	req, _ = http.NewRequest("POST", "/education", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "post response body parser")

}
