package usecase

import (
	"api/model"
	"api/repository"
)

type IUserUsecase interface {
	GetUserByFirebaseUID(firebaseUID string) (model.UserResponse, error)
	CreateUser(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) GetUserByFirebaseUID(firebaseUID string) (model.UserResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserByFirebaseUID(&user, firebaseUID); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return resUser, nil
}

func (uu *userUsecase) CreateUser(user model.User) (model.UserResponse, error) {
	if err := uu.ur.CreateUser(&user); err != nil {
		return model.UserResponse{}, err
	}
	resScore := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	return resScore, nil
}
