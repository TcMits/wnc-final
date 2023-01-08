package usecase

import (
	"github.com/TcMits/wnc-final/pkg/error/wrapper"
)

const DEBUG = true

func WrapError(err error) error {
	if err != nil && !DEBUG {
		if _, ok := err.(interface{ Code() string }); !ok {
			return wrapper.NewDBError(err)
		}
	}
	return err
}

func ValidationError(err error) error {
	return WrapError(wrapper.NewValidationError(err))
}
