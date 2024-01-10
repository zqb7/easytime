package commands

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var flagShiftDay = &pflag.FlagSet{}
var flagShiftHour = &pflag.FlagSet{}
var flagShiftSecond = &pflag.FlagSet{}

var (
	shiftDay, shiftHour, shiftSecond *int
)

func init() {
	shiftDay = flagShiftDay.Int("day", 0, "--day -1")
	shiftHour = flagShiftHour.Int("hour", 0, "--hour 2 or --hour -3")
	shiftSecond = flagShiftSecond.Int("second", 0, "--second 20 or --second -30")
}

func ShiftRun(cmd *cobra.Command) {
	if cmd.Flag("day") != nil {
		PARSE_TIME = PARSE_TIME.Add(time.Duration(*shiftDay) * 24 * time.Hour)
	}
	if cmd.Flag("hour") != nil {
		PARSE_TIME = PARSE_TIME.Add(time.Duration(*shiftHour) * time.Hour)
	}
	if cmd.Flag("second") != nil {
		PARSE_TIME = PARSE_TIME.Add(time.Duration(*shiftSecond) * time.Second)
	}
}

func init() {
	nowCmd.PersistentFlags().AddFlagSet(flagShiftDay)
	getCmd.PersistentFlags().AddFlagSet(flagShiftDay)
	nowCmd.PersistentFlags().AddFlagSet(flagShiftHour)
	getCmd.PersistentFlags().AddFlagSet(flagShiftHour)
	nowCmd.PersistentFlags().AddFlagSet(flagShiftSecond)
	getCmd.PersistentFlags().AddFlagSet(flagShiftSecond)
}
