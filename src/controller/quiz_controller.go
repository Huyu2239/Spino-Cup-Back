package controller

import (
	"api/model"
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IQuizController interface {
	GetFilteredQuizzes(c echo.Context) error
	CreateQuiz(c echo.Context) error
	UpdateQuiz(c echo.Context) error
	DeleteQuiz(c echo.Context) error
	GetQuizAnswer(c echo.Context) error
	CheckQuiz(c echo.Context) error
}

type quizController struct {
	qu usecase.IQuizUsecase
}

func NewQuizContoller(qu usecase.IQuizUsecase) IQuizController {
	return &quizController{qu}
}

func (qc *quizController) GetFilteredQuizzes(c echo.Context) error {
	queryFilters := c.QueryParam("filters")

	filters, err := parseFilters(queryFilters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	quizRes, err := qc.qu.GetFilteredQuizzes(filters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, quizRes)
}

func (qc *quizController) CreateQuiz(c echo.Context) error {
	quiz := model.Quiz{}
	if err := c.Bind(&quiz); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	quiz.UserId = c.Get("user").(model.UserResponse).ID

	quizRes, err := qc.qu.CreateQuiz(quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, quizRes)
}

func (qc *quizController) UpdateQuiz(c echo.Context) error {
	quiz := model.Quiz{}
	if err := c.Bind(&quiz); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	quiz.UserId = c.Get("user").(model.UserResponse).ID

	id := c.Param("quizId")
	quizId, _ := strconv.Atoi(id)
	quizRes, err := qc.qu.UpdateQuiz(quiz, uint(quizId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, quizRes)
}

func (qc *quizController) DeleteQuiz(c echo.Context) error {
	id := c.Param("quizId")
	quizId, _ := strconv.Atoi(id)
	err := qc.qu.DeleteQuiz(uint(quizId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (qc *quizController) GetQuizAnswer(c echo.Context) error {
	id := c.Param("quizId")
	quizId, _ := strconv.Atoi(id)
	ansRes, err := qc.qu.GetQuizAnswer(uint(quizId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ansRes)
}

func (qc *quizController) CheckQuiz(c echo.Context) error {
	id := c.Param("quizId")
	quizId, _ := strconv.Atoi(id)

	x := c.QueryParam("x")
	y := c.QueryParam("y")
	ansX, _ := strconv.Atoi(x)
	ansY, _ := strconv.Atoi(y)

	checkAnsRes, err := qc.qu.CheckQuiz(uint(quizId), uint(ansX), uint(ansY))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, checkAnsRes)
}
