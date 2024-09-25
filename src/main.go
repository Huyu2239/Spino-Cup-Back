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
	quizRepository := repository.NewQuizRepository(db)
	quizUsecase := usecase.NewQuizUsecase(quizRepository)
	quizController := controller.NewQuizContoller(quizUsecase)
	e := router.NewRouter(quizController)
	e.Logger.Fatal(e.Start(":1323"))
}
