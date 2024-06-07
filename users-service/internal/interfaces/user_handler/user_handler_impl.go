package userhandler

import (
	"net/http"
	"strconv"
	userservice "users-service/internal/domain/service/user_service"
	"users-service/internal/interfaces/api/request"
	"users-service/internal/interfaces/api/response"

	"github.com/labstack/echo/v4"
)

type UserHandlerImpl struct {
	service userservice.UserService
}

func NewUserHandler(service userservice.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		service: service,
	}
}

func (controller *UserHandlerImpl) SaveUser(c echo.Context) error {
	user := new(request.UserServiceRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	// if err := c.Validate(user); err != nil {
	// 	return err
	// }

	// masukan data setelah di validasi
	saveUser, errSaveUser := controller.service.SaveUser(*user)

	if errSaveUser != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, errSaveUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "Berhasil membuat data", saveUser))
}

func (controller *UserHandlerImpl) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getUser, errGetUser := controller.service.GetUser(id)

	if errGetUser != nil {
		return c.JSON(http.StatusNotFound, response.ResponseToClient(http.StatusNotFound, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "success", getUser))
}

func (controller *UserHandlerImpl) GetUsers(c echo.Context) error {
	getUsers, errGetUsers := controller.service.GetUsers()

	if errGetUsers != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseToClient(http.StatusBadRequest, errGetUsers.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseToClient(http.StatusOK, "success", getUsers))
}
