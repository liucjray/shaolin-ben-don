package main

import (
	"database/sql"
	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"
	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

func readMockFile() ([]*typesjson.ProgressItem, error) {
	var (
		items []*typesjson.ProgressItem
	)

	j, err := os.ReadFile(*flagMock)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var progress typesjson.Progress
	if err := json.Unmarshal(j, &progress); err != nil {
		return nil, errors.WithStack(err)
	}

	// convert to slice of pointers
	items = make([]*typesjson.ProgressItem, len(progress.Data))
	for i := range progress.Data {
		items[i] = &progress.Data[i]
	}

	calcRemainingTimeForMocks(items)

	return items, nil
}

func calcRemainingTimeForMocks(items []*typesjson.ProgressItem) {
	now := time.Now()
	for _, item := range items {
		if !item.ExpireDate.Valid || !item.ExpireDate.Time.After(now) {
			item.ExpireDate = sql.NullTime{Valid: false}
			item.RemainSecondBeforeExpire = 0
			continue
		}

		item.RemainSecondBeforeExpire = item.ExpireDate.Time.Sub(now)
	}
}
