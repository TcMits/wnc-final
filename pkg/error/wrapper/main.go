package wrapper

import "github.com/TcMits/wnc-final/pkg/error/code"

type (
	wrapperError struct {
		Msg error
	}
	ValidationError struct {
		*wrapperError
	}
	NotFoundError struct {
		*wrapperError
	}
	DBError struct {
		*wrapperError
	}
)

func (v *wrapperError) Error() string {
	return v.Msg.Error()
}

func (v *wrapperError) Unwrap() error {
	return v.Msg
}

func (v *ValidationError) Code() string {
	return code.ValidationError
}

func (v *NotFoundError) Code() string {
	return code.NotFoundError
}

func (v *DBError) Code() string {
	return code.DBError
}

func NewValidationError(err error) *ValidationError {
	return &ValidationError{
		wrapperError: &wrapperError{
			Msg: err,
		},
	}
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{
		wrapperError: &wrapperError{
			Msg: err,
		},
	}
}

func NewDBError(err error) *DBError {
	return &DBError{
		wrapperError: &wrapperError{
			Msg: err,
		},
	}
}
