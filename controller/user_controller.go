package controller

import (
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
	return nil
}

func (uc *userController) LogIn(c echo.Context) error {
	return nil
}

func (uc *userController) LogOut(c echo.Context) error {
	return nil
}