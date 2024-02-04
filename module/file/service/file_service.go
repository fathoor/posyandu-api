package service

import "github.com/fathoor/posyandu-api/module/file/model"

type FileService interface {
	Upload(request *model.FileRequest) (model.FileResponse, error)
	Get(fileType, fileName string) (string, error)
	Delete(fileType, fileName string) error
}
