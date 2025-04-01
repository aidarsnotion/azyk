package throttling

import (
	"time"

	"azyk/util/logger"
	"sync"
)

// LoginThrottler — интерфейс для ограничения количества неудачных попыток авторизации.
type LoginThrottler interface {
	IsBlocked(email string) bool
	RecordFailedAttempt(email string)
	ClearAttempts(email string)
}

type loginAttempt struct {
	count       int
	lastAttempt time.Time
}

type InMemoryThrottler struct {
	attempts    map[string]*loginAttempt
	mu          sync.Mutex
	maxAttempts int
	window      time.Duration
}

func NewInMemoryThrottler(maxAttempts int, window time.Duration) LoginThrottler {
	return &InMemoryThrottler{
		attempts:    make(map[string]*loginAttempt),
		maxAttempts: maxAttempts,
		window:      window,
	}
}

func (lt *InMemoryThrottler) IsBlocked(email string) bool {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	attempt, exists := lt.attempts[email]
	if !exists {
		return false
	}
	if time.Since(attempt.lastAttempt) > lt.window {
		delete(lt.attempts, email)
		return false
	}
	return attempt.count >= lt.maxAttempts
}

func (lt *InMemoryThrottler) RecordFailedAttempt(email string) {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	attempt, exists := lt.attempts[email]
	if !exists {
		lt.attempts[email] = &loginAttempt{
			count:       1,
			lastAttempt: time.Now(),
		}
		logger.Log.WithField("email", email).Warn("Первая неудачная попытка входа")
	} else {
		if time.Since(attempt.lastAttempt) > lt.window {
			attempt.count = 1
		} else {
			attempt.count++
		}
		attempt.lastAttempt = time.Now()
		logger.Log.WithField("email", email).Warnf("Неудачная попытка входа #%d", attempt.count)
	}
}

func (lt *InMemoryThrottler) ClearAttempts(email string) {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	delete(lt.attempts, email)
	logger.Log.WithField("email", email).Info("Сброс попыток авторизации")
}
