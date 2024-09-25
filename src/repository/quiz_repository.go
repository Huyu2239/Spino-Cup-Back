package repository

import (
	"api/model"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IQuizRepository interface {
	GetFilteredQuizzes(quizzes *[]model.Quiz2, filters []model.Filter, limit int, random bool) error
	GetQuizByID(quiz *model.Quiz2, quizID uint) error
	CreateQuiz(quiz *model.Quiz2) error
	UpdateQuiz(quiz *model.Quiz2, quizID uint) error
	DeleteQuiz(quizID uint) error
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) IQuizRepository {
	return &quizRepository{db}
}

func (qr *quizRepository) GetFilteredQuizzes(quizzes *[]model.Quiz2, filters []model.Filter, limit int, random bool) error {

	db, err := applyFilters(qr.db, filters)

	if err != nil {
		return err
	}

	if limit <= 0 {
		limit = 10
	}

	if random {
		db = db.Order("RANDOM()")
	} else {
		db = db.Order("created_at")
	}

	if err := db.Limit(int(limit)).Find(quizzes).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) GetQuizByID(quiz *model.Quiz2, quizID uint) error {
	if err := qr.db.Where("id=?", quizID).First(&quiz).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) CreateQuiz(quiz *model.Quiz2) error {
	log.Printf("Quiz: %+v", quiz)
	if err := qr.db.Create(quiz).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) UpdateQuiz(quiz *model.Quiz2, quizID uint) error {

	result := qr.db.Model(quiz).Clauses(clause.Returning{}).
		Where("id=?", quizID).
		Updates(map[string]interface{}{
			"question":      quiz.Question,
			"code":          quiz.Code,
			"input_sample":  quiz.InputSample,
			"output_sample": quiz.OutputSample,
			"input_secret":  quiz.InputSecret,
			"output_secret": quiz.OutputSecret,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exit")
	}
	return nil

}

func (qr *quizRepository) DeleteQuiz(quizID uint) error {
	result := qr.db.Where("id=?", quizID).Delete(&model.Quiz2{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exit")
	}
	return nil
}
