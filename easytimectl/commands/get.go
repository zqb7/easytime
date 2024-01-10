package commands

import (
	"fmt"

	"github.com/qilook/easytime"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := easytime.Get(args[0])
		fmt.Println(t, err)
	},
}
