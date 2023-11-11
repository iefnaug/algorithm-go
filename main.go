package main

import (
	"algorithm/math"
	"fmt"
)

func A0013() {
	sum := math.RomanToInt("MCMXCIV")
	fmt.Println(sum)
}

func A0191() {
	weight := math.HammingWeight2(10)
	fmt.Println(weight)
}

func main() {
	A0191()
}
