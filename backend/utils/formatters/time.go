package formatters

import "time"

func GetTimestap() int64 {
	return time.Now().Unix()
}
