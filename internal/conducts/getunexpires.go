package conducts

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wolftotem4/shaolin-ben-don/internal/app"
	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

func GetUnexpiredItems(ctx context.Context, app *app.App, interfaceValue int) (*FetchItemInfo, error) {
	info, err := FetchItems(ctx, app, interfaceValue)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var items []*typesjson.ProgressItem
	for _, item := range info.Items {
		if !item.IsExpired() {
			items = append(items, item)
		}
	}

	return &FetchItemInfo{
		Items:     items,
		Interface: info.Interface,
	}, nil
}
