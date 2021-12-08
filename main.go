package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var program = []string{
	"08:55AM",
	"09:40AM",
	"10:25AM",
	"11:10AM",
	"11:55AM",
	"01:15PM",
	"02:00PM",
	"02:45PM",
	"03:30PM",
}

const cyan = "\u001b[36;1m"
const reset = "\u001b[0m"

func fmtTime(mins string, secs string) string {
	if len(mins) < 2 {
		mins = "0" + mins
	}
	if len(secs) < 2 {
		secs = "0" + secs
	}
	return mins + ":" + secs
}

func parseTime(t time.Duration) string {
	mins := strconv.Itoa(int(math.Floor(t.Minutes())))
	secs := strconv.Itoa(int(int(math.Round(t.Seconds())) % 60))
	return cyan + "time left: " + reset + fmtTime(mins, secs)
}

func doEvery(d time.Duration, f func()) {
	f()
	for range time.Tick(d) {
		f()
	}
}

func main() {
	current := time.Now()
	for i := range program {
		lessonTime, _ := time.Parse(time.Kitchen, program[i])
		lessonTime = time.Date(
			current.Year(),
			current.Month(),
			current.Day(),
			lessonTime.Hour(),
			lessonTime.Minute(),
			lessonTime.Second(),
			current.Nanosecond(),
			current.Location())
		isAfter := (lessonTime.After(current))
		if isAfter {
			fmt.Println()
			fmt.Println(cyan+"ends on:"+reset, program[i])
			doEvery(time.Second, func() {
				fmt.Printf("\r%0s", parseTime(lessonTime.Sub(time.Now()))+" ")
			})
			break
		}
	}
}
