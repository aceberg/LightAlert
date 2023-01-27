package check

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TimeToSec checks if time is correct and converts it to seconds
func TimeToSec(timeout string) int64 {
	var seconds int64

	isDays := regexp.MustCompile(`^\d+d$`)
	if isDays.MatchString(timeout) {

		numStr := strings.TrimSuffix(timeout, "d")

		days, err := strconv.ParseInt(numStr, 10, 64)
		IfError(err)

		seconds = (days * 24 * 60 * 60)

	} else {

		t, err := time.ParseDuration(timeout)
		IfError(err)

		seconds = int64(t.Seconds())
	}

	return seconds
}
