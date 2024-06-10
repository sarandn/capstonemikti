package app

import (
	"net/http"
	"os"
	"users-service/config"
	roleservice "users-service/internal/domain/service/role_service"
	userservice "users-service/internal/domain/service/user_service"
	"users-service/internal/infra/db"
	rolerepository "users-service/internal/infra/role_repository"
	userrepository "users-service/internal/infra/user_repository"
	"users-service/internal/interfaces/api/response"
	rolehandler "users-service/internal/interfaces/role_handler"
	userhandler "users-service/internal/interfaces/user_handler"
	"users-service/pkg/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitializedServer() *echo.Echo {
	database := db.DBConnection()

	// user
	userRepo := userrepository.NewUserRepository(database)
	userToken := helper.NewTokenUseCase()
	userService := userservice.NewUserService(userRepo, userToken)
	userHandler := userhandler.NewUserHandler(userService)

	// role
	roleRepo := rolerepository.NewRoleRepository(database)
	roleService := roleservice.NewRoleService(roleRepo)
	roleHandler := rolehandler.NewRoleHandler(roleService)

	config.LoadConfig()

	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = helper.BindAndValidate

	r.POST("/registrasi", userHandler.SaveUser)
	r.GET("/user/:id", userHandler.GetUser, JWTProtection())
	r.GET("/users", userHandler.GetUsers, JWTProtection())
	r.PUT("/user/update/:id", userHandler.UpdateUser, JWTProtection())
	r.DELETE("/user/delete/:id", userHandler.DeleteUser)
	r.POST("/user/login", userHandler.LoginUser)

	r.POST("/registrasi/role", roleHandler.SaveRole)
	r.GET("/role/:id", roleHandler.GetRole, JWTProtection())
	r.GET("/roles", roleHandler.GetRoleList, JWTProtection())

	return r
}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, response.ResponseToClient(http.StatusUnauthorized, "anda harus login untuk mengakses resource ini !", nil))
		},
	})
}
