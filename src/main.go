package main

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/router"
	"api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	quizRepository := repository.NewQuizRepository(db)
	scoreRepository := repository.NewScoreRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	quizUsecase := usecase.NewQuizUsecase(quizRepository)
	scoreUsecase := usecase.NewScoreUsecase(scoreRepository)
	quizController := controller.NewQuizContoller(quizUsecase)
	scoreController := controller.NewScoreContoller(scoreUsecase)
	e := router.NewRouter(quizController, scoreController, userUsecase)
	e.Logger.Fatal(e.Start(":1323"))
}
