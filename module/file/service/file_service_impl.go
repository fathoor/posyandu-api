package service

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/helper"
	"github.com/fathoor/posyandu-api/module/file/model"
	"github.com/fathoor/posyandu-api/module/file/validation"
	"path"
	"strings"
)

type fileServiceImpl struct {
}

func (service *fileServiceImpl) Upload(request *model.FileRequest) (model.FileResponse, error) {
	filePath, err := validation.ValidateFileRequest(request)
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	return model.FileResponse{
		Path: strings.TrimPrefix(filePath, "storage/"),
		URL:  filePath,
	}, nil
}

func (service *fileServiceImpl) Get(fileType, fileName string) (string, error) {
	filePath := path.Join("storage", fileType, fileName)

	return helper.GetFile(filePath)
}

func (service *fileServiceImpl) Delete(fileType, fileName string) error {
	filePath := path.Join("storage", fileType, fileName)

	return helper.RemoveFile(filePath)
}

func ProvideFileService() FileService {
	return &fileServiceImpl{}
}
