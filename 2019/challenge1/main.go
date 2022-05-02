package challenge1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const input  = "challenge1/input1.txt"

func checkError(err error) {
	if err != nil {
		log.Fatal("[Error]: ", err)
		panic(err)
	}
}

func fuelReqForMass(mass int) int {
	return (mass / 3) - 2
}

/// Calculate the fuel requirements based on the masses in the input file.
///
/// Ignore additional mass added by fuel
func part1(scanner *bufio.Scanner) int {
	fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		checkError(err)

		fuel += fuelReqForMass(mass)
	}

	return fuel
}

/// Calculate the fuel requirements based on the masses in the input file including additional mass added by fuel
///
/// We do this by calculating the fuel required for the mass, and then recursively calculating the fuel required for
/// the additional fuel we added.
///
/// Note: The instructions specify that this is *PER MODULE* not *TOTAL*. That means there will be cumulation of
/// unaccounted mass from the leftovers of each module, which will NOT be included in the total
func part2(scanner *bufio.Scanner) int {
	fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		checkError(err)

		for fuelReqForMass(mass) > 0 {
			fuel += fuelReqForMass(mass)
			mass = fuelReqForMass(mass)
		}
	}

	return fuel
}

func Run() {
	fmt.Println("Running Challenge 1")

	file, err := os.Open(input)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Uncomment one to get the fuel calculations
	// fuel := part1(scanner) 							// Value: 3511949
	fuel := part2(scanner)                           // Value: 5265045

	fmt.Println("Total fuel required: ", fuel)
}

