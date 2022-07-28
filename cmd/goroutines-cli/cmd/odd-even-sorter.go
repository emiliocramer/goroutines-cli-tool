package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goroutine-cli/cmd/goroutines-cli/tooling"
)

var oddEvenCmd = &cobra.Command{
	Use:   "odd-even-sorter",
	Short: "This command will sort whatever int parameters you pass with it into odd and even numbers",
	Long:  "This teaches us to have two concurrent goroutines and channels (sorting a bucket)",
	Run: func(cmd *cobra.Command, args []string) {

		// Create the channels that will flow
		// In this case we want to have two different channels

		chOdd := make(chan int)
		chEven := make(chan int)

		// We run our functions in concurrency and pass our empty channels to them

		go odd(chOdd)
		go even(chEven)

		// We go down the list of our argument and as it sorts the integers it passes the value into it's respective channel
		// It can pass it into the goroutines that are being run

		for _, v := range args {
			v := tooling.ToInt(v)

			if v%2 == 0 {
				chEven <- v
			} else {
				chOdd <- v
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(oddEvenCmd)
}

// This is essentially a print function for when it is odd
// But what makes it special is that as its running, it's always open because it's parameter is "open"

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("ODD: ", v)
	}
}

// Same thing but for when the value is even

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("EVEN: ", v)
	}
}
