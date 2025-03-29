package utils

import "time"

// Takes interval in seconds and callback
// as input, returns callback which will
// endlessly fire this callback with given interval.
func InfiniteRepetitiveCall(intervalInSeconds uint, callback func()) func() {
	return func() {
		for {
			time.Sleep(time.Duration(intervalInSeconds) * time.Second)
			callback()
		}
	}
}
