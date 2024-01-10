package easytime

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		s       string
		wont    time.Time
		wontErr error
	}{
		{s: "2024-01-10T09:49:00.999999+08:00", wont: time.Date(2024, 1, 10, 9, 49, 0, 999999, time.Local)},
		{s: "2024-01-10T09:49:00.999999-08:00", wont: time.Date(2024, 1, 10, 9, 49, 0, 999999, time.Local)},
		{s: "978282000", wont: time.Date(2001, 1, 1, 1, 0, 0, 0, time.Local)},
		{s: "2023-01-09 15:39:00", wont: time.Date(2023, 1, 9, 15, 39, 0, 0, time.Local)},
		{s: "2023-01-09", wont: time.Date(2023, 1, 9, 00, 00, 0, 0, time.Local)},
	}
	for _, tt := range tests {
		v, err := Get(tt.s)
		if err != nil && err != tt.wontErr {
			t.Errorf("Got:%s, Wont:%s", v.Format(time.RFC3339Nano), tt.wont.Format(time.RFC3339Nano))
		}
	}
}
