package app

import (
	roleservice "users-service/internal/domain/service/role_service"
	userservice "users-service/internal/domain/service/user_service"
	"users-service/internal/infra/db"
	rolerepository "users-service/internal/infra/role_repository"
	userrepository "users-service/internal/infra/user_repository"
	rolehandler "users-service/internal/interfaces/role_handler"
	userhandler "users-service/internal/interfaces/user_handler"

	"github.com/labstack/echo/v4"
)

func InitializedServer() *echo.Echo {
	database := db.DBConnection()

	// user
	userRepo := userrepository.NewUserRepository(database)
	userService := userservice.NewUserService(userRepo)
	userHandler := userhandler.NewUserHandler(userService)

	// role
	roleRepo := rolerepository.NewRoleRepository(database)
	roleService := roleservice.NewRoleService(roleRepo)
	roleHandler := rolehandler.NewRoleHandler(roleService)

	r := echo.New()
	r.POST("/registrasi", userHandler.SaveUser)
	r.GET("/user/:id", userHandler.GetUser)
	r.GET("/users", userHandler.GetUsers)
	r.PUT("/user/update/:id", userHandler.UpdateUser)
	r.DELETE("/user/delete/:id", userHandler.DeleteUser)

	r.POST("/registrasi/role", roleHandler.SaveRole)
	r.GET("/role/:id", roleHandler.GetRole)
	r.GET("/roles", roleHandler.GetRoleList)

	return r
}
