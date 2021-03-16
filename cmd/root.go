package cmd

import (
	"github.com/YangzhenZhao/goquotes/quotes/stock"
	"github.com/spf13/cobra"
)

var mode int8
var code string
var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		println(mode)
		if code != "" {
			quote := stock.SinaQuote{}
			quote.Tick(code)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "")
	rootCmd.Flags().StringVarP(&code, "code", "c", "", "")
}
