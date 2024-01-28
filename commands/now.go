package commands

import (
	"fmt"
	"time"

	"github.com/qilook/easytime/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nowCmd)
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "输出当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		PARSE_TIME = time.Now()
		ShiftRun(cmd)
		fmt.Println(utils.TimeStd(PARSE_TIME))
	},
}