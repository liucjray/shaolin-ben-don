package faker

import (
	"database/sql"
	"encoding/base64"
	"math/rand"
	"time"

	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

func MakeProgressItems(n int) []*typesjson.ProgressItem {
	var items = make([]*typesjson.ProgressItem, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range items {
		var (
			expireMinute time.Duration = time.Duration(r.Intn(20)+1) * time.Minute
		)

		items[i] = &typesjson.ProgressItem{
			OrderHashId: func() string {
				var orderHashBytes = make([]byte, 10)
				r.Read(orderHashBytes)
				return base64.StdEncoding.EncodeToString(orderHashBytes)
			}(),
			ExpireDate: func() sql.NullTime {
				return sql.NullTime{
					Time:  time.Now().Add(expireMinute),
					Valid: true,
				}
			}(),
			MaxQty:       0,
			MaxTotalCost: 0,
			InProgress: func() bool {
				return r.Intn(2) == 1
			}(),
			Size: func() int {
				return r.Intn(20) + 1
			}(),
			Total: func() int {
				return r.Intn(5000)
			}(),
			Announcement: "",
			ShopName: func() string {
				var b = make([]byte, 10)
				r.Read(b)
				return base64.StdEncoding.EncodeToString(b)
			}(),
			Originator:               "1",
			GroupId:                  1,
			Unlockable:               true,
			PasswordLocked:           false,
			RemainSecondBeforeExpire: expireMinute,
		}
	}
	return items
}

func MakeProgressItem() *typesjson.ProgressItem {
	return MakeProgressItems(1)[0]
}
