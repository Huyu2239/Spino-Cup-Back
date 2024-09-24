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
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	if err := middleware.InitFirebase(); err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.AuthMiddleware(uu))

	q := e.Group("/quizzes")
	q.GET("", qc.GetFilteredQuizzes)
	q.POST("", qc.CreateQuiz)
	q.PUT("/:quizID", qc.UpdateQuiz)
	q.DELETE("/:quizID", qc.DeleteQuiz)
	q.GET("/ans/:quizID", qc.GetQuizAnswer)
	q.GET("/check/:quizID", qc.CheckQuiz)

	s := e.Group("/scores")
	s.GET("", sc.GetFilteredScores)
	s.POST("", sc.CreateScore)
	s.DELETE("/:scoreID", sc.DeleteScore)

	return e
}
