package one

import "fmt"
import "log"

import "github.com/cpjk/advent_of_code_2019/util"

// https://adventofcode.com/2019/day/1
// https://adventofcode.com/2019/day/1#part2

func One() {
	moduleMassesInt, err := util.InjestIntegerList("input.txt")
	if err != nil {
		log.Panic(err)
	}

	totalFuel := totalFuelForModules(moduleMassesInt)
	fmt.Println("Day 1 answer: ", totalFuel)
}

func totalFuelForModules(masses []int) int {
	totalFuel := 0
	for _, mass := range masses {
		totalFuel += fuelForModule(mass)
	}

	return totalFuel
}

func fuelForModule(moduleMass int) int {
	return fuelForMass(moduleMass)
}

func fuelForMass(mass int) int {
	return doFuelForMass(mass, 0)
}

// Recursively build up the fuel required for the input mass,
// memoizing the current sum in currentFuel
func doFuelForMass(mass int, currentFuel int) int {
	if mass <= 0 {
		return currentFuel
	}

	newFuelMass := (mass / 3) - 2
	if newFuelMass <=0 {
		return doFuelForMass(newFuelMass, currentFuel)
	}

	newFuelTotal := newFuelMass + currentFuel
	return doFuelForMass(newFuelMass, newFuelTotal)
}


