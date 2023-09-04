package utils

import (
	"bufio"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"path/filepath"
)

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
		log.Error("获取项目根目录失败：%v", err.Error())
	}

	// 构造文件的绝对路径
	absPath := filepath.Join(rootDir, filename)
	absPath = filepath.Clean(absPath)

	return absPath, nil
}
