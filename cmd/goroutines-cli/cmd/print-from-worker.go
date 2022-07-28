package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goroutine-cli/cmd/goroutines-cli/tooling"
	"time"
)

var printWorkerCmd = &cobra.Command{
	Use:   "print-from-worker",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		printFromWorker(args[0], tooling.ToInt(args[1]))
	},
}

func init() {
	rootCmd.AddCommand(printWorkerCmd)
}

func printFromWorker(taskName string, workerCount int) {

	// Create the results as a tally of what task is done, so it can be emptied as a waitgroup
	results := make(chan int, 1)

	// Here we start up the right number of goroutines
	for i := 1; i <= workerCount; i++ {
		go worker(i, taskName, results)
	}

	// We empty results when each task has been completed
	// This acts as a way to wait for the goroutines to finish before closing the application
	for a := 1; a <= workerCount; a++ {
		<-results
	}
}

// This syntax for tasks (the second parameter) is like a continuously open gate
func worker(id int, task string, results chan<- int) {

	// Because tasks is a channel, you can iterate over it
	fmt.Println("worker: ", id, " has begun task: ", task)

	// This simulates a difficult task
	time.Sleep(time.Second)
	fmt.Println("worker: ", id, " has finished task: ", task)

	// This is to fill up our results channel so we can  "dump" it to act as a waitgroup
	results <- 1

}
