package controllers

import (
	"bytes"
	"content-assistant/pkg/common"
	"content-assistant/pkg/utils"
	"content-assistant/plugins/fileupload/aliyun"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	img, err := c.FormFile("file")

	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}

	// Upload the file
	imgBytes, err := utils.MultiPartFileHeaderToBytes(img)
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}

	fPath := filepath.Join("imgs", img.Filename)
	err = aliyun.Bucket.PutObject(fPath, bytes.NewReader(imgBytes))
	if err != nil {
		zap.L().Error("upload file failed", zap.Error(err))
		common.Error(c, 500, "upload file failed "+err.Error())
		return
	}

	common.Success(c, aliyun.GetObjectUrl(fPath))
}
