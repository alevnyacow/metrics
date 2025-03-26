package utils

import "time"

func InfiniteRepetitiveCall(intervalInSeconds uint, callback func()) func() {
	return func() {
		for {
			time.Sleep(time.Duration(intervalInSeconds) * time.Second)
			callback()
		}
	}
}
