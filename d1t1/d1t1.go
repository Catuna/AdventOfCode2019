package main

import (
	"common"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calcModuleFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func main() {
	var args = os.Args

	var input, err = common.ReadInputFile(args[1])
	if err != nil {
		log.Fatal("Failed to read task input", err)
	}

	var sum = 0
	for _, moduleMassStr := range input {
		moduleMass, _ := strconv.Atoi(moduleMassStr)
		sum += calcModuleFuel(moduleMass)
	}

	fmt.Println(sum)

}
