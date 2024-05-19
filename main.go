package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize the scheduler
	scheduler := Scheduler{}

	// Example tasks
	scheduler.AddTask("Task 1", time.Now().Add(5*time.Second), 0, false, 1, true)
	scheduler.AddTask("Task 2", time.Now().Add(10*time.Second), 2*time.Second, true, 2, true)

	// Start the scheduler
	go scheduler.Run()

	// Handle termination signals for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Logging
	logger := log.New(os.Stdout, "[Task Scheduler] ", log.LstdFlags)

	logger.Println("Task Scheduler started.")

	// Keep the application running until terminated
	select {
	case sig := <-signalCh:
		logger.Printf("Received signal %v. Shutting down...\n", sig)
		// Perform cleanup or graceful shutdown here
		// For example, stop accepting new tasks and wait for ongoing tasks to complete
		logger.Println("Task Scheduler stopped.")
	}
}
