package ctrl

import typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"

type ReportedItems struct {
	Items map[string]bool
}

func (reported *ReportedItems) ExtractUnreported(items []*typesjson.ProgressItem) []*typesjson.ProgressItem {
	var reports []*typesjson.ProgressItem
	for _, item := range items {
		if !reported.Items[item.OrderHashId] {
			reports = append(reports, item)
		}
	}
	return reports
}

func (reported *ReportedItems) MarkReported(items []*typesjson.ProgressItem) {
	var result = make(map[string]bool)
	for _, item := range items {
		result[item.OrderHashId] = true
	}
	reported.Items = result
}
