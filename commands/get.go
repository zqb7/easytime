package commands

import (
	"fmt"

	"github.com/qilook/easytime/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "转换时间",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := utils.Get(args[0])
		if err != nil {
			return err
		}
		PARSE_TIME = t
		ShiftRun(cmd)
		fmt.Println(utils.TimeStd(PARSE_TIME))
		return nil
	},
}
