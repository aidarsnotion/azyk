package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
	"azyk/util"
	"errors"
)

type UserUsecase interface {
	RegisterUser(data models.CreateUserRequest) (*models.User, error)
	Logout(userID int) error
	GetUser(id int) (*models.User, error)
	UpdateUser(id int, data models.UpdateUserRequest) error
	DeleteUser(id int) error
}

type userUsecase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewUserUsecase(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) UserUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

// 🔹 Регистрация пользователя
func (u *userUsecase) RegisterUser(data models.CreateUserRequest) (*models.User, error) {
	if err := util.ValidateStruct(data); err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
		OrgID:    data.OrgID,
	}

	// Проверяем, если это корпоративный пользователь
	if user.Role == models.RoleCorporate && user.OrgID == nil {
		return nil, errors.New("corporate users must belong to an organization")
	}

	// Добавляем пользователя в базу
	err := u.userRepo.CreateUser(user)
	return user, err
}

// 🔹 Выход из системы
func (u *userUsecase) Logout(userID int) error {
	return u.sessionRepo.DeleteSession(userID)
}

// 🔹 Получение пользователя
func (u *userUsecase) GetUser(id int) (*models.User, error) {
	return u.userRepo.GetUserByID(id)
}

// 🔹 Обновление пользователя
func (u *userUsecase) UpdateUser(id int, data models.UpdateUserRequest) error {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	if data.Name != nil {
		user.Name = *data.Name
	}
	if data.Email != nil {
		user.Email = *data.Email
	}

	return u.userRepo.UpdateUser(user)
}

// 🔹 Удаление пользователя
func (u *userUsecase) DeleteUser(id int) error {
	return u.userRepo.DeleteUser(id)
}
