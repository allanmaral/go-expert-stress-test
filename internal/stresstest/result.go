package stresstest

import "time"

type Result struct {
	Code     int
	Duration time.Duration
	Error    string
}
