package validation

import (
	"fmt"
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/helper"
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/file/model"
	"path"
)

func ValidateFileRequest(request *model.FileRequest) (string, error) {
	valid := validation.Validator.Struct(request)
	if valid != nil {
		return "", exception.BadRequestError{
			Message: "Invalid request data",
		}
	}

	switch request.Type {
	case "image":
		ext := path.Ext(request.File.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return "", exception.BadRequestError{
				Message: fmt.Sprintf("Invalid file extension: %s is not supported", ext),
			}
		}

		if request.File.Size > 2*1024*1024 {
			return "", exception.BadRequestError{
				Message: "File size is too large, maximum is 2MB",
			}
		}
	default:
		return "", exception.BadRequestError{
			Message: fmt.Sprintf("Invalid file type: %s is not supported", request.Type),
		}
	}

	var (
		fileType = request.Type
		fileExt  = path.Ext(request.File.Filename)
		filePath = helper.GenerateFilePath(fileType, fileExt)
	)

	return filePath, nil
}
