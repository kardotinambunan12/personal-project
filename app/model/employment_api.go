package model

import "github.com/go-openapi/strfmt"

type EmploymentAPI struct {
	JobTitle    string       `json:"job_title,omitempty"`
	Employer    string       `json:"employer,omitempty"`
	StartDate   *strfmt.Date `json:"start_date,omitempty" gorm:"type:date" format:"date" swaggertype:"string"`
	EndDate     *strfmt.Date `json:"end_date,omitempty" gorm:"type:date" format:"date" swaggertype:"string"`
	City        string       `json:"city,omitempty"`
	Description string       `json:"description,omitempty"`
}
