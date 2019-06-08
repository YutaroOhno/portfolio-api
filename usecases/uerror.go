package usecases

import (
	"github.com/jinzhu/gorm"
)

type UError struct {
	Code int
	Msg string
}

const (
	BadRequest = iota + 1
	Unauthorized
	NotFound
	Conflict
	InternalException
)

func GetUErrorByError(err error) *UError {
	if err != nil {
		var code int
		if gorm.IsRecordNotFoundError(err) {
			code = NotFound
		} else {
			code = InternalException
		}

		return &UError{
			Code: code,
			Msg: err.Error(),
		}
	}

	return nil
}