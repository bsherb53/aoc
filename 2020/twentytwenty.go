package twentytwenty

import (
	"fmt"
	"time"

	ten "github.com/aoc/2020/10_ten"
	eleven "github.com/aoc/2020/11_eleven"
	twelve "github.com/aoc/2020/12_twelve"
	thirteen "github.com/aoc/2020/13_thirteen"
	fourteen "github.com/aoc/2020/14_fourteen"
	fifteen "github.com/aoc/2020/15_fifteen"
	sixteen "github.com/aoc/2020/16_sixteen"
	seventeen "github.com/aoc/2020/17_seventeen"
	eighteen "github.com/aoc/2020/18_eighteen"
	nineteen "github.com/aoc/2020/19_nineteen"
	one "github.com/aoc/2020/1_one"
	twenty "github.com/aoc/2020/20_twenty"
	twentyone "github.com/aoc/2020/21_twentyone"
	twentytwo "github.com/aoc/2020/22_twentytwo"
	twentythree "github.com/aoc/2020/23_twentythree"
	twentyfour "github.com/aoc/2020/24_twentyfour"
	twentyfive "github.com/aoc/2020/25_twentyfive"
	twentysix "github.com/aoc/2020/26_twentysix"
	twentyseven "github.com/aoc/2020/27_twentyseven"
	twentyeight "github.com/aoc/2020/28_twentyeight"
	twentynine "github.com/aoc/2020/29_twentynine"
	two "github.com/aoc/2020/2_two"
	thirty "github.com/aoc/2020/30_thirty"
	thirtyone "github.com/aoc/2020/31_thirtyone"
	three "github.com/aoc/2020/3_three"
	four "github.com/aoc/2020/4_four"
	five "github.com/aoc/2020/5_five"
	six "github.com/aoc/2020/6_six"
	seven "github.com/aoc/2020/7_seven"
	eight "github.com/aoc/2020/8_eight"
	nine "github.com/aoc/2020/9_nine"
)

var allDays = []func(l bool){
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

func Run2020(day, times int) {
	fmt.Println("Advent Of Code 2020")
	for i := range allDays {
		if i < day {
			log := i+1 == day
			if times > 1 {
				log = false
			}
			var avg time.Duration
			for j := 0; j < times; j++ {
				start := time.Now()
				allDays[i](log)
				t := time.Since(start)
				avg += t
			}
			fmt.Println(fmt.Sprintf("Day %v ran %v times with an average of %v", i+1, times, time.Duration(int64(avg)/int64(times))))
		}
	}
}
