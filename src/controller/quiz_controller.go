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
}

type quizController struct {
	qu usecase.IQuizUsecase
}

func NewQuizContoller(qu usecase.IQuizUsecase) IQuizController {
	return &quizController{qu}
}

func (qc *quizController) GetFilteredQuizzes(c echo.Context) error {
	queryFilters := c.QueryParam("filters")
	queryLimit := c.QueryParam("limit")
	queryRandom := c.QueryParam("random")

	filters, err := parseFilters(queryFilters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	limit, _ := strconv.Atoi(queryLimit)
	randomInt, _ := strconv.Atoi(queryRandom)
	random := false

	if randomInt == 1 {
		random = true
	}

	quizRes, err := qc.qu.GetFilteredQuizzes(filters, limit, random)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, quizRes)
}

func (qc *quizController) CreateQuiz(c echo.Context) error {
	quiz := model.Quiz2{}
	if err := c.Bind(&quiz); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	quizRes, err := qc.qu.CreateQuiz(quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, quizRes)
}

func (qc *quizController) UpdateQuiz(c echo.Context) error {
	quiz := model.Quiz2{}
	if err := c.Bind(&quiz); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("quizID")
	quizID, _ := strconv.Atoi(id)
	quizRes, err := qc.qu.UpdateQuiz(quiz, uint(quizID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, quizRes)
}

func (qc *quizController) DeleteQuiz(c echo.Context) error {
	id := c.Param("quizID")
	quizID, _ := strconv.Atoi(id)
	err := qc.qu.DeleteQuiz(uint(quizID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
