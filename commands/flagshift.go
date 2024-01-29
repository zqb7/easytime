package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var flagShiftDay = &pflag.FlagSet{}
var flagShiftHour = &pflag.FlagSet{}
var flagShiftSecond = &pflag.FlagSet{}
var flagShiftUTC = &pflag.FlagSet{}

var (
	shiftDay, shiftHour, shiftSecond *int
	shiftUTC                         *string
)

func init() {
	shiftDay = flagShiftDay.Int("day", 0, "--day -1")
	shiftHour = flagShiftHour.Int("hour", 0, "--hour 2 or --hour -3")
	shiftSecond = flagShiftSecond.Int("second", 0, "--second 20 or --second -30")
	shiftUTC = flagShiftUTC.String("utc", "", "--utc 8 or --utc -8")
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
	if *shiftUTC != "" {
		offset, _ := strconv.Atoi(*shiftUTC)
		for _, tzone := range TimeZones {
			// 注意：Etc/GMT+5 实际上是 UTC -05:00
			if tzone.Offset == -offset {
				PARSE_TIME = PARSE_TIME.In(tzone.Zone)
				break
			}
		}
	} else {
		PARSE_TIME = PARSE_TIME.In(time.Local)
	}
}

func init() {
	nowCmd.PersistentFlags().AddFlagSet(flagShiftDay)
	getCmd.PersistentFlags().AddFlagSet(flagShiftDay)
	nowCmd.PersistentFlags().AddFlagSet(flagShiftHour)
	getCmd.PersistentFlags().AddFlagSet(flagShiftHour)
	nowCmd.PersistentFlags().AddFlagSet(flagShiftSecond)
	getCmd.PersistentFlags().AddFlagSet(flagShiftSecond)
	nowCmd.PersistentFlags().AddFlagSet(flagShiftUTC)
	getCmd.PersistentFlags().AddFlagSet(flagShiftUTC)
}

type TimeZone struct {
	Offset int
	Zone   *time.Location
}

var TimeZones []*TimeZone

func init() {
	for i := -9; i <= 9; i++ {
		var name string
		if i >= 0 {
			name = fmt.Sprintf("Etc/GMT+%d", i)
		} else {
			name = fmt.Sprintf("Etc/GMT%d", i)
		}
		v, err := time.LoadLocation(name)
		if err != nil {
			panic(err)
		}
		TimeZones = append(TimeZones, &TimeZone{Offset: i, Zone: v})
	}
}
