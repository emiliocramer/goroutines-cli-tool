package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goroutine-cli/cmd/goroutines-cli/tooling"
	"time"
)

var printRateCmd = &cobra.Command{
	Use:   "print-rate-limited",
	Short: "Will print your [first-arg], [second-arg] amount of times, with [third-arg] seconds in between each iteration.",
	Long:  "This utilizes how to limit the amount of requests coming in using 'rate-limiting' of some kind.",
	Run: func(cmd *cobra.Command, args []string) {
		printRate(args[0], tooling.ToInt(args[1]), tooling.ToInt(args[2]))
	},
}

func init() {
	rootCmd.AddCommand(printRateCmd)
}

func printRate(text string, iters int, delay int) {

	// Sets the delay between each iteration
	limiter := time.Tick(time.Duration(delay) * time.Second)

	// Iterating through the amount of times we want it printed
	// But waiting for the ticker on each iteration
	// Limiter is a channel that we are waiting for it to be "dump-able"
	for i := 0; i < iters; i++ {
		<-limiter
		fmt.Println(text)
	}

}
