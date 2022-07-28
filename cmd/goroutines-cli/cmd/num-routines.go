package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

// Done

var numRoutinesCmd = &cobra.Command{
	Use:   "num-routines",
	Short: "Enter the amount of goroutines you'd like to run.",
	Long:  "This utilizes a single goroutine being iterated through a user determined parameter.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		go numRoutines(args[0])

		time.Sleep(time.Second)
	},
}

func init() {
	rootCmd.AddCommand(numRoutinesCmd)
}

func numRoutines(cap string) {

	intCap, err := strconv.Atoi(cap)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= intCap; i++ {
		fmt.Println("goroutine number: ", i)
	}
}
