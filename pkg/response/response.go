package response

import (
	"errors"
	appErrors "github.com/agung96tm/miblog/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

type Response struct {
	Code    int   `json:"-"`
	Pretty  bool  `json:"-"`
	Data    any   `json:"data,omitempty"`
	Message any   `json:"message"`
	Error   error `json:"-"`
}

func MessageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field must be email"
	case "min":
		return "This field has min length"
	case "max":
		return "This field has max length"
	}
	return fe.Tag()
}

func (a Response) JSONValidationError(ctx echo.Context) error {
	a.Code = http.StatusBadRequest

	if err, ok := a.Error.(validator.ValidationErrors); ok && err != nil {
		var validationErrors []ValidationError

		for _, e := range err {
			validationErrors = append(validationErrors, ValidationError{
				Field:   e.Field(),
				Message: MessageForTag(e),
			})
		}

		a.Data = validationErrors
	} else {
		a.Message = a.Error.Error()
	}

	return a.JSON(ctx)
}

func (a Response) JSONPolicyError(ctx echo.Context) error {
	a.Message = a.Error.Error()
	if errors.Is(a.Error, appErrors.ErrPolicyUnauthorized) {
		a.Code = http.StatusUnauthorized
	}
	if errors.Is(a.Error, appErrors.ErrPolicyForbidden) {
		a.Code = http.StatusForbidden
	}

	return a.JSON(ctx)
}

func (a Response) JSON(ctx echo.Context) error {
	if a.Message == "" || a.Message == nil {
		a.Message = http.StatusText(a.Code)
	}

	if err, ok := a.Message.(error); ok {
		a.Message = err.Error()
	}

	if a.Pretty {
		return ctx.JSONPretty(a.Code, a, "\t")
	}

	return ctx.JSON(a.Code, a)
}
