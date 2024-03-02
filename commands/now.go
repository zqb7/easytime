package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/zqb7/easytime/pkg/utils"
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
