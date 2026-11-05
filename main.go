package main

import (
	"fmt"
	"time"
)

// Action is the core functionality hook. It should be changed by the user.
func runAction() {
	// Initial functionality: print "hello"
	fmt.Println("hello")

	// Future actions will be implemented here.
}

func main() {
	// 1. Initial run
	fmt.Println("--- Chappie Initial Run ---")
	runAction()

	// 2. Set up the hourly loop until the end of the day
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// The target stop time (00:00 of the next day)
	// We calculate the date for tomorrow at midnight.
	tomorrow := time.Now().Add(24 * time.Hour)
	midnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.Local)

	fmt.Printf("--- Chappie Entering Hourly Loop. Target stop time: %s ---\n", midnight.Format(time.Stamp))

	// Run the first action immediately
	runAction()

	for {
		select {
		case <-ticker.C:
			// Check if we have reached or passed midnight
			if time.Now().After(midnight) || time.Now().Equal(midnight) {
				fmt.Println("--- End of Day. Shutting down Chappie. ---")
				return
			}

			fmt.Println("--- Hourly Checkpoint Reached ---")
			runAction()
		default:
			// Sleep briefly to prevent 100% CPU usage if the ticker hasn't fired yet
			time.Sleep(1 * time.Second)
		}
	}
}