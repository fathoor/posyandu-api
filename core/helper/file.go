package helper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"os"
	"path"
)

func GenerateFilePath(fileType, fileExt string) string {
	return path.Join("storage", fileType, uuid.New().String()+fileExt)
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func GetFile(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", exception.NotFoundError{
			Message: "File not found",
		}
	}

	return filePath, nil
}
