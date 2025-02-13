package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
	"azyk/util"
	"errors"
)

type AuthUseCase interface {
	Login(email, password, deviceId string) (*models.User, error)
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

func (a *authUseCase) Login(email, password, deviceId string) (*models.User, error) {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if !util.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	if user.Role == models.RoleIndividual || user.Role == models.RoleLegalEntity || user.Role == models.RoleScientific {
		existingSession, _ := a.sessionRepo.GetActiveSession(user.ID)
		if existingSession != nil {
			// Kill session
			a.sessionRepo.DeleteSession(user.ID)
		}
	}

	//Creating new session
	_, err = a.sessionRepo.CreateSession(user.ID, deviceId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authUseCase) Logout(userId int) error {
	return a.sessionRepo.DeleteSession(userId)
}
