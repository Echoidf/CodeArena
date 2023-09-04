package embeds

import (
	"embed"
	"go.uber.org/zap"
	"path/filepath"
)

//go:embed static/*
var EmbedStatic embed.FS

const EmbedStaticRoot = "static"

var logger *zap.Logger

func ReadEmbeds(fileName string) []byte {
	fileData, err := EmbedStatic.ReadFile(filepath.Join(EmbedStaticRoot, fileName))
	if err != nil {
		logger.Error("读取文件错误:", zap.Error(err))
	}
	return fileData
}
