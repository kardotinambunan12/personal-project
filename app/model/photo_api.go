package model

type PhotoAPI struct {
	ImageProfile *string `json:"image_profile,omitempty" gorm:"type:varchar(255)" example:"scheme://domain.ltd/path/to/image.extension"`
}
