package validation

import (
	"github.com/itsLeonB/posyandu-api/core/validation"
	"github.com/itsLeonB/posyandu-api/module/chat/model"
)

func ValidateChatCreateRequest(request *model.ChatCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateChatRoomCreateRequest(request *model.ChatRoomCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateChatUpdateRequest(request *model.ChatUpdateRequest) error {
	return validation.Validator.Struct(request)
}
