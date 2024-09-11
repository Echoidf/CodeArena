package controllers

import (
	"bytes"
	"codearena/pkg/response"
	"codearena/pkg/utils"
	"codearena/plugins/fileupload/aliyun"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func UploadFile(c *fiber.Ctx) error {
	img, err := c.FormFile("file")

	if err != nil {
		return response.BadRequest(c, err)
	}

	// Upload the file
	imgBytes, err := utils.MultiPartFileHeaderToBytes(img)
	if err != nil {
		return response.Error(c, err)
	}

	fPath := filepath.Join("imgs", img.Filename)
	err = aliyun.Bucket.PutObject(fPath, bytes.NewReader(imgBytes))
	if err != nil {
		zap.L().Error("upload file failed", zap.Error(err))
		return response.Error(c, err)
	}

	return response.Success(c, aliyun.GetObjectUrl(fPath))
}
