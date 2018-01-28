package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func printHelp() {
	fmt.Println("Usage: rand [min] max")
}

func toInt(value string) int {
	maxNum, err := strconv.Atoi(value)
	if err == nil {
		return maxNum
	}

	fmt.Println(err)
	os.Exit(1)
	return -1
}

func main() {
	args := os.Args
	size := len(os.Args)

	var min, max int

	switch size {
	case 2:
		min = 0
		max = toInt(args[1])
	case 3:
		min = toInt(args[1])
		max = toInt(args[2])
	default:
		printHelp()
		os.Exit(0)
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomValue := random.Intn(max - min)
	randomRealValue := randomValue + min

	fmt.Println(randomRealValue)
	os.Exit(0)
}
