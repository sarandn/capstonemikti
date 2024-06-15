package rolehandler

import (
	"net/http"
	"strconv"
	roleservice "users-service/internal/domain/service/role_service"
	"users-service/internal/interfaces/api/request"
	"users-service/internal/interfaces/api/response"

	"github.com/labstack/echo/v4"
)

type RoleHandlerImpl struct {
	roleService roleservice.RoleService
}

func NewRoleHandler(service roleservice.RoleService) *RoleHandlerImpl {
	return &RoleHandlerImpl{
		roleService: service,
	}
}

func (controller *RoleHandlerImpl) SaveRole(c echo.Context) error {
	role := new(request.RoleServiceRequest)

	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	// if err := c.Validate(role); err != nil {
	// 	return err
	// }

	saveRole, errSaveRole := controller.roleService.SaveRole(*role)

	if errSaveRole != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, errSaveRole.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "berhasil membuat Role", saveRole))
}

func (controller *RoleHandlerImpl) GetRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getRole, errGetRole := controller.roleService.GetRole(id)

	if errGetRole != nil {
		return c.JSON(http.StatusNotFound, response.ResponseToClient(http.StatusBadRequest, errGetRole.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "success", getRole))
}

func (controller *RoleHandlerImpl) GetRoleList(c echo.Context) error {
	getRole, errGetRole := controller.roleService.GetRoles()

	if errGetRole != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseToClient(http.StatusInternalServerError, errGetRole.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "success", getRole))
}

func (controller *RoleHandlerImpl) UpdateRole(c echo.Context) error {
	role := new(request.RoleUpdateServiceRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	roleUpdate, errRoleUpdate := controller.roleService.UpdateRole(*role, id)

	if errRoleUpdate != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, errRoleUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "data berhasil diupdate", roleUpdate))
}

func (controller *RoleHandlerImpl) DeleteRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delRole, errDelRole := controller.roleService.DeleteData(id)

	if errDelRole != nil {
		return c.JSON(http.StatusNotFound, response.ResponseToClient(http.StatusNotFound, errDelRole.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "Berhasil delete user", delRole))
}
