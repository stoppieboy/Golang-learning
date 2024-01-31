package main

import "fmt"

func main() {
	// Create new tasks
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // No. of workers that can run at a time
	}

	// Run the pool
	wp.Run()
	fmt.Println("All tasks have been processed")
}