package admin

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/error/code"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
)

const (
	_unknownError                = "UNKNOWN"
	_useCaseInputValidationError = "USECASE_INPUT_VALIDATION_ERROR"
)

func getCodeFromError(err error) string {
	haveCodeErr, ok := err.(interface{ Code() string })
	if !ok {
		return _unknownError
	}
	return haveCodeErr.Code()
}

func getStatusCodeFromCode(c string) int {
	switch c {
	case code.PermissionDeniedError:
		return iris.StatusForbidden
	case code.AuthenticationError:
		return iris.StatusUnauthorized
	case code.InternalServerError, _unknownError:
		return iris.StatusInternalServerError
	case code.ValidationError, _useCaseInputValidationError:
		return iris.StatusBadRequest
	case code.NotFoundError:
		return iris.StatusNoContent
	case code.DBError:
		return iris.StatusNotAcceptable
	default:
		return iris.StatusInternalServerError
	}
}

func logError(err error, c string, l logger.Interface) {
	switch c {
	case code.PermissionDeniedError,
		code.AuthenticationError,
		code.ValidationError,
		code.NotFoundError,
		_useCaseInputValidationError:
		l.Info(err.Error())
	case code.DBError:
		l.Warn(err.Error())
	case code.InternalServerError, _unknownError:
		l.Error(err)
	default:
		l.Info(err.Error())
	}
}

func getTranslateFunc(tr func(string, ...any) string) model.TranslateFunc {
	return func(m *i18n.Message, a ...any) string {
		return tr(m.ID, a...)
	}
}

func handleError(ctx iris.Context, err error, l logger.Interface) {
	code := getCodeFromError(err)
	statusCode := getStatusCodeFromCode(code)
	message := ""
	detail := ""

	switch foundedError := err.(type) {
	case model.TranslatableError:
		unwrapErr := foundedError.Unwrap()
		logError(unwrapErr, code, l)
		translatableErr := foundedError.SetTranslateFunc(getTranslateFunc(ctx.Tr))
		message = translatableErr.Error()
		detail = unwrapErr.Error()
	case *model.TranslatableError:
		unwrapErr := foundedError.Unwrap()
		logError(unwrapErr, code, l)
		translatableErr := foundedError.SetTranslateFunc(getTranslateFunc(ctx.Tr))
		message = translatableErr.Error()
		detail = unwrapErr.Error()
	default:
		logError(err, code, l)
		message = foundedError.Error()
		detail = message
	}

	ctx.StopWithJSON(statusCode, errorResponse{
		Code:    code,
		Message: message,
		Detail:  detail,
	})
}

// alias handleError.
func HandleError(ctx iris.Context, err error, l logger.Interface) {
	handleError(ctx, err, l)
}

func translatableErrorFromValidationErrors(
	inputStructure any, errs *validator.ValidationErrors,
) *model.TranslatableError {
	verboser, ok := inputStructure.(interface {
		GetErrorMessageFromStructField(error) *i18n.Message
	})
	i18nMessage := _oneOrMoreFieldsFailedToBeValidatedMessage
	if ok {
		for _, validationErr := range *errs {
			i18nMessage = verboser.GetErrorMessageFromStructField(validationErr)
			break
		}
	}
	return model.NewTranslatableError(errs, i18nMessage, _useCaseInputValidationError, nil)
}

func handleBindingError(
	ctx iris.Context,
	err error,
	l logger.Interface,
	input any,
	wrapTranslateError func(error) error,
) {
	ae := err
	switch actualErr := err.(type) {
	case validator.ValidationErrors:
		ae = translatableErrorFromValidationErrors(input, &actualErr)
	case *validator.ValidationErrors:
		ae = translatableErrorFromValidationErrors(input, actualErr)
	case *json.UnmarshalTypeError:
		ae = model.NewTranslatableError(actualErr, _oneOrMoreFieldsFailedToBeValidatedMessage, _useCaseInputValidationError, nil)
	default:
		if wrapTranslateError != nil {
			ae = wrapTranslateError(err)
		}
	}
	handleError(ctx, ae, l)
}
