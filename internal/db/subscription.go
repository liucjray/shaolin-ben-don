package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Handler struct {
	DB *sqlx.DB
}

func (handler *Handler) AddSubscription(ctx context.Context, chatId int64) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := handler.DB.ExecContext(ctx, "INSERT OR IGNORE INTO subscriptions (chat_id) VALUES(?)", chatId)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (handler *Handler) DeleteSubscription(ctx context.Context, chatId int64) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := handler.DB.ExecContext(ctx, "DELETE FROM subscriptions WHERE chat_id = ?", chatId)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (handler *Handler) GetSubscriptions(ctx context.Context) ([]int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := handler.DB.QueryContext(ctx, "SELECT chat_id FROM subscriptions")
	if err != nil {
		return []int64{}, nil
	}

	var result []int64
	for rows.Next() {
		var chatId int64
		if err := rows.Scan(&chatId); err != nil {
			return []int64{}, nil
		}
		result = append(result, chatId)
	}

	return result, nil
}
