package auth

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
	"azyk/internal/usecase/throttling"
	"azyk/util"
	"azyk/util/logger"
	"errors"
)

// AuthUseCase — интерфейс для работы с авторизацией.
type AuthUseCase interface {
	Login(email, password, deviceId string) (*models.User, string, error)
	Logout(userId int) error
}

type authUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	throttler   throttling.LoginThrottler
}

// NewAuthUseCase создаёт новый экземпляр AuthUseCase, инжектируя зависимости.
func NewAuthUseCase(userRepo repository.UserRepository, sessionRepo repository.SessionRepository, throttler throttling.LoginThrottler) AuthUseCase {
	return &authUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		throttler:   throttler,
	}
}

func (a *authUseCase) Login(email, password, deviceId string) (*models.User, string, error) {
	// Проверяем, не заблокирован ли пользователь из-за большого количества неудачных попыток.
	if a.throttler.IsBlocked(email) {
		logger.Log.WithField("email", email).Warn("Пользователь заблокирован из-за слишком большого числа неудачных попыток")
		return nil, "", errors.New("слишком много попыток авторизации, попробуйте позже")
	}

	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		a.throttler.RecordFailedAttempt(email)
		logger.Log.WithField("email", email).Warn("Не удалось получить пользователя по email")
		return nil, "", errors.New("invalid email or password")
	}

	// Проверяем пароль.
	if !util.CheckPasswordHash(password, user.Password) {
		a.throttler.RecordFailedAttempt(email)
		logger.Log.WithField("email", email).Warn("Неверный пароль")
		return nil, "", errors.New("invalid email or password")
	}

	// При успешном логине сбрасываем счётчик неудачных попыток.
	a.throttler.ClearAttempts(email)
	logger.Log.WithField("email", email).Info("Пользователь успешно авторизовался")

	// Если для данного типа пользователя разрешено только одно устройство, удаляем предыдущую сессию.
	if user.Role == models.RoleIndividual || user.Role == models.RoleLegalEntity || user.Role == models.RoleScientific {
		existingSession, _ := a.sessionRepo.GetActiveSession(user.ID)
		if existingSession != nil {
			logger.Log.WithField("userID", user.ID).Info("Удаляем предыдущую активную сессию")
			a.sessionRepo.DeleteSession(user.ID)
		}
	}

	// Создаём новую сессию.
	if err := a.sessionRepo.CreateSession(user.ID, deviceId); err != nil {
		logger.Log.WithError(err).Error("Ошибка при создании новой сессии")
		return nil, "", err
	}

	// Генерируем JWT-токен.
	token, err := util.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка генерации JWT-токена")
		return nil, "", err
	}

	return user, token, nil
}

func (a *authUseCase) Logout(userId int) error {
	err := a.sessionRepo.DeleteSession(userId)
	if err != nil {
		logger.Log.WithField("userID", userId).Error("Ошибка выхода пользователя")
	} else {
		logger.Log.WithField("userID", userId).Info("Пользователь успешно вышел из системы")
	}
	return err
}
