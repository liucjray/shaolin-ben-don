package typesjson

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"
)

func TestUnmarshalProgress(t *testing.T) {
	var data Progress

	j, err := embedFS.ReadFile("testdata/progress.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(j, &data); err != nil {
		t.Fatal(err)
	}

	if len(data.Data) == 0 {
		t.Fatal("failed to unmarshal JSON")
	}

	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		t.Fatal(err)
	}

	if !time.Date(2022, 2, 8, 10, 0, 0, 0, loc).Equal(data.Data[0].ExpireDate.Time) {
		t.Errorf("unexpected date: %s", data.Data[0].ExpireDate.Time.String())
	}

	if data.Data[0].RemainSecondBeforeExpire != 2983*time.Second {
		t.Errorf("unexpected value: %d", data.Data[0].RemainSecondBeforeExpire)
	}
}

func TestIsExpiring(t *testing.T) {
	var item = &ProgressItem{RemainSecondBeforeExpire: 5 * 60 * time.Second}

	if !item.IsExpiring(6 * time.Minute) {
		t.Error("expect positive value")
	}
	if !item.IsExpiring(5 * time.Minute) {
		t.Error("expect positive value")
	}
	if item.IsExpiring(4 * time.Minute) {
		t.Error("expect negative value")
	}
	if item.IsExpiring(0) {
		t.Error("expect negative value")
	}
}

func TestUpdateRemainSecondBeforeExpire(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2022-02-18T12:00:00+08:00")

	item := &ProgressItem{ExpireDate: sql.NullTime{Time: now.Add(3 * time.Minute), Valid: true}}
	item.UpdateRemainSecondBeforeExpire(now)

	if item.RemainSecondBeforeExpire != 3*time.Minute {
		t.Errorf("unexpected value (expected %s, actual %s)", 3*time.Minute, item.RemainSecondBeforeExpire)
	}

	item.UpdateRemainSecondBeforeExpire(now.Add(4 * time.Minute))
	if item.RemainSecondBeforeExpire != 0 {
		t.Errorf("unexpected value (expected 0, actual %s)", item.RemainSecondBeforeExpire)
	}
}
