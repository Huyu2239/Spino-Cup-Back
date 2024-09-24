package usecase

import (
	"api/model"
	"api/repository"
)

type IQuizUsecase interface {
	GetFilteredQuizzes(filters []model.Filter) ([]model.QuizResponse, error)
	CreateQuiz(quiz model.Quiz) (model.QuizResponse, error)
	UpdateQuiz(quiz model.Quiz, quizID uint) (model.QuizResponse, error)
	DeleteQuiz(quizID uint) error
	GetQuizAnswer(quizID uint) (model.AnswerResponse, error)
	CheckQuiz(quizID uint, ansX uint, ansY uint) (model.CheckResponse, error)
}

type quizUsecase struct {
	qr repository.IQuizRepository
}

func NewQuizUsecase(qr repository.IQuizRepository) IQuizUsecase {
	return &quizUsecase{qr}
}

func (qu *quizUsecase) GetFilteredQuizzes(filters []model.Filter) ([]model.QuizResponse, error) {
	quizzes := []model.Quiz{}
	if err := qu.qr.GetFilteredQuizzes(&quizzes, filters); err != nil {
		return nil, err
	}
	resQuizzes := []model.QuizResponse{}
	for _, v := range quizzes {
		q := model.QuizResponse{
			ID:         v.ID,
			Question:   v.Question,
			Difficulty: v.Difficulty,
			Language:   v.Language,
		}
		resQuizzes = append(resQuizzes, q)
	}
	return resQuizzes, nil
}

func (qu *quizUsecase) CreateQuiz(quiz model.Quiz) (model.QuizResponse, error) {
	if err := qu.qr.CreateQuiz(&quiz); err != nil {
		return model.QuizResponse{}, err
	}
	resQuiz := model.QuizResponse{
		ID:         quiz.ID,
		Question:   quiz.Question,
		Difficulty: quiz.Difficulty,
		Language:   quiz.Language,
	}
	return resQuiz, nil
}

func (qu *quizUsecase) UpdateQuiz(quiz model.Quiz, quizID uint) (model.QuizResponse, error) {
	if err := qu.qr.UpdateQuiz(&quiz, quizID); err != nil {
		return model.QuizResponse{}, err
	}

	resQuiz := model.QuizResponse{
		ID:         quiz.ID,
		Question:   quiz.Question,
		Difficulty: quiz.Difficulty,
		Language:   quiz.Language,
	}
	return resQuiz, nil
}

func (qu *quizUsecase) DeleteQuiz(quizID uint) error {
	if err := qu.qr.DeleteQuiz(quizID); err != nil {
		return err
	}
	return nil
}

func (qu *quizUsecase) GetQuizAnswer(quizID uint) (model.AnswerResponse, error) {
	quiz := model.Quiz{}
	if err := qu.qr.GetQuizByID(&quiz, quizID); err != nil {
		return model.AnswerResponse{}, err
	}
	resAns := model.AnswerResponse{
		ID:          quiz.ID,
		AnswerX:     quiz.AnswerX,
		AnswerY:     quiz.AnswerY,
		Explanation: quiz.Explanation,
	}

	return resAns, nil
}

func (qu *quizUsecase) CheckQuiz(quizID uint, ansX uint, ansY uint) (model.CheckResponse, error) {
	quiz := model.Quiz{}
	if err := qu.qr.GetQuizByID(&quiz, quizID); err != nil {
		return model.CheckResponse{}, err
	}

	if quiz.AnswerX == ansX && quiz.AnswerY == ansY {
		return model.CheckResponse{ID: quizID, IsCorrect: true}, nil
	}

	return model.CheckResponse{ID: quizID, IsCorrect: false}, nil
}
