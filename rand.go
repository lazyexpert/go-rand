package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	max := args[1]

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxNum, err := strconv.Atoi(max)
	if err == nil {
		fmt.Println(random.Intn(maxNum))

		os.Exit(0)
	}

	fmt.Println(err)
	os.Exit(1)
}
