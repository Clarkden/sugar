package auth

import (
	sugar "sugar/data"
	"time"

	"github.com/google/uuid"
)

func CreateSessionParams(userId int64) sugar.CreateSessionParams {
	sessionId := uuid.New().String()
	sessionCreatedAt := time.Now().Unix()
	sessionExpiresAt := time.Now().Add(24 * 30 * time.Hour).Unix()

	sessionParams := sugar.CreateSessionParams{
		UserID:    &userId,
		SessionID: &sessionId,
		CreatedAt: &sessionCreatedAt,
		ExpiresAt: &sessionExpiresAt,
	}

	return sessionParams
}
