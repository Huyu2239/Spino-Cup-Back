package usecase

import (
	"api/model"
	"api/repository"
)

type IScoreUsecase interface {
	GetFilteredScores(filters []model.Filter) ([]model.ScoreResponse, error)
	CreateScore(score model.Score) (model.ScoreResponse, error)
	DeleteScore(scoreID uint) error
}

type scoreUsecase struct {
	sr repository.IScoreRepository
}

func NewScoreUsecase(sr repository.IScoreRepository) IScoreUsecase {
	return &scoreUsecase{sr}
}

func (su *scoreUsecase) GetFilteredScores(filters []model.Filter) ([]model.ScoreResponse, error) {
	scores := []model.Score{}
	if err := su.sr.GetFilteredScores(&scores, filters); err != nil {
		return nil, err
	}
	resScores := []model.ScoreResponse{}
	for _, v := range scores {
		s := model.ScoreResponse{
			ID:    v.ID,
			Point: v.Point,
		}
		resScores = append(resScores, s)
	}
	return resScores, nil
}

func (su *scoreUsecase) CreateScore(score model.Score) (model.ScoreResponse, error) {
	if err := su.sr.CreateScore(&score); err != nil {
		return model.ScoreResponse{}, err
	}
	resScore := model.ScoreResponse{
		ID:    score.ID,
		Point: score.Point,
	}
	return resScore, nil
}

func (su *scoreUsecase) DeleteScore(scoreID uint) error {
	if err := su.sr.DeleteScore(scoreID); err != nil {
		return err
	}
	return nil
}
