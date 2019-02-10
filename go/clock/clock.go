package clock

import "fmt"

type Clock int

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}

func (clock Clock) Add(minutes int) Clock {
	time := (clock + Clock(minutes)) % 1440
	if time < 0 {
		time += 1440
	}
	return Clock(time)
}

func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}

func New(hour int, minute int) Clock {
	time := (hour*60 + minute) % 1440
	if time < 0 {
		time += 1440
	}
	return Clock(time)
}
