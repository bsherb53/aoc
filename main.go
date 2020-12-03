package main

import (
	"fmt"
	"time"

	ten "github.com/aoc/days/10_ten"
	eleven "github.com/aoc/days/11_eleven"
	twelve "github.com/aoc/days/12_twelve"
	thirteen "github.com/aoc/days/13_thirteen"
	fourteen "github.com/aoc/days/14_fourteen"
	fifteen "github.com/aoc/days/15_fifteen"
	sixteen "github.com/aoc/days/16_sixteen"
	seventeen "github.com/aoc/days/17_seventeen"
	eighteen "github.com/aoc/days/18_eighteen"
	nineteen "github.com/aoc/days/19_nineteen"
	one "github.com/aoc/days/1_one"
	twenty "github.com/aoc/days/20_twenty"
	twentyone "github.com/aoc/days/21_twentyone"
	twentytwo "github.com/aoc/days/22_twentytwo"
	twentythree "github.com/aoc/days/23_twentythree"
	twentyfour "github.com/aoc/days/24_twentyfour"
	twentyfive "github.com/aoc/days/25_twentyfive"
	twentysix "github.com/aoc/days/26_twentysix"
	twentyseven "github.com/aoc/days/27_twentyseven"
	twentyeight "github.com/aoc/days/28_twentyeight"
	twentynine "github.com/aoc/days/29_twentynine"
	two "github.com/aoc/days/2_two"
	thirty "github.com/aoc/days/30_thirty"
	thirtyone "github.com/aoc/days/31_thirtyone"
	three "github.com/aoc/days/3_three"
	four "github.com/aoc/days/4_four"
	five "github.com/aoc/days/5_five"
	six "github.com/aoc/days/6_six"
	seven "github.com/aoc/days/7_seven"
	eight "github.com/aoc/days/8_eight"
	nine "github.com/aoc/days/9_nine"
)

var allDays = []func(){
	one.Do,
	two.Do,
	three.Do,
	four.Do,
	five.Do,
	six.Do,
	seven.Do,
	eight.Do,
	nine.Do,
	ten.Do,
	eleven.Do,
	twelve.Do,
	thirteen.Do,
	fourteen.Do,
	fifteen.Do,
	sixteen.Do,
	seventeen.Do,
	eighteen.Do,
	nineteen.Do,
	twenty.Do,
	twentyone.Do,
	twentytwo.Do,
	twentythree.Do,
	twentyfour.Do,
	twentyfive.Do,
	twentysix.Do,
	twentyseven.Do,
	twentyeight.Do,
	twentynine.Do,
	thirty.Do,
	thirtyone.Do,
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
