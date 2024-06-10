package helper

import (
	"fmt"
	"net/http"
	"users-service/internal/interfaces/api/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidate(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s field ini wajib diisi", err.Field())
				report.Code = http.StatusBadRequest
			case "email":
				report.Message = fmt.Sprintf("%s ini bukan email valid", err.Field())
				report.Code = http.StatusBadRequest
			case "gte":
				report.Message = fmt.Sprintf("%s nomor harus lebih dari 10", err.Field())
				report.Code = http.StatusBadRequest
			case "lte":
				report.Message = fmt.Sprintf("%s nomor harus kurang dari 15", err.Field())
				report.Code = http.StatusBadRequest
			case "numeric":
				report.Message = fmt.Sprintf("%s harus berbentuk nomor", err.Field())
				report.Code = http.StatusBadRequest

			}
		}
	}

	c.Logger().Error(report.Message)
	c.JSON(report.Code, response.ResponseToClient(report.Code, report.Message.(string), nil))
}
