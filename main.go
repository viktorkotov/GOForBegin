/*
Вычисление стандартного отклонения:
Вычислим среднее: (1 + 4 + 10) / 3 = 5
Вычисляем квадраты отклонений от их среднего: (5-1)**2=16 (5-4)**2=1 (10-5)**2=25
Вычислим среднее арифметическое (дисперсию) этих значений: (16 + 1 + 25) / 3 = 14
Стандартное отклонение равно квадратному корню дисперсии: квадратный корень из 14 = 3.7 = 4
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var dist []uint64

	entry := os.Args[1]

	src, err := os.Open(entry)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			i, _ := strconv.ParseUint(line, 10, 16)
			dist = append(dist, i)
		}
	}
	defer src.Close()

	fmt.Println(calculateLocation(dist))
}

func calculateLocation(houses []uint64) float64 {
	var sum uint64 = 0
	for i := 0; i < len(houses); i++ {
		sum += houses[i]
	}

	sum /= uint64(len(houses))

	var fsum float64 = 0
	for i := 0; i < len(houses); i++ {
		fsum += math.Pow((float64(sum) - float64(houses[i])), 2)
	}

	fsum /= float64(len(houses))

	return math.Round(math.Sqrt(fsum))
}
