package utils

import (
	"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func MultiPartFileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 转为字节数组
	imgBytes := make([]byte, fileHeader.Size)
	_, err = file.Read(imgBytes)
	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

func GetRootDir() (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	rootDir = filepath.Dir(rootDir)

	// 确保rootDir不是空字符串，并且是一个存在的目录。
	if rootDir == "" || !isDir(rootDir) {
		return "", fmt.Errorf("unable to determine root directory")
	}

	return rootDir, nil
}

// isDir 检查给定的路径是否是一个目录。
func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// CopyFile 高效地拷贝文件，使用底层操作系统的零拷贝特性，不需要将整个文件的内容加载到内存中。
func CopyFile(srcPath, dstPath string) (err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer func() { _ = srcFile.Close() }()

	dstFile, err := CreateFile(dstPath)
	if err != nil {
		return
	}
	defer func() { _ = dstFile.Close() }()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return
	}

	err = dstFile.Sync()
	return
}

func NotExistFile(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func WriteFile(path string, dataBytes []byte) error {
	if err := MkdirAll(filepath.Dir(path)); err != nil {
		return err
	}

	return os.WriteFile(path, dataBytes, 0644)
}

func CreateFile(path string) (file *os.File, err error) {
	if err = MkdirAll(filepath.Dir(path)); err != nil {
		return
	}

	return os.Create(path)
}

func ReadFile(path string) (content []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if content != nil {
			content = append(content, '\n')
		}
		content = append(content, fileScanner.Bytes()...)
	}
	return
}

func MkdirAll(dirname string) error {
	return os.MkdirAll(dirname, os.ModePerm)
}

func GetAbsPath(filename string) (string, error) {
	// 获取项目根目录路径
	rootDir, err := os.Getwd()
	if err != nil {
		zap.L().Error(fmt.Sprintf("获取项目根目录失败：%v", err.Error()))
	}

	// 构造文件的绝对路径
	absPath := filepath.Join(rootDir, filename)
	absPath = filepath.Clean(absPath)

	return absPath, nil
}
