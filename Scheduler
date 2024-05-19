package main

import (
	"fmt"
	"time"
)

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

// Scheduler manages scheduled tasks.
type Scheduler struct {
	Tasks []Task
}

// AddTask adds a new task to the scheduler.
func (s *Scheduler) AddTask(name string, schedule time.Time, interval time.Duration, recurring bool, priority int, notify bool) {
	task := Task{
		ID:        len(s.Tasks) + 1,
		Name:      name,
		Schedule:  schedule,
		Interval:  interval,
		Recurring: recurring,
		Priority:  priority,
		Notify:    notify,
		Done:      false,
	}
	s.Tasks = append(s.Tasks, task)
}

// Run starts the scheduler and executes tasks.
func (s *Scheduler) Run() {
	for {
		now := time.Now()
		for i, task := range s.Tasks {
			if !task.Done && task.Schedule.Before(now) {
				fmt.Printf("Executing task: %s\n", task.Name)
				s.Tasks[i].Done = true
				if task.Notify {
					// Add notification logic here
					fmt.Printf("Notification: Task %s completed\n", task.Name)
				}
				if task.Recurring {
					s.Tasks[i].Schedule = now.Add(task.Interval)
					s.Tasks[i].Done = false
					fmt.Printf("Rescheduled recurring task: %s\n", task.Name)
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}
