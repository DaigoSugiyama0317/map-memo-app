package main

import (
	"github.com/DaigoSugiyama0317/map-memo-app/controller"
	"github.com/DaigoSugiyama0317/map-memo-app/db"
	"github.com/DaigoSugiyama0317/map-memo-app/repository"
	"github.com/DaigoSugiyama0317/map-memo-app/router"
	"github.com/DaigoSugiyama0317/map-memo-app/usecase"
)

func main() {
	db := db.NewDB()
	//User Repositoryの初期化
	userRepository := repository.NewUserRepository(db)
	//User Usecaseの初期化
	userUsecase := usecase.NewUserUsecase(userRepository)
	////User controllerの初期化
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}