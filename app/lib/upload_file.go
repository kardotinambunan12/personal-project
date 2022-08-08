package lib

// import (
// 	"mime/multipart"
// 	"personal-project/app/model"
// 	"personal-project/app/services"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// func PostPhoto(c *fiber.Ctx) error {
// 	db := services.DB
// 	form, err := c.MultipartForm()
// 	if nil != err {
// 		return ErrorBadRequest(c, err.Error())
// 	}

// 	owner := GetXUserID(c)

// 	var photo *model.Photo

// 	if files, ok := form.File["files"]; ok && len(files) > 0 {
// 		photo = createPhoto(c, db, owner, files[0])
// 	} else {
// 		return ErrorBadRequest(c, "invalid files field")
// 	}

// 	if nil == photo {
// 		return ErrorInternal(c, "Upload failed")
// 	}

// 	db.Create(photo).
// 		First(&photo)

// 	return OK(c, photo)
// }

// func createPhoto(c *fiber.Ctx, db *gorm.DB, owner *uuid.UUID, file *multipart.FileHeader) *model.PhotoAPI{

// 	uploadDirectory := lib.StorageDirectory()
// 	ext := strings.ToLower(regexp.MustCompile(".*\\.([^\\.]+)$").ReplaceAllString(file.Filename, "$1"))
// 	id := uuid.New()
// 	newFileName := id.String() + "." + ext

// 	return nil

// }
