package main

import (
	"fmt"
	//"log"
	//"time"
	"sync"

	//"github.com/golang/protobuf/proto"
	//"github.com/golang/protobuf/ptypes"
	pb "github.com/imorph/proto-calendar/pkg/cal"
)

type Calendar struct {
	mx   sync.RWMutex
	pCal *pb.Calendar
}

func NewCalendar() *Calendar {
	cal := &Calendar{
		pCal: &pb.Calendar{},
	}
	return cal
}

func (c *Calendar) AddEvent(e *pb.Event) {
	c.mx.Lock()
	c.pCal.Events = append(c.pCal.Events, e)
	c.mx.Unlock()
}

func (c *Calendar) UpdateEventByName(e *pb.Event) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for i, ev := range c.pCal.Events {
		if ev.GetName() == e.GetName() {
			c.pCal.Events[i] = e
			return
		}
	}
}

func (c *Calendar) DeleteEventByName(name string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for i, ev := range c.pCal.Events {
		if ev.GetName() == name {
			c.pCal.Events[i] = c.pCal.Events[len(c.pCal.Events)-1]
			c.pCal.Events[len(c.pCal.Events)-1] = nil
			c.pCal.Events = c.pCal.Events[:len(c.pCal.Events)-1]
			return
		}
	}
}

func (c *Calendar) PrintEventNames() {
	fmt.Println("Events that you have:")
	c.mx.RLock()
	for _, ev := range c.pCal.Events {
		fmt.Println("Event:")
		fmt.Println("  Name:", ev.GetName())
		fmt.Println("  Type:", ev.GetType())
	}
	c.mx.RUnlock()
}
