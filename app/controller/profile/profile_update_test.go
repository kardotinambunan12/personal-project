package profile

import (
	"bytes"
	"net/http"

	// "personal-project/app/model"

	// "strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutProfile(t *testing.T) {
	// services.InitDatabase()
	// db := services.DB
	app := fiber.New()

	app.Put("/profile/:id", ProfilePut)

	// address := "Test string"
	// city := "Konoha"
	// country := "Indonesia"
	// wanted_job_title := "Anonim"
	// first_name:="Jhon"
	// last_name:="DOe"
	// email := "anonim@mail.com"
	// phone:= "08877777"
	// postal_code:= 12334
	// driving_license := "test licence"
	// nationality := "test natioanality"
	// place_of_birth:= "Humbahas"
	// date_of_birth := "2022-08-07"

	// init := model.Profile{
	// 	ProfileAPI: model.ProfileAPI{
	// 		Address:        &address,
	// 		City:           &city,
	// 		Country:        &country,
	// 		WantedJobTitle: &wanted_job_title,
	// 		FirstName:      &first_name,
	// 		LastName:       &last_name,
	// 		Email:          &email,
	// 		Phone:          &phone,
	// 		PostalCode:     &postal_code,
	// 		DrivingLicense: &driving_license,
	// 		Nationality:    &nationality,
	// 		PlaceOfBirth:   &place_of_birth,
	// 		DateOfBirth: &date_of_birth,

	// 	},
	// }
	// db.Create(&init)

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

	// id := strconv.Itoa(*init.ID)
	id := "1"

	uri := "/profile/" + id

	//positive case
	req, _ := http.NewRequest("PUT", uri, payload)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "put response success")

	//negative case
	req, _ = http.NewRequest("PUT", uri, nil)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response error payload header")

	req, _ = http.NewRequest("PUT", uri, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response error content-type header")

}
