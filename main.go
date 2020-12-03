package main

import (
	"fmt"
	"time"

	__one "github.com/aoc/days/1_one"
	__two "github.com/aoc/days/2_two"
	__three "github.com/aoc/days/3_three"
	__four "github.com/aoc/days/4_four"
)

var allDays = []func(){
	__one.Do,
	__two.Do,
	__three.Do,
	__four.Do,
}

func main() {
	fmt.Println("Advent Of Code 2020")
	// enter the day, 0 runs all days

	day := time.Now().Day()

	for i := range allDays {
		if i == day-1 || day == 0 {
			start := time.Now()
			allDays[i]()
			t := time.Since(start)
			fmt.Println(fmt.Sprintf("Day %v ran in %s", i+1, t))
		}
	}

}
