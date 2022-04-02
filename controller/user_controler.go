package controller

import (
	"gorm-api/config"
	m "gorm-api/middleware"
	"gorm-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserReq struct {
	ID       int    `json:"id" param:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
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
	var user []models.User
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	db := config.DBManager()
	db = db.Find(&user, req.ID)

	return c.JSON(http.StatusOK, user)
}

func UpdateUserById(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := models.User{}

	newUser := models.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	}

	db := config.DBManager()
	db = db.Model(&user).Where("id = ?", req.ID).Updates(newUser)
	return c.JSON(http.StatusCreated, user)
}

func DeleteUserById(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var user []models.User
	db := config.DBManager()
	db = db.Delete(&user, req.ID)

	result := map[string]string{
		"response_code": "200",
		"message":       "succsess",
	}

	return c.JSON(http.StatusOK, result)
}

func LoginUser(c echo.Context) (err error) {

	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := models.User{}
	db := config.DBManager()

	db = db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user)

	token, err := m.CreateToken(req.ID, req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Sucsess",
		"user":    user,
		"token":   token,
	})
}
