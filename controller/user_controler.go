package controller

import (
	"fmt"
	"gorm-api/config"
	"gorm-api/models"
	"net/http"

	"github.com/labstack/echo"
)

type UserReq struct {
	ID    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(c echo.Context) error {
	var user []models.User

	db := config.DBManager()
	db = db.Find(&user)
	return c.JSON(http.StatusOK, user)

}

func CreateUser(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	db := config.DBManager()
	db = db.Create(&user)

	return c.JSON(http.StatusOK, user)
}

func GetUserById(c echo.Context) (err error) {
	req := new(UserReq)
	c.Bind(req)

	if err = c.Bind(req); err != nil {
		fmt.Println("---->", c.Bind(req))
		fmt.Println("---->", req)
		fmt.Print(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	fmt.Print(err)

	var user []models.User

	db := config.DBManager()
	db = db.First(&user, req.ID)

	// fmt.Println("====> ", req.ID)

	return c.JSON(http.StatusOK, user)
}

// func UpdateUserById(c echo.Context) error {

// }

// func DeleteUserById(c echo.Context) error {

// }
