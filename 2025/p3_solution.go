package main

import (
	"fmt"
	"math"
	"strconv"
)

func P3GetFirstStar(batteries []string) int {
	maxJoltage := 0
	for _, battery := range batteries {
		joltage := 0
		for d1Idx, d1Str := range battery {
			for d2Idx := d1Idx + 1; d2Idx < len(battery); d2Idx++ {
				d1, _ := strconv.Atoi(string(d1Str))
				d2, _ := strconv.Atoi(string(battery[d2Idx]))
				if d1*10+d2 > joltage {
					joltage = d1*10 + d2
				}
			}
		}
		maxJoltage += joltage
	}

	return maxJoltage
}

func P3GetSecondStar(in []string) int {
	const BatteryBankSize = 12
	maxJoltage := 0

	for _, rawBatteryBank := range in {
		if len(rawBatteryBank) < BatteryBankSize {
			fmt.Printf("Invalid battery bank: %v\n", rawBatteryBank)
			return 0
		}

		batteryBank := NewBatteryBank(rawBatteryBank)

		// Start with an initial battery using the right most digits
		optimalBatteryConfig := batteryBank[len(batteryBank)-BatteryBankSize:]

		for idx := len(batteryBank) - BatteryBankSize - 1; idx >= 0; idx-- {
			optimalBatteryConfig = AddBattery(optimalBatteryConfig, batteryBank[idx])
		}

		joltage := 0
		for i := 0; i < len(optimalBatteryConfig); i++ {
			joltage += int(math.Pow10(len(optimalBatteryConfig)-i-1)) * optimalBatteryConfig[i]
		}

		maxJoltage += joltage
	}

	return maxJoltage
}

func NewBatteryBank(in string) []int {
	out := make([]int, len(in))
	for i := range out {
		c, err := strconv.Atoi(string(in[i]))
		if err != nil {
			fmt.Printf("Invalid input: %v | %v\n", in, err)
			return nil
		}

		out[i] = c
	}

	return out
}

// AddBattery takes a battery and a cell to add.
// It will try to add the cell to the battery from left to right.
// If it finds a cell to replace, it will then try to move the current cell downward
func AddBattery(battery []int, cell int) []int {
	if len(battery) == 0 {
		return []int{}
	}

	if battery[0] <= cell {
		return append([]int{cell}, AddBattery(battery[1:], battery[0])...)
	}

	return battery
}
