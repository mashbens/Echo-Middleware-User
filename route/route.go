package route

import (
	"gorm-api/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/user/list", controller.GetUsers)
	e.POST("/user", controller.CreateUser)
	e.GET("/user/:id", controller.GetUserById)
	// e.PUT("/users/:id", controller.UpdateUserById)
	// e.DELETE("/users/:id", controller.DeleteUserById)
	return e
}
