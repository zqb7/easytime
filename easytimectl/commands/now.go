package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nowCmd)
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "输出当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now().Unix())
	},
}
