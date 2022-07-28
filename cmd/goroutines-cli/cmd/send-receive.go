package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var alternatingCmd = &cobra.Command{
	Use:   "send-receive",
	Short: "Will alternate pass its parameters between functions with channels. For as many parameters as you pass it will iterate through them",
	Long:  "This utilizes having data passed between functions and having channels as function parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		sender1 := sender(args)

		receiver(sender1)
	},
}

func init() {
	rootCmd.AddCommand(alternatingCmd)
}

// will return "the insertion of a string channel" (this is what you get when you call the function)

// when this is called, it'll initiate a go function that will be iterating through the params in concurrency
// passing it through, through each iteration

// the for loop will continue passing and returning a new version of ch? i guess?
// while the close isn't fulfilled until after it's iterated through the whole params list

func sender(params []string) <-chan string {
	ch := make(chan string)

	go func() {
		for _, str := range params {
			ch <- str
		}
		close(ch)
	}()

	return ch
}

// will print on a new line each value in ch, which continues to change?
// takes in not a channel, but the input of a channel

func receiver(ch <-chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}
