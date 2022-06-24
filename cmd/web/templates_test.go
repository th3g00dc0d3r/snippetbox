package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm time.Time
		want string
	} {
		{
			name: "UTC",
			tm: time.Date(2022, 4, 22, 11, 15, 0, 0, time.UTC),
			want: "22 Apr 2022 at 11:15",
		},
		{
			name: "Empty",
			tm: time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm: time.Date(2022, 4, 22, 11, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "22 Apr 2022 at 10:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}