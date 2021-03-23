package cmd

import (
	"fmt"

	"github.com/YangzhenZhao/goquotes/quotes/stock"
	"github.com/spf13/cobra"
)

var tickCode string
var currentPriceCode string
var tickCodes []string
var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if tickCode != "" {
			quote := stock.SinaQuote{}
			tick, err := quote.Tick(tickCode)
			if err != nil {
				fmt.Println(err)
			} else {
				tick.Print()
			}
		}
		if currentPriceCode != "" {
			quote := stock.SinaQuote{}
			price, err := quote.Price(currentPriceCode)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%.2f\n", price)
			}
		}
		if len(tickCodes) > 0 {
			quote := stock.SinaQuote{}
			tickMap := quote.TickMap(tickCodes)
			for _, tick := range tickMap {
				tick.Print()
				println()
			}
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&tickCode, "tick", "t", "", "Set a code to get tick.")
	rootCmd.Flags().StringVarP(&currentPriceCode, "current_price", "c", "", "Set a code to get current price.")
	rootCmd.Flags().StringArrayVarP(&tickCodes, "ticks", "T", []string{}, "Set some code to get tick.")
}
