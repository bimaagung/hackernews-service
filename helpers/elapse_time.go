package helpers

import (
	"fmt"
	"math"
	"time"
)

func ElapseTime(date float64) string {
	currentTime := time.Now()
	timestamp :=  float64(currentTime.UnixNano() / int64(time.Millisecond))
	delta := math.Abs(timestamp / 1000 - date)
	days := int(math.Floor(delta / 86400))

	if days != 0 {
		return fmt.Sprintf("%d days", days)
	}

	delta = delta - float64(days) * 86400
	hours := int(math.Floor(delta / 3600)) % 24

	if hours != 0 {
		return fmt.Sprintf("%d days", days)
	}

	return ""
}