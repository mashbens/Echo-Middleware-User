package route

import (
	"gorm-api/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/user/list", controller.GetUsers)
	e.POST("/user", controller.CreateUser)
	e.GET("/user/:id", controller.GetUserById)
	e.PUT("/user/:id", controller.UpdateUserById)
	e.DELETE("/user/:id", controller.DeleteUserById)
	return e

}
