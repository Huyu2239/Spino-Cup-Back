package router

import (
	"api/controller"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(qc controller.IQuizController) *echo.Echo {
	e := echo.New()
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost", "http://spino.huyu2239.work"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	q := e.Group("/quizzes")
	q.GET("", qc.GetFilteredQuizzes)
	q.POST("", qc.CreateQuiz)
	q.PUT("/:quizID", qc.UpdateQuiz)
	q.DELETE("/:quizID", qc.DeleteQuiz)

	return e
}
