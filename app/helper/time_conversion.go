package helper

import (
	"fmt"
	"time"
)

func TimeConversion() {

	target := time.UnixMilli(1591502343000)
	dur := time.Since(target)

	sec := dur.Seconds()
	minute := dur.Minutes()
	hour := dur.Hours()

	switch {
	case sec < 60:
		fmt.Println(dur.Seconds(), "second")
	case minute < 60:
		fmt.Println(dur.Minutes(), "minute")
	case hour < 24:
		fmt.Println(dur.Hours(), "hours")
	case hour < 730:
		fmt.Println(dur.Hours()/24, "day")
	case hour < 8760:
		fmt.Println(dur.Hours()/730, "month")
	default:
		fmt.Println(dur.Hours()/8760, "year")
	}

}
