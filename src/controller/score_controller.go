package controller

import (
	"api/model"
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IScoreController interface {
	GetFilteredScores(c echo.Context) error
	CreateScore(c echo.Context) error
	DeleteScore(c echo.Context) error
}

type scoreController struct {
	su usecase.IScoreUsecase
}

func NewScoreContoller(su usecase.IScoreUsecase) IScoreController {
	return &scoreController{su}
}

func (sc *scoreController) GetFilteredScores(c echo.Context) error {
	queryFilters := c.QueryParam("filters")

	filters, err := parseFilters(queryFilters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	scoreRes, err := sc.su.GetFilteredScores(filters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, scoreRes)
}

func (sc *scoreController) CreateScore(c echo.Context) error {
	score := model.Score{}
	if err := c.Bind(&score); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	score.UserID = c.Get("user").(model.UserResponse).ID

	scoreRes, err := sc.su.CreateScore(score)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, scoreRes)
}

func (sc *scoreController) DeleteScore(c echo.Context) error {
	id := c.Param("scoreID")
	scoreID, _ := strconv.Atoi(id)
	err := sc.su.DeleteScore(uint(scoreID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
