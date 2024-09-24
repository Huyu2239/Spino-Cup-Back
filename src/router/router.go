package router

import (
	"api/controller"
	"api/middleware"
	"api/usecase"
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(qc controller.IQuizController, sc controller.IScoreController, uu usecase.IUserUsecase) *echo.Echo {
	e := echo.New()

	if err := middleware.InitFirebase(); err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.AuthMiddleware(uu))

	q := e.Group("/quizzes")
	q.GET("", qc.GetFilteredQuizzes)
	q.POST("", qc.CreateQuiz)
	q.PUT("/:quizId", qc.UpdateQuiz)
	q.DELETE("/:quizId", qc.DeleteQuiz)
	q.GET("/ans/:quizId", qc.GetQuizAnswer)
	q.GET("/check/:quizId", qc.CheckQuiz)

	s := e.Group("/scores")
	s.GET("", sc.GetFilteredScores)
	s.POST("", sc.CreateScore)
	s.DELETE("/:scoreId", sc.DeleteScore)

	return e
}
