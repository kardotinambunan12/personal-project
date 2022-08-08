package model

import "github.com/go-openapi/strfmt"

type EducationAPI struct {
	School      string       `json:"school,omitempty"`
	Degree      string       `json:"degree,omitempty"`
	StartDate   *strfmt.Date `json:"start_date,omitempty" gorm:"type:date" format:"date" swaggertype:"string"`
	EndDate     *strfmt.Date `json:"end_date,omitempty" gorm:"type:date" format:"date" swaggertype:"string"`
	City        string       `json:"city,omitempty"`
	Description string       `json:"description,omitempty"`
}
