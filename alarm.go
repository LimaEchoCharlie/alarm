package alarm

import (
	"fmt"
	"time"
)

// Time represents a 24 hour time %d:%d:%d with second precision
type Time struct {
	Hour   int
	Minute int
	Second int
}

// equalTime returns true is Time is equivalent to the standard GO time
func (a Time) equalTime(t time.Time) bool {
	return a.Hour == t.Hour() && a.Minute == t.Minute() && a.Second == t.Second()
}

// String satisfies the stringer interface
func (a Time) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", a.Hour, a.Minute, a.Second)
}

// NewTime checks the inputs and initialises a new Time
func NewTime(h, m, s int) (Time, error) {
	var t Time
	if h < 0 || h >= 24 {
		return t, fmt.Errorf("Invalid hour %d", h)
	}
	if m < 0 || m >= 60 {
		return t, fmt.Errorf("Invalid minute %d", h)
	}
	if s < 0 || s >= 60 {
		return t, fmt.Errorf("Invalid seconds %d", h)
	}
	return Time{h, m, s}, nil
}

// TimeFromStandardTime creates a Time from the given standard GO time
func TimeFromStandardTime(t time.Time) Time {
	return Time{t.Hour(), t.Minute(), t.Second()}
}

// Alarm sounds (returns ticks on a channel) if the current time matches any of the times provided ar initialisation.
// Alarm is a wrapper around a time.Ticker. Depending on the accuracy of the tick, multiple ticks can be sent within the
// second when the alarm is sounding
type Alarm struct {
	C      chan Time // The channel on which the ticks are delivered.
	ticker *time.Ticker
}

// Stop stops the Alarm
func (a *Alarm) Stop() {
	a.ticker.Stop()
}

// NewAlarm creates a new Alarm
func NewAlarm(accuracy time.Duration, alarmTimes ...Time) *Alarm {
	a := &Alarm{
		C:      make(chan Time, 1),
		ticker: time.NewTicker(accuracy),
	}
	go func() {
		for t := range a.ticker.C {
			for _, alarmTime := range alarmTimes {
				if alarmTime.equalTime(t) {
					a.C <- alarmTime
				}
			}
		}
	}()
	return a
}
