package response

import (
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
	Code    int  `json:"-"`
	Pretty  bool `json:"-"`
	Data    any  `json:"data,omitempty"`
	Message any  `json:"message"`
	Error   any  `json:"-"`
}

func MessageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field must be email"
	}
	return fe.Error()
}

func (a Response) JSONValidationError(ctx echo.Context) error {
	if err, ok := a.Error.(validator.ValidationErrors); ok && err != nil {
		var validationErrors []ValidationError

		for _, e := range err {
			validationErrors = append(validationErrors, ValidationError{
				Field:   e.Field(),
				Message: MessageForTag(e),
			})
		}

		a.Data = validationErrors
		a.Code = http.StatusBadRequest
		return a.JSON(ctx)
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
