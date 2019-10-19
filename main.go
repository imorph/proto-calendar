//go:generate protoc --go_out=. pkg/cal/calendar.proto
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/imorph/proto-calendar/pkg/cal"
)

func main() {

	fmt.Println("This is simple calendar demo.")

	evStartDate, err := ptypes.TimestampProto(time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC))
	if err != nil {
		log.Println(err)
	}
	evStopDate, err := ptypes.TimestampProto(time.Date(2019, time.November, 10, 23, 5, 0, 0, time.UTC))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Creating 3 events")
	ev1 := pb.Event{
		Name:      "Meeting with Alice ;-)",
		StartTime: evStartDate,
		StopTime:  evStopDate,
		Type:      pb.Event_MEETING,
	}

	ev2 := pb.Event{
		Name:      "Other event",
		StartTime: evStartDate,
		StopTime:  evStopDate,
		Type:      pb.Event_OTHER,
	}

	ev3 := pb.Event{
		Name:      "Remind me",
		StartTime: evStartDate,
		StopTime:  evStopDate,
		Type:      pb.Event_REMINDER,
	}

	fmt.Println("Creating calendar")
	myCalendar := NewCalendar()
	fmt.Println("adding 3 events to calendar")
	myCalendar.AddEvent(&ev1)
	myCalendar.AddEvent(&ev2)
	myCalendar.AddEvent(&ev3)

	myCalendar.PrintEventNames()

	fmt.Println(`Updating "Remind me" Event`)
	ev4 := pb.Event{
		Name:      "Remind me",
		StartTime: evStartDate,
		StopTime:  evStopDate,
		Type:      pb.Event_OTHER,
	}

	myCalendar.UpdateEventByName(&ev4)
	myCalendar.PrintEventNames()

	fmt.Println(`Deleting "Meeting with Alice ;-)" Event`)
	myCalendar.DeleteEventByName("Meeting with Alice ;-)")
	myCalendar.PrintEventNames()

}
