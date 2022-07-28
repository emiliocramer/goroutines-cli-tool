package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
)

var writeBackwardsCmd = &cobra.Command{
	Use:   "write-backwards",
	Short: "Will print out all of your inputs in new lines, then it will print the reverse of the words in a new line",
	Long:  "This teaches to use wait-groups to wait until all tasks are completed",
	Run: func(cmd *cobra.Command, args []string) {

		// Define a new waitgroup
		wg := new(sync.WaitGroup)

		// Iterate through the strings that we want reversed
		for _, v := range args {

			// For each one we add a new waitgroup instance
			wg.Add(1)

			// Apparently this is needed
			// Re-define the index in what we are iterating through
			v := v

			// Run our function
			// Defer closing our waitgroup, so this will have many running at the same time
			// As they finish, so does the waitgroup
			go func() {
				defer wg.Done()
				forwardBack(v)
			}()
		}

		// Will wait until all waitgroups that we've "added" are finished
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(writeBackwardsCmd)
}

func forwardBack(str string) {

	// Print it straight out
	fmt.Println(str)

	// Create an empty array of strings of the right size to fill with the reverse
	reverse := make([]string, len(str))

	// iterating through the original string
	// 0...5, bytes of letters
	for i, v := range str {

		// index-reverse of where it is now, relative to the rest of the string
		// Setting it to the string version of the current index(bytes)
		reverse[(len(str)-1)-i] = string(v)
	}

	// iterate through reversed list
	// print each index out on same line
	for _, v := range reverse {
		fmt.Print(v)
	}
	fmt.Println("")
}
