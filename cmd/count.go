package cmd

import (
	"fmt"

	"github.com/vulski/dong/core"
	"github.com/spf13/cobra"
)

func init() {
	dongCmd.AddCommand(countCmd)
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Print the amount of dongs you have stored!",
	Long:  `Print the amount of dongs you have stored!`,
	Run: func(cmd *cobra.Command, args []string) {
		var count int
		db.Model(&core.Dong{}).Count(&count)
		fmt.Println(count)
	},
}
