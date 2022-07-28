package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goroutine-cli/cmd/goroutines-cli/tooling"
)

// Done

var passLastCmd = &cobra.Command{
	Use:   "pass-last",
	Short: "Enter the amount of goroutines you want that pass the last iteration through a channel.",
	Long:  "This utilizes a single buffered channel being iterated through a list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		last := make(chan int, 1)
		cap := tooling.ToInt(args[0])

		go passLast(cap, last)
		for i := 1; i <= cap; i++ {
			fmt.Println("Last: ", <-last, ". Current: ", i)
		}
	},
}

func init() {
	rootCmd.AddCommand(passLastCmd)
}

func passLast(cap int, last chan int) {
	defer close(last)
	for i := 0; i < cap; i++ {
		last <- i
	}
}
