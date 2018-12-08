# alarm
Alarm is a wrapper around a time.Ticker that sounds (returns ticks on a channel) if the current time matches any of the 
times provided at initialisation. Alarms are set using a 24 hour clock time 23:59:59 with second precision. Unlike
the definition of time in the standard go library, alarm times don't have a date component and so will sound at the same time each day.

Depending on the accuracy of the tick, multiple ticks can be sent within the
second when the alarm is sounding
## Example
~~~go
import (
	"fmt"
	"github.com/limaechocharlie/alarm"
	"time"
)

// create as many alarm times as required
on := alarm.Time{Hour:21, Minute:0, Second:0}
off := alarm.Time{Hour:21, Minute:0, Second:30}

// create the alarm
a := alarm.NewAlarm(200*time.Millisecond, on, off)

go func() {
	
	// wait for the alarm to fire
	for currentAlarm := range a.C {
		
		// choose an action depending on which alarm
		if currentAlarm == on {
			
			// the alarm may fire >1 in a second.
			// If it is important that the action 
			// occurs only once, add a filter
			fmt.Println(p, "ON:", currentAlarm)
		} else if currentAlarm == off {
			
			fmt.Println(p, "OFF:", currentAlarm)
			}
		}
	}
}()

time.Sleep(2 * time.Minute)

// if required, stop the alarm from another process
a.Stop()
fmt.Println("Alarm stopped")
~~~
