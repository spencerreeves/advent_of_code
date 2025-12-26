package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMirrorRepeats(t *testing.T) {
	type test struct {
		Name    string
		Lower   int
		Upper   int
		Repeats []int
	}

	tests := []test{
		{Name: "11-22", Lower: 11, Upper: 22, Repeats: []int{11, 22}},
		{Name: "95-115", Lower: 95, Upper: 115, Repeats: []int{99}},
		{Name: "998-1012", Lower: 998, Upper: 1012, Repeats: []int{1010}},
		{Name: "1188511880-1188511890", Lower: 1188511880, Upper: 1188511890, Repeats: []int{1188511885}},
		{Name: "1698522-1698528", Lower: 1698522, Upper: 1698528, Repeats: nil},
		{Name: "446443-446449", Lower: 446443, Upper: 446449, Repeats: []int{446446}},
		{Name: "38593856-38593862", Lower: 38593856, Upper: 38593862, Repeats: []int{38593859}},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			repeats := GetMirrorRepeats(tc.Lower, tc.Upper)
			assert.Equal(t, tc.Repeats, repeats)
		})
	}
}

func TestIsValidPattern(t *testing.T) {
	type test struct {
		Name    string
		Pattern string
		ID      string
		Valid   bool
	}

	tests := []test{
		{Name: "repeats one", Pattern: "11", ID: "11", Valid: true},
		{Name: "repeats twice", Pattern: "11885", ID: "1188511885", Valid: true},
		{Name: "repeats many", Pattern: "12", ID: "121212121212", Valid: true},
		{Name: "does not repeat", Pattern: "98", ID: "99", Valid: false},
		{Name: "longer ID", Pattern: "11", ID: "1", Valid: false},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Valid, IsValidPattern(tc.ID, tc.Pattern))
		})
	}
}

func TestIsInvalidID(t *testing.T) {
	type test struct {
		ID      string
		Valid   bool
		Pattern string
	}

	tests := []test{
		{ID: "11", Valid: true, Pattern: "1"},
		{ID: "2121212121", Valid: true, Pattern: "21"},
		{ID: "824824824", Valid: true, Pattern: "824"},
		{ID: "123", Valid: false, Pattern: ""},
	}

	for _, tc := range tests {
		t.Run(tc.ID, func(t *testing.T) {
			valid, pattern := IsInvalidID(tc.ID)

			assert.Equal(t, tc.Valid, valid)
			if tc.Valid {
				assert.Equal(t, tc.Pattern, pattern)
			}
		})
	}
}
