package route

import (
	"gorm-api/constants"
	"gorm-api/controller"
	m "gorm-api/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("/user", controller.CreateUser)
	e.POST("/login", controller.LoginUser)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/user/list", controller.GetUsers)
	eJwt.GET("/user/:id", controller.GetUserById)
	eJwt.PUT("/user/:id", controller.UpdateUserById)
	eJwt.DELETE("/user/:id", controller.DeleteUserById)
	return e

}
