package ctrl

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/wolftotem4/shaolin-ben-don/internal/faker"
	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

func TestPendingItems(t *testing.T) {
	priorTime := 10 * time.Minute
	data := NewPendingItems(priorTime)

	// expect 3 expiring items and 10 unexpiring items
	{
		for i := 0; i < 100; i++ {
			data.Update(GenerateExpiringItemsForTesting(3, 10, priorTime))
			if size := len(data.ExtractExpiringItems()); size != 3 {
				t.Fatalf("size of expiring items is unmatched (size: %d)", size)
			}
			if nextTime := data.NextTime(); nextTime <= 0 {
				t.Fatalf("unexpected value (next time: %s)", nextTime)
			}
			if data.Chan() == nil {
				t.Fatal("unexpected nil timer")
			}
		}
	}

	// expect no timer with no expiring items
	{
		data.Update(GenerateExpiringItemsForTesting(15, 0, priorTime))
		if nextTime := data.NextTime(); nextTime != 0 {
			t.Fatalf("expected zero value (next time: %s)", data.NextTime())
		}
	}

	// test timer
	{
		item := faker.MakeProgressItem()
		item.RemainSecondBeforeExpire = priorTime + 2*time.Minute
		item.ExpireDate = sql.NullTime{Time: time.Now().Add(item.RemainSecondBeforeExpire), Valid: true}

		data.Update([]*typesjson.ProgressItem{item})

		if nextTime := data.NextTime(); nextTime != 2*time.Minute {
			t.Fatalf("unexpected value (next time: %s", data.NextTime())
		}
	}

	// ignore zero item.RemainSecondBeforeExpire
	{
		items := GenerateExpiringItemsForTesting(0, 2, priorTime)
		items[1].ExpireDate = sql.NullTime{Valid: false}
		items[1].RemainSecondBeforeExpire = 0

		data.Update(items)

		if data.NextTime() == 0 {
			t.Fatal("expected non-zero value")
		}
	}
}

func GenerateExpiringItemsForTesting(expiring int, remote int, priorTime time.Duration) []*typesjson.ProgressItem {
	var items []*typesjson.ProgressItem

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	{
		prepares := faker.MakeProgressItems(expiring)
		for i := range prepares {
			expireMinute := time.Duration(r.Intn(int(priorTime/time.Minute))+1) * time.Minute
			prepares[i].ExpireDate = sql.NullTime{Time: time.Now().Add(expireMinute), Valid: true}
			prepares[i].RemainSecondBeforeExpire = expireMinute
		}

		items = append(items, prepares...)
	}

	{
		prepares := faker.MakeProgressItems(remote)
		for i := range prepares {
			expireMinute := priorTime + time.Duration(r.Intn(20)+1)*time.Minute
			prepares[i].ExpireDate = sql.NullTime{Time: time.Now().Add(expireMinute), Valid: true}
			prepares[i].RemainSecondBeforeExpire = expireMinute
		}

		items = append(items, prepares...)
	}

	return items
}
