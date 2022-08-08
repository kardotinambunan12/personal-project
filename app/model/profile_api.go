package model

import "github.com/go-openapi/strfmt"

type ProfileAPI struct {
	WantedJobTitle *string      `json:"wanted_job_title,omitempty"`
	FirstName      *string      `json:"first_name,omitempty"`
	LastName       *string      `json:"last_name,omitempty"`
	Email          *string      `json:"email,omitempty"`
	Phone          *string      `json:"phone,omitempty"`
	Country        *string      `json:"country,omitempty"`
	City           *string      `json:"city,omitempty"`
	Address        *string      `json:"address,omitempty"`
	PostalCode     *int         `json:"postal_code,omitempty" gorm:"type:smallint; null"`
	DrivingLicense *string      `json:"driving_license,omitempty"`
	Nationality    *string      `json:"nationality,omitempty"`
	PlaceOfBirth   *string      `json:"place_of_birth,omitempty"`
	DateOfBirth    *strfmt.Date `json:"date_of_birth,omitempty" gorm:"type:date" format:"date" swaggertype:"string"` // Date Of Birth
}
