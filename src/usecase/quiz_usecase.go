package usecase

import (
	"api/model"
	"api/repository"
)

type IQuizUsecase interface {
	GetFilteredQuizzes(filters []model.Filter, limit int, random bool) ([]model.QuizResponse, error)
	CreateQuiz(quiz model.Quiz2) (model.QuizResponse, error)
	UpdateQuiz(quiz model.Quiz2, quizID uint) (model.QuizResponse, error)
	DeleteQuiz(quizID uint) error
}

type quizUsecase struct {
	qr repository.IQuizRepository
}

func NewQuizUsecase(qr repository.IQuizRepository) IQuizUsecase {
	return &quizUsecase{qr}
}

func (qu *quizUsecase) GetFilteredQuizzes(filters []model.Filter, limit int, random bool) ([]model.QuizResponse, error) {
	quizzes := []model.Quiz2{}
	if err := qu.qr.GetFilteredQuizzes(&quizzes, filters, limit, random); err != nil {
		return nil, err
	}
	resQuizzes := []model.QuizResponse{}
	for _, v := range quizzes {
		q := model.QuizResponse{
			ID:           v.ID,
			Question:     v.Question,
			Code:         v.Code,
			InputSample:  v.InputSample,
			OutputSample: v.OutputSample,
			InputSecret:  v.InputSecret,
			OutputSecret: v.OutputSecret,
		}
		resQuizzes = append(resQuizzes, q)
	}
	return resQuizzes, nil
}

func (qu *quizUsecase) CreateQuiz(quiz model.Quiz2) (model.QuizResponse, error) {
	if err := qu.qr.CreateQuiz(&quiz); err != nil {
		return model.QuizResponse{}, err
	}
	resQuiz := model.QuizResponse{
		ID:            quiz.ID,
		Question:      quiz.Question,
		Code:          quiz.Code,
		InCorrectCode: quiz.InCorrectCode,
		InputSample:   quiz.InputSample,
		OutputSample:  quiz.OutputSample,
		InputSecret:   quiz.InputSecret,
		OutputSecret:  quiz.OutputSecret,
	}
	return resQuiz, nil
}

func (qu *quizUsecase) UpdateQuiz(quiz model.Quiz2, quizID uint) (model.QuizResponse, error) {
	if err := qu.qr.UpdateQuiz(&quiz, quizID); err != nil {
		return model.QuizResponse{}, err
	}

	resQuiz := model.QuizResponse{
		ID:            quiz.ID,
		Question:      quiz.Question,
		InCorrectCode: quiz.InCorrectCode,
		Code:          quiz.Code,
		InputSample:   quiz.InputSample,
		OutputSample:  quiz.OutputSample,
		InputSecret:   quiz.InputSecret,
		OutputSecret:  quiz.OutputSecret,
	}
	return resQuiz, nil
}

func (qu *quizUsecase) DeleteQuiz(quizID uint) error {
	if err := qu.qr.DeleteQuiz(quizID); err != nil {
		return err
	}
	return nil
}
