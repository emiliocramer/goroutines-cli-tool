package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "An in depth description for what goroutines are.",
	Run: func(cmd *cobra.Command, args []string) {
		printDescription()
	},
}

func init() {
	rootCmd.AddCommand(descriptionCmd)
}

func printDescription() {
	fmt.Println(`
	A goroutine is a lightweight execution thread in the Go programming language and a function that executes 
	concurrently with the rest of the program.

	Goroutines are incredibly cheap when compared to traditional threads as the overhead of creating a goroutine 
	is very low. Therefore, they are widely used in Go for concurrent programming.	
	`)
}
