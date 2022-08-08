package model

import "github.com/go-openapi/strfmt"

type User struct {
	Base
	UserName  *string          `json:"user_name,omitempty"`
	Email     *string          `json:"email,omitempty"`
	BirthDate *strfmt.DateTime `json:"id,omitempty"`
	Address   *string          `json:"addess,omitempty"`
	Country   *string          `json:"country,omitempty"`
}
