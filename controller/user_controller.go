package controller

import (
	"net/http"

	"github.com/DaigoSugiyama0317/map-memo-app/model"
	"github.com/DaigoSugiyama0317/map-memo-app/usecase"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	//リクエストをUser構造体に入れる。
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//User新規登録
	resUser, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resUser)
}

func (uc *userController) LogIn(c echo.Context) error {
	return nil
}

func (uc *userController) LogOut(c echo.Context) error {
	return nil
}