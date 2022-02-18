package ctrl

import (
	"time"

	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

type PendingItems struct {
	items     []*typesjson.ProgressItem
	priorTime time.Duration
	timer     *time.Timer
}

func NewPendingItems(priorTime time.Duration) *PendingItems {
	var data = PendingItems{
		priorTime: priorTime,
	}
	return &data
}

func (info *PendingItems) Update(items []*typesjson.ProgressItem) {
	info.items = items
	info.UpdateTimer()
}

func (info *PendingItems) UpdateTimer() {
	info.updateTimer(info.NextTime())
}

func (info *PendingItems) UpdateRemainSecondBeforeExpireValues() {
	for _, item := range info.items {
		item.UpdateRemainSecondBeforeExpire(time.Now())
	}
}

func (info *PendingItems) NextTime() time.Duration {
	var next time.Duration
	for _, item := range info.items {
		if item.RemainSecondBeforeExpire > 0 && (next == 0 || item.RemainSecondBeforeExpire < next) && !item.IsExpiring(info.priorTime) {
			next = item.RemainSecondBeforeExpire
		}
	}
	if next == 0 {
		return 0
	}

	return next - info.priorTime
}

func (info *PendingItems) Size() int {
	return len(info.items)
}

func (info *PendingItems) Chan() <-chan time.Time {
	if info.timer != nil {
		return info.timer.C
	}
	return nil
}

func (info *PendingItems) ExtractExpiringItems() []*typesjson.ProgressItem {
	var (
		reports  []*typesjson.ProgressItem
		pendings []*typesjson.ProgressItem
	)
	for _, item := range info.items {
		if item.IsExpiring(info.priorTime) {
			reports = append(reports, item)
		} else {
			pendings = append(pendings, item)
		}
	}
	info.items = pendings
	return reports
}

func (info *PendingItems) updateTimer(next time.Duration) {
	if info.timer != nil {
		if !info.timer.Stop() {
			<-info.timer.C
		}
	}

	if next > 0 {
		if info.timer == nil {
			info.timer = time.NewTimer(next)
		} else {
			info.timer.Reset(next)
		}
	} else {
		info.timer = nil
	}
}
