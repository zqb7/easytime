package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	PARSE_TIME = time.Time{}
)

var rootCmd = &cobra.Command{
	Use:   "easytime",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
