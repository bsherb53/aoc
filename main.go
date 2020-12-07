package main

import (
	"time"

	twentytwenty "github.com/aoc/2020"
)

func main() {
	day := time.Now().Day()
	day = 0
	twentytwenty.Run2020(day, 1000)
}
