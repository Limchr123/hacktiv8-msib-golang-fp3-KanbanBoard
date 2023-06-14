package helpers

import (
	"github.com/asaskevich/govalidator"
	"kanban_board/pkg/errs"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest("Error occurred while trying to validate data")
	}

	return nil
}
