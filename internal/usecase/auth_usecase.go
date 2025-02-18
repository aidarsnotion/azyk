package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
	"azyk/util"
	"errors"
)

type AuthUseCase interface {
	Login(email, password, deviceId string) (*models.User, string, error)
	Logout(userId int) error
}

type authUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewAuthUseCase(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) AuthUseCase {
	return &authUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (a *authUseCase) Login(email, password, deviceId string) (*models.User, string, error) {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Проверяем пароль
	if !util.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid email or password")
	}

	// Проверяем активные сессии (если роль требует 1 устройство)
	if user.Role == models.RoleIndividual || user.Role == models.RoleLegalEntity || user.Role == models.RoleScientific {
		existingSession, _ := a.sessionRepo.GetActiveSession(user.ID)
		if existingSession != nil {
			a.sessionRepo.DeleteSession(user.ID) // Удаляем старую сессию
		}
	}

	// Создаем новую сессию
	err = a.sessionRepo.CreateSession(user.ID, deviceId)
	if err != nil {
		return nil, "", err
	}

	// Генерируем JWT-токен
	token, err := util.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (a *authUseCase) Logout(userId int) error {
	return a.sessionRepo.DeleteSession(userId)
}
