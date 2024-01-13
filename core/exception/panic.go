package exception

import (
	"errors"
	"gorm.io/gorm"
)

func PanicIfNeeded(err error) {
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			panic(BadRequestError{
				Message: "Data already exists",
			})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(NotFoundError{
				Message: "Data not found",
			})
		} else {
			panic(err)
		}
	}
}
