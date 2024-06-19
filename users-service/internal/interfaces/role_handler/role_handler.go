package rolehandler

import "github.com/labstack/echo/v4"

type RoleHandler interface {
	SaveRole(c echo.Context) error
	GetRole(c echo.Context) error
	GetRoleList(c echo.Context) error
	UpdateRole(c echo.Context) error
	DeleteRole(c echo.Context) error
}
