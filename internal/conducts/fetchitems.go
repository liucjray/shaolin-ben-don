package conducts

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/wolftotem4/shaolin-ben-don/internal/action"
	"github.com/wolftotem4/shaolin-ben-don/internal/api"
	"github.com/wolftotem4/shaolin-ben-don/internal/app"
	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

// list and recommanded update time
type FetchItemInfo struct {
	// prompting items
	Items []*typesjson.ProgressItem

	Interface int
}

func FetchItems(ctx context.Context, app *app.App, interfaceValue int) (*FetchItemInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	var (
		pendingItems []*typesjson.ProgressItem
	)

	if interfaceValue == 0 {
		// Refresh session status by loading the main page
		act := &action.DashboardAction{Client: app.Client}
		data, err := act.Update(ctx)
		// We don't need to handle the error here

		// Re-login if needed
		logged, err := action.PerformLoginIfRequired(ctx, app.Client, app.Config, err)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if logged != nil {
			interfaceValue = logged.Interface
		} else {
			interfaceValue = data.Interface
		}
	}

	// Load all pending items
	{
		data, err := api.CallProgress(ctx, app.Client)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// convert to slice of pointers
		for i := range data.Data {
			pendingItems = append(pendingItems, &data.Data[i])
		}
	}

	return &FetchItemInfo{
		Items:     pendingItems,
		Interface: interfaceValue,
	}, nil
}
