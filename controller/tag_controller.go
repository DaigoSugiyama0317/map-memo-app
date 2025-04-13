package controller

import (
	"net/http"

	"github.com/DaigoSugiyama0317/map-memo-app/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITagController interface {
	GetAllTags(c echo.Context) error
	CreateTag(c echo.Context) error
	UpdateTag(c echo.Context) error
	DeleteTag(c echo.Context) error
}

type tagController struct {
	tu usecase.ITagUsecase
}

func NewTagController(tu usecase.ITagUsecase) ITagController {
	return &tagController{tu}
}

func (tc tagController) GetAllTags(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	tagRes, err := tc.tu.GetAllTags(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tagRes)
}

func (tc tagController) CreateTag(c echo.Context) error {

}

func (tc tagController) UpdateTag(c echo.Context) error {

}

func (tc tagController) DeleteTag(c echo.Context) error {

}