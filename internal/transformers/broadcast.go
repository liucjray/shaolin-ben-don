package transformers

import (
	"fmt"
	"strings"

	"github.com/wolftotem4/shaolin-ben-don/internal/client"
	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

type LinkItems struct {
	Items  []*typesjson.ProgressItem
	Client client.Client
}

func (link *LinkItems) String() string {
	var buffer []string
	for _, item := range link.Items {
		u, _ := link.Client.Endpoint(item.GetPath())

		buffer = append(buffer, fmt.Sprintf("%s\n%s", item.ShopName, u.String()))
	}
	return strings.Join(buffer, "\n\n")
}
