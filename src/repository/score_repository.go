package repository

import (
	"api/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type IScoreRepository interface {
	GetFilteredScores(scores *[]model.Score, filters []model.Filter) error
	CreateScore(score *model.Score) error
	DeleteScore(scoreId uint) error
}

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) IScoreRepository {
	return &scoreRepository{db}
}

func (sr *scoreRepository) GetFilteredScores(scores *[]model.Score, filters []model.Filter) error {

	db, err := applyFilters(sr.db, filters)

	if err != nil {
		return err
	}

	if err := db.Order("created_at").Find(scores).Error; err != nil {
		return err
	}
	return nil
}

func (sr *scoreRepository) CreateScore(score *model.Score) error {
	log.Printf("Score: %+v", score)
	if err := sr.db.Create(score).Error; err != nil {
		return err
	}
	return nil
}

func (sr *scoreRepository) DeleteScore(scoreId uint) error {
	result := sr.db.Where("id=?", scoreId).Delete(&model.Score{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exit")
	}
	return nil
}
