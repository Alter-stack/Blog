package api

import (
	"github.com/Alter/blog/pkg/logging"
	"github.com/Alter/blog/pkg/retcode"
	"github.com/Alter/blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	code := retcode.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		ResponseData(c, retcode.ERROR, data)
		return
	}

	if image == nil {
		ResponseData(c, retcode.INVALID_PARAMS, data)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName
	if ! upload.CheckImageExt(imageName) || ! upload.CheckImageSize(file) {
		ResponseData(c, retcode.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, data)
		return
	}

	if err := upload.CheckImage(fullPath); err != nil {
		logging.Warn(err)
		ResponseData(c, retcode.ERROR_UPLOAD_CHECK_IMAGE_FAIL, data)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		ResponseData(c, retcode.ERROR_UPLOAD_SAVE_IMAGE_FAIL, data)
		return
	}
	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	ResponseData(c, code, data)
}
