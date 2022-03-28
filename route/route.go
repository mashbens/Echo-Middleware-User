package route

import (
	"gorm-api/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.CreateUser)
	e.GET("/users/:id", controller.GetUserById)
	// e.PUT("/users/:id", controller.UpdateUserById)
	e.DELETE("/users/:id", controller.DeleteUserById)
	return e
}
