package repository

import (
	"api/model"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IQuizRepository interface {
	GetFilteredQuizzes(quizzes *[]model.Quiz, filters []model.Filter) error
	GetQuizByID(quiz *model.Quiz, quizId uint) error
	CreateQuiz(quiz *model.Quiz) error
	UpdateQuiz(quiz *model.Quiz, quizId uint) error
	DeleteQuiz(quizId uint) error
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) IQuizRepository {
	return &quizRepository{db}
}

func (qr *quizRepository) GetFilteredQuizzes(quizzes *[]model.Quiz, filters []model.Filter) error {

	dbUser := qr.db.Joins("User")
	db, err := applyFilters(dbUser, filters)

	if err != nil {
		return err
	}

	if err := db.Order("created_at").Find(quizzes).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) GetQuizByID(quiz *model.Quiz, quizId uint) error {
	if err := qr.db.Where("id=?", quizId).First(&quiz).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) CreateQuiz(quiz *model.Quiz) error {
	log.Printf("Quiz: %+v", quiz)
	if err := qr.db.Create(quiz).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) UpdateQuiz(quiz *model.Quiz, quizId uint) error {

	result := qr.db.Model(quiz).Clauses(clause.Returning{}).
		Where("id=?", quizId).
		Updates(map[string]interface{}{
			"question":    quiz.Question,
			"answer_x":    quiz.AnswerX,
			"answer_y":    quiz.AnswerY,
			"difficulty":  quiz.Difficulty,
			"language":    quiz.Language,
			"explanation": quiz.Explanation,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exit")
	}
	return nil

}

func (qr *quizRepository) DeleteQuiz(quizId uint) error {
	result := qr.db.Where("id=?", quizId).Delete(&model.Quiz{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exit")
	}
	return nil
}
