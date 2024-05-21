package main

import (
	"fmt"
	"sort"
)

type Solution struct {
	count      int
	frequency  map[int]int
	percentage map[int]float64
}

func (s *Solution) attack(health, level int) {
	if health <= 0 {
		s.frequency[level]++
		s.count++
		return
	}

	health += 2

	for i := 20; i <= 35; i++ {
		s.attack(health-i, level+1)
	}
}

func (s *Solution) calculatePercentage() {
	for level, freq := range s.frequency {
		s.percentage[level] = float64(freq) / float64(s.count) * 100
	}
}

func (s *Solution) printResult() {
	fmt.Println("Total Number: ", s.count)

	fmt.Println()

	var levels []int
	for level := range s.frequency {
		levels = append(levels, level)
	}
	sort.Ints(levels)

	for _, level := range levels {
		fmt.Printf("Level %d: %d\n", level, s.frequency[level])
	}

	fmt.Println()

	s.calculatePercentage()
	for _, level := range levels {
		fmt.Printf("Level %d: %.5f%%\n", level, s.percentage[level])
	}
}

func main() {
	misyoner := &Solution{
		count:      0,
		frequency:  make(map[int]int),
		percentage: make(map[int]float64),
	}

	misyoner.attack(98, 0)
	misyoner.printResult()
}
