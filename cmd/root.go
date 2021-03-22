package cmd

import (
	"fmt"

	"github.com/YangzhenZhao/goquotes/quotes/stock"
	"github.com/spf13/cobra"
)

var tickCode string
var priceCode string
var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if tickCode != "" {
			quote := stock.SinaQuote{}
			tick, err := quote.Tick(tickCode)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%+v\n", tick)
			}
		}
		if priceCode != "" {
			quote := stock.SinaQuote{}
			price, err := quote.Price(priceCode)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(price)
			}
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&tickCode, "tick", "t", "", "Set a code to get tick.")
	rootCmd.Flags().StringVarP(&priceCode, "price", "p", "", "Set a code to get price.")
}
