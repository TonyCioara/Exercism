package clock

import "fmt"

// Clock is a struct that keeps track of time
type Clock int

// String presents a clock's time in string format
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}

// Add adds a given number of minutes to a clock
func (clock Clock) Add(minutes int) Clock {
	time := (clock + Clock(minutes)) % 1440
	if time < 0 {
		time += 1440
	}
	return Clock(time)
}

// Subtract subtracts a given number of minutes to a clock
func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}

// New creates a new clock
func New(hour int, minute int) Clock {
	time := (hour*60 + minute) % 1440
	if time < 0 {
		time += 1440
	}
	return Clock(time)
}
