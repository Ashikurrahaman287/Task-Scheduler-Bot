package main

import "time"

// Task represents a scheduled task.
type Task struct {
	ID        int
	Name      string
	Schedule  time.Time
	Interval  time.Duration
	Recurring bool
	Priority  int
	Notify    bool
	Done      bool
}
