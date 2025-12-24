package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDial(t *testing.T) {
	type test struct {
		Name          string
		Size          int
		StartAt       int
		Invalid       bool
		ExpectedValue int
	}

	tests := []test{
		{Name: "Invalid Dial", Size: 0, StartAt: 0, Invalid: true},
		{Name: "Minimal size", Size: 1, StartAt: 0, ExpectedValue: 0},
		{Name: "Expect size", Size: 99, StartAt: 50, ExpectedValue: 50},
		{Name: "Negative position", Size: 99, StartAt: -100, ExpectedValue: 0},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			d := NewDial(tc.Size, tc.StartAt)
			assert.True(t, (tc.Invalid && d == nil) || (!tc.Invalid && d != nil && d.CurrNode.Value == tc.ExpectedValue), "Expected valid dial")
		})
	}
}

func TestDial(t *testing.T) {
	type test struct {
		Name           string
		Dial           *Dial
		Actions        []int
		ExpectedValues []int
	}

	tests := []test{
		{Name: "Example input", Dial: NewDial(99, 50), Actions: []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}, ExpectedValues: []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			for i, action := range tc.Actions {
				tc.Dial.Turn(action)
				assert.Equal(t, tc.ExpectedValues[i], tc.Dial.CurrNode.Value)
			}
		})
	}
}
