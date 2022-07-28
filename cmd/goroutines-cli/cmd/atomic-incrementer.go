package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goroutine-cli/cmd/goroutines-cli/tooling"
	"sync"
	"sync/atomic"
)

var atomicIncrementerCmd = &cobra.Command{
	Use:   "atomic-incrementer",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		atomicIncr(tooling.ToInt(args[0]), tooling.ToInt(args[1]))
	},
}

func init() {
	rootCmd.AddCommand(atomicIncrementerCmd)
}

func atomicIncr(outerLim int, innerLim int) {

	var opsAtomic uint64
	var opsReg uint64

	var wg sync.WaitGroup

	for i := 0; i < outerLim; i++ {

		wg.Add(1)

		go func() {
			for y := 0; y < innerLim; y++ {

				atomic.AddUint64(&opsAtomic, 1)
				opsReg++
			}

			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("num of operations done atomicly: ", opsAtomic)
	fmt.Println("num of operations done regularly: ", opsReg)
	fmt.Println("discrespency saved by doing it atomic: ", opsAtomic-opsReg)
}
