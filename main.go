package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var dist []int64

	entry := os.Args[1]

	src, err := os.Open(entry)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			i, _ := strconv.ParseInt(line, 10, 64)
			dist = append(dist, i)
		}
	}
	defer src.Close()

	fmt.Println(calculateLocation(dist))
}

func calculateLocation(houses []int64) float32 {
	var result float32 = 0

	for _, item := range houses {
		result += float32(item)
	}

	result = result / float32(len(houses))
	return result
}
