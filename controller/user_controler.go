package controller

import (
	"gorm-api/config"
	"gorm-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserReq struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var UserList = make(map[int]models.User, 0)

func CreateUser(c echo.Context) error {
	req := new(UserReq)
	c.Bind(req)
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"data":    user,
	})
}

func GetUsers(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func GetUserById(c echo.Context) error {
	var user []models.User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Find(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"data":    user,
	})
}

// func UpdateUserById(c echo.Context) error {

// }

func DeleteUserById(c echo.Context) error {
	req := new(UserReq)
	c.Bind(req)
	var user []models.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Delete(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succsess delete user",
	})

}
