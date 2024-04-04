package main

import (
	"APITOOLS/Controller"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Initialize controllers
	emailCtrl := &Controller.EmailController{}
	cacheCtrl := Controller.NewCacheController()
	taskCtrl := &Controller.TaskController{
		EmailCtrl: emailCtrl,
		CacheCtrl: cacheCtrl,
	}

	// Initialize SchedulerController with Interval field
	schedulerCtrl := Controller.NewSchedulerController(5 * time.Second)

	// Run SchedulerController in a goroutine
	go func() {
		for {
			time.Sleep(schedulerCtrl.Interval)
			fmt.Println("Scheduler dijalankan pada:", time.Now())
			taskCtrl.SubmitTaskNotification("if-22008@students.ithb.ac.id", "Tugas berkala")
		}
	}()

	// Endpoint for receiving task requests via HTTP POST
	http.HandleFunc("/submit-task", func(w http.ResponseWriter, r *http.Request) {
		// Get taskName and email from the request
		taskName := r.FormValue("tugas")
		email := r.FormValue("email")

		// Set task to cache and submit task notification
		cacheCtrl.Set(email, taskName)
		taskCtrl.SubmitTaskNotification(email, taskName)

		// Respond with success message
		fmt.Fprintf(w, "Tugas '%s' telah berhasil diajukan kepada %s", taskName, email)
	})

	// Start HTTP server
	fmt.Println("Server HTTP berjalan di port 8080...")
	http.ListenAndServe(":8080", nil)
}
