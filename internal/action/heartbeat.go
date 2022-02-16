package action

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/wolftotem4/shaolin-ben-don/internal/client"
)

type HeartbeatAction struct {
	Client client.Client
}

func (action *HeartbeatAction) Call(ctx context.Context, interfaceValue int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	u := fmt.Sprintf("/do/?wicket:interface=:%d:heartbeat:-1:IUnversionedBehaviorListener&wicket:behaviorId=0&wicket:ignoreIfNotActive=true&random=", interfaceValue)

	_, err := action.Client.Call(ctx, u, client.Webpage)
	if err != nil {
		errors.WithStack(err)
	}

	return nil
}
