package stresstest

import "math"

type Statistics struct {
	count int64
	sum   float64
	sumSq float64
	min   float64
	max   float64
}

func newStatistics() Statistics {
	return Statistics{
		min: math.Inf(1),
		max: math.Inf(-1),
	}
}

func (s *Statistics) Add(value float64) {
	s.count++
	s.sum += value
	s.sumSq += value * value
	if value < s.min {
		s.min = value
	}
	if value > s.max {
		s.max = value
	}
}

func (s *Statistics) Mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

func (s *Statistics) StdDev() float64 {
	if s.count == 0 {
		return 0
	}
	return math.Sqrt((s.sumSq - ((s.sum * s.sum) / float64(s.count))) / float64(s.count-1))
}
