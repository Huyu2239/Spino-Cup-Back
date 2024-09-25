package repository

import (
	"api/model"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IQuizRepository interface {
	GetFilteredQuizzes(quizzes *[]model.Quiz, filters []model.Filter, limit int, random bool) error
	GetAllLanguages(languages *[]string) error
	GetQuizByID(quiz *model.Quiz, quizID uint) error
	CreateQuiz(quiz *model.Quiz) error
	UpdateQuiz(quiz *model.Quiz, quizID uint) error
	DeleteQuiz(quizID uint) error
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) IQuizRepository {
	return &quizRepository{db}
}

func (qr *quizRepository) GetFilteredQuizzes(quizzes *[]model.Quiz, filters []model.Filter, limit int, random bool) error {

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

func (qr *quizRepository) GetQuizByID(quiz *model.Quiz, quizID uint) error {
	if err := qr.db.Where("id=?", quizID).First(&quiz).Error; err != nil {
		return err
	}
	return nil
}

func (qr *quizRepository) GetAllLanguages(languages *[]string) error {
	if err := qr.db.Model(&model.Quiz{}).Distinct("language").Pluck("language", &languages).Error; err != nil {
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

func (qr *quizRepository) UpdateQuiz(quiz *model.Quiz, quizID uint) error {

	result := qr.db.Model(quiz).Clauses(clause.Returning{}).
		Where("id=?", quizID).
		Updates(map[string]interface{}{
			"question":    quiz.Question,
			"answer_x":    quiz.AnswerX,
			"answer_y":    quiz.AnswerY,
			"edited_text": quiz.EditedText,
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

func (qr *quizRepository) DeleteQuiz(quizID uint) error {
	result := qr.db.Where("id=?", quizID).Delete(&model.Quiz{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exit")
	}
	return nil
}
