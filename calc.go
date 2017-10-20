package calc

import "time"

func Sum(x int, y int) int {
	return x + y
}

func SumSlow(x int, y int) int {
	time.Sleep(2 * time.Second)
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}
