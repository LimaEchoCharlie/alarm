package alarm

import (
	"fmt"
	"testing"
	"time"
)

// Test checks that NewTime creates the right time when given valid values
func TestValidTime(t *testing.T) {
	testCases := []struct {
		h, m, s int
	}{
		{1, 13, 16},
		{18, 1, 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc), func(t *testing.T) {
			alarmTime, err := NewTime(tc.h, tc.m, tc.s)
			if err != nil {
				t.Fatalf("Unexpected err: %s", err)
			}
			expected := Time{tc.h, tc.m, tc.s}
			if alarmTime != expected {
				t.Errorf("Time mismatch")
			}
		})
	}
}

// Test that NewTime raises an error if given a value outside of usual time ranges
func TestInvalidTime(t *testing.T) {
	testCases := []struct {
		h, m, s int
	}{
		{-1, 13, 16},
		{24, 13, 16},
		{12, -1, 16},
		{12, 60, 16},
		{12, 13, -1},
		{12, 13, 60},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc), func(t *testing.T) {

			if _, err := NewTime(tc.h, tc.m, tc.s); err == nil {
				t.Errorf("Expected an error for invalid time %02d:%02d:%02d", tc.h, tc.m, tc.s)
			}
		})
	}
}

// Test that the alarm time can be create from a standard GO time
func TestTimeFromStandardTime(t *testing.T) {
	testCases := []struct {
		h, m, s int
	}{
		{1, 13, 16},
		{18, 1, 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc), func(t *testing.T) {
			standardTime := time.Date(2018, time.December, 25, tc.h, tc.m, tc.s, 0, time.UTC)
			alarmTime := TimeFromStandardTime(standardTime)
			expected := Time{tc.h, tc.m, tc.s}
			if alarmTime != expected {
				t.Errorf("Time mismatch")
			}
		})
	}
}

// Test that the alarm stops
func TestAlarmStops(t *testing.T) {
	delta := 10 * time.Millisecond
	now := time.Now()
	a := NewAlarm(delta, TimeFromStandardTime(now), TimeFromStandardTime(now.Add(time.Second)))
	a.Stop()
	time.Sleep(2 * delta)
	select {
	case <-a.C:
		t.Fatal("Alarm did not shut down")
	default:
		// ok
	}
}

// Test that the alarm sounds
func TestAlarmSounds(t *testing.T) {
	delta := 10 * time.Millisecond
	now := time.Now()
	a := NewAlarm(delta, TimeFromStandardTime(now), TimeFromStandardTime(now.Add(time.Second)))
	count := 0
	go func() {
		for _ = range a.C {
			count++
		}
	}()
	time.Sleep(5 * delta)
	a.Stop()
	if count == 0 {
		t.Fatal("Alarm did not sound")
	}
}
