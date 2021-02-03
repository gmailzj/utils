package utils

import "time"

func LocalTime() time.Time {
	return time.Now().In(time.FixedZone("CST", 8*3600))
}

func TimeSecond() int64 {
	return LocalTime().Unix()
}

func TimeMillisecond() int64 {
	return LocalTime().UnixNano() / 1e6
}

func Location() *time.Location {
	return time.FixedZone("CST", 8*3600)
}

func TimeSecondStr() string {
	return LocalTime().Format("20060102150405")
}
