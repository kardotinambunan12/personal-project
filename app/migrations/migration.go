package migrations

import "personal-project/app/model"

// "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Profile{},
	&model.Education{},
	&model.Employment{},
	&model.Skill{},
	&model.WorkingExperience{},
	&model.Photo{},
}
