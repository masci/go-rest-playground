package storage

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMakeID(t *testing.T) {
	rand.Seed(42) // make tests reproducible by having the same rand sequence every time

	var tests = []struct {
		input string
		want  string
	}{
		{"foo", "FO2305"},
		{"bar", "BA4987"},
		{"BAZ", "BA1668"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			id := makeID(tt.input)
			if id != tt.want {
				t.Errorf("got %s, want %s", id, tt.want)
			}
		})
	}
}

func TestCreateTime(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"2022-01-21", "2022-01-21 00:00:00 +0000 UTC"},
		{"1999-01-01", "1999-01-01 00:00:00 +0000 UTC"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			date := createTime(tt.input)
			if date.String() != tt.want {
				t.Errorf("got %s, want %s", date, tt.want)
			}
		})
	}
}
