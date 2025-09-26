package controller

import (
	"fmt"
	"log"
	"time"
)

// 最大リトライ回数
const maxRetry = 3
const retryDelay = 30 * time.Second

func RegisterUserWithRetry() error {
	var lastErr error
	var updateErr error

	for i := 0; i < maxRetry; i++ {
		lastErr = RegisterUserOnce()
		if lastErr != nil {
			log.Printf("初回ユーザ登録失敗: %v リトライ %d/%d", lastErr, i+1, maxRetry)
			time.Sleep(retryDelay)
			continue
		}
		updateErr = ConnectUserData()
		if updateErr != nil {
			log.Printf("初回ユーザ接続失敗: %v リトライ %d/%d", lastErr, i+1, maxRetry)
			time.Sleep(retryDelay)
			continue
		}
		return nil
	}

	return fmt.Errorf("初回ユーザ登録に失敗しました: %w", lastErr)
}
