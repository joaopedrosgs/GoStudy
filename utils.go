package main

import "time"

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}
