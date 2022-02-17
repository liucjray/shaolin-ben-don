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
	info.updateTimer(info.NextTime())
}

func (info *PendingItems) NextTime() time.Duration {
	var refTime time.Duration
	for _, item := range info.items {
		if (refTime == 0 || item.RemainSecondBeforeExpire < refTime) && !item.IsExpiring(info.priorTime) {
			refTime = item.RemainSecondBeforeExpire
		}
	}
	if refTime == 0 {
		return 0
	}
	return refTime - info.priorTime
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
	info.updateTimer(info.NextTime())
	return reports
}

func (info *PendingItems) updateTimer(refTime time.Duration) {
	if info.timer != nil {
		if !info.timer.Stop() {
			<-info.timer.C
		}
	}

	if refTime > 0 && len(info.items) > 0 {
		if info.timer == nil {
			info.timer = time.NewTimer(refTime)
		} else {
			info.timer.Reset(refTime)
		}
	} else {
		info.timer = nil
	}
}
