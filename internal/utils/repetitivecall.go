package utils

import "time"

func RepetitiveCall(intervalInSeconds int, callback func()) func() {
	return func() {
		for {
			time.Sleep(time.Duration(intervalInSeconds) * time.Second)
			callback()
		}
	}
}
