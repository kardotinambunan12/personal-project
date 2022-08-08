package routes

import (
	"personal-project/app/controller"
	"personal-project/app/controller/educations"
	"personal-project/app/controller/employment"
	"personal-project/app/controller/photo"
	"personal-project/app/controller/profile"
	"personal-project/app/controller/skill"
	"personal-project/app/controller/workingexperience"
	"personal-project/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()
	// services.InitRedis()

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)

	// routes profile
	api.Get("/profile", profile.GetAllprofile)
	api.Post("/profile", profile.ProfilePost)
	api.Put("/profile/:id", profile.ProfilePut)

	//routes educations
	api.Get("/education", educations.GetAllEducation)
	api.Get("education", educations.EducationPost)
	api.Delete("/education/:id", educations.EducationDelete)

	//routes employment
	api.Get("/employment", employment.GetAllEmployment)
	api.Post("/employment", employment.EmploymentPost)
	api.Delete("/employment/:id", employment.EmploymentDelete)

	//routes skill
	api.Get("/skill", skill.GetAllSkill)
	api.Post("/skill", skill.SkillPost)
	api.Delete("/skill/:id", skill.SkillDelete)

	//routes working experience
	api.Get("/working_experience", workingexperience.GetAllWorkingExperience)
	api.Delete("/working_experience/:id", workingexperience.PutWorkingexperience)

	//routes photo
	api.Get("/photo", photo.GetPhoto)
	api.Post("/photo", photo.PostPhoto)
	api.Delete("/photo", photo.PhotoDelete)

}
