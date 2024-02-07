package lib

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"reflect"
	"strings"
)

type HttpHandler struct {
	Engine *echo.Echo
}

type Validator struct {
	validate *validator.Validate
}

func (a *Validator) Validate(i interface{}) error {
	return a.validate.Struct(i)
}

func NewHttpHandler() HttpHandler {
	engine := echo.New()
	engine.HidePort = true
	engine.HideBanner = true
	engine.Binder = &BinderWithValidation{}

	httpHandler := HttpHandler{Engine: engine}

	// validator
	httpHandler.Engine.Validator = func() echo.Validator {
		v := validator.New()
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		return &Validator{validate: v}
	}()

	return httpHandler
}

type BinderWithValidation struct{}

func (BinderWithValidation) Bind(i interface{}, ctx echo.Context) error {
	binder := &echo.DefaultBinder{}

	if err := binder.Bind(i, ctx); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}
