package userhandler

import "github.com/labstack/echo/v4"

type UserHandler interface {
	SaveUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
