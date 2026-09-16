package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var iso8601DurationPattern = regexp.MustCompile(
	`^P(?:(\d+(?:\.\d+)?)D)?(?:T(?:(\d+(?:\.\d+)?)H)?(?:(\d+(?:\.\d+)?)M)?(?:(\d+(?:\.\d+)?)S)?)?$`,
)

// parseISO8601Duration parses ISO 8601 durations such as PT5S, PT1M30S and P1DT2H.
// Years and months are intentionally unsupported because their length is variable.
func parseISO8601Duration(value string) (time.Duration, error) {
	matches := iso8601DurationPattern.FindStringSubmatch(value)
	if matches == nil || matches[1] == "" && matches[2] == "" && matches[3] == "" && matches[4] == "" {
		return 0, fmt.Errorf("must be a non-zero ISO 8601 duration, for example PT5S")
	}

	units := []struct {
		value  string
		factor float64
	}{
		{matches[1], 24 * 60 * 60},
		{matches[2], 60 * 60},
		{matches[3], 60},
		{matches[4], 1},
	}

	var seconds float64
	for _, unit := range units {
		if unit.value == "" {
			continue
		}

		amount, err := strconv.ParseFloat(unit.value, 64)
		if err != nil {
			return 0, fmt.Errorf("parse %q: %w", value, err)
		}
		seconds += amount * unit.factor
	}

	if seconds <= 0 || seconds > float64(1<<63-1)/float64(time.Second) {
		return 0, fmt.Errorf("duration %q is outside the supported range", value)
	}

	return time.Duration(seconds * float64(time.Second)), nil
}
