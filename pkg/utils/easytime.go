package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	ErrInvalidTimeFormat    = errors.New("Invalid time format")
	ErrInvalidTimeComponent = errors.New("Invalid time component provided")
)

var REG_TIME = regexp.MustCompile(`^(\d{2})(?:\:?(\d{2}))?(?:\:?(\d{2}))?(?:([\.\,])(\d+))?$`)

type TimeFormat struct {
	Key   string
	Value string
}

var TimeFormats = []TimeFormat{
	{"YYYY-MM-DD", "2006-01-02"},
	{"YYYY-M-DD", "2006-1-02"},
	{"YYYY-M-D", "2006-1-2"},
	{"YYYY/MM/DD", "2006/01/02"},
	{"YYYY/M/DD", "2006/1/02"},
	{"YYYY/M/D", "2006/1/2"},
	{"YYYY.MM.DD", "2006.01.02"},
	{"YYYY.M.DD", "2006.1.02"},
	{"YYYY.M.D", "2006.1.2"},
	{"YYYYMMDD", "20060102"},
	{"YYYY-DDDD", "2006-0002"},
	{"YYYYDDDD", "20060002"},
	{"YYYY-MM", "2006-01"},
	{"YYYY/MM", "2006/01"},
	{"YYYY.MM", "2006.01"},
	{"YYYY", "2006"},
}

func isTimestamp(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	for _, v := range s {
		if !unicode.IsDigit(v) && v != '.' {
			return 0, false
		}
	}
	timestampFloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}
	var timestamp int64
	subSecParts := strings.Split(s, ".")
	if len(subSecParts) == 1 {
		timestamp = int64(timestampFloat * 1e6)
	} else if len(subSecParts[1]) == 3 {
		timestamp = int64(timestampFloat * 1e3)
	} else if len(subSecParts[1]) == 6 {
		timestamp = int64(timestampFloat * 1e6)
	}
	return timestamp, true
}

func Get(s string) (t time.Time, err error) {
	var hasSpaceDivider = strings.Contains(s, " ")
	var hasT_Divider = strings.Contains(s, "T")
	numSpaces := strings.Count(s, " ")

	if hasSpaceDivider && numSpaces != 1 || (hasT_Divider && numSpaces > 0) {
		return time.Time{}, ErrInvalidTimeFormat
	}
	var hasTime = hasSpaceDivider || hasT_Divider
	var hasTZ bool
	var timeString, TZFormat string
	var formats []string

	for _, v := range TimeFormats {
		formats = append(formats, v.Value)
	}

	if hasTime {
		var v []string
		if hasSpaceDivider {
			v = strings.Split(s, " ")
		} else {
			v = strings.Split(s, "T")
		}
		timeString = v[1]
		timeParts := regexp.MustCompile(`[\+\-Z]`).Split(timeString, 2)
		timeComponents := REG_TIME.FindStringSubmatch(timeParts[0])
		if len(timeComponents) == 0 {
			return time.Time{}, ErrInvalidTimeComponent
		}
		// timeComponents  index=0是正则匹配的完整数据
		var hasSubSeconds, hasSeconds, hasMinutes bool
		// hours, _ = strconv.Atoi(timeComponents[0])

		timeComponents = timeComponents[1:]
		if len(timeComponents) > 1 && timeComponents[1] != "" {
			hasMinutes = true
		}
		if len(timeComponents) > 2 && timeComponents[2] != "" {
			hasSeconds = true
		}
		var subSecondsSep string
		if len(timeComponents) > 3 {
			subSecondsSep = timeComponents[3]
		}
		if len(timeComponents) > 4 && timeComponents[4] != "" {
			hasSubSeconds = true
		}

		hasTZ = len(timeParts) == 2

		var isBasicTimeFormat bool
		if !strings.Contains(timeParts[0], ":") {
			isBasicTimeFormat = true
		}

		if hasTZ && strings.Contains(timeParts[1], ":") {
			if strings.Contains(timeString, "+") {
				TZFormat = "+" + timeParts[1]
			} else {
				TZFormat = "-" + timeParts[1]
			}
		}
		var timeSep = ""
		if !isBasicTimeFormat {
			timeSep = ":"
		}
		if hasSubSeconds {
			timeString = fmt.Sprintf("15%s04%s05%s000000", timeSep, timeSep, subSecondsSep)
		} else if hasSeconds {
			timeString = fmt.Sprintf("15%s04%s05", timeSep, timeSep)
		} else if hasMinutes {
			timeString = fmt.Sprintf("15%s04%s", timeSep, timeSep)
		} else {
			timeString = "15"
		}

		for index := range formats {
			if hasSpaceDivider {
				formats[index] += " " + timeString
			} else {
				formats[index] += "T" + timeString
			}
		}
	}
	if hasTime && hasTZ {
		for index := range formats {
			formats[index] += TZFormat
		}
	}
	var isMatchFormat bool
	for _, f := range formats {
		t, err = time.Parse(f, s)
		if err != nil {
			continue
		}
		isMatchFormat = true
		break
	}
	if !isMatchFormat {
		if timestamp, ok := isTimestamp(s); ok {
			return time.UnixMicro(timestamp), nil
		}
	}
	if t.IsZero() {
		return t, ErrInvalidTimeFormat
	}
	return t, nil
}

func TimeStd(t time.Time) string {
	zone, _ := t.Zone()
	return fmt.Sprintf("date:%s,zone:%s,timestamp:%d", t.Format(time.RFC3339), zone, t.Unix())
}
