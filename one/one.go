package one

import "fmt"
import "log"

import "github.com/cpjk/advent_of_code_2019/util"

/*
Fuel required to launch a given module is based on its mass.
Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.
The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.

What is the sum of the fuel requirements for all of the modules on your spacecraft?
*/

func One() {
	moduleMassesInt, err := util.InjestIntegerList("input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(moduleMassesInt)

	totalFuel := totalFuelForModules(moduleMassesInt)
	fmt.Println("total fuel: ", totalFuel)
}

func fuelForModule(moduleMass int) int {
	fuel := (moduleMass / 3) - 2
	return fuel
}

func totalFuelForModules(masses []int) int {
	totalFuel := 0
	for _, mass := range masses {
		totalFuel += fuelForModule(mass)
	}

	return totalFuel
}

