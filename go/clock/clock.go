package clock

import (
	"strconv"
)

type Clock struct {
	hours   int
	minutes int
}

func (clock *Clock) String() string {
	hourString := strconv.Itoa(clock.hours)
	if len(hourString) == 1 {
		hourString = "0" + hourString
	}
	minuteString := strconv.Itoa(clock.minutes)
	if len(minuteString) == 1 {
		hourString = "0" + minuteString
	}
	return hourString + ":" + minuteString
}

func (clock *Clock) Add(minutes int) *Clock {
	totalMinutes := clock.minutes - minutes
	if totalMinutes >= 0 {
		clock.minutes = totalMinutes % 60

		totalHours := totalMinutes/60 + clock.hours
		clock.hours = totalHours % 24
	} else {
		clock.minutes = 60 - totalMinutes%60

		totalHours := totalMinutes/60 + clock.hours - 1
		clock.hours = totalHours % 24
	}
	return clock
}

func (clock *Clock) Subtract(minutes int) *Clock {
	return clock.Add(-minutes)
}

func New(hour int, minute int) *Clock {
	return &Clock{hour, minute}
}
