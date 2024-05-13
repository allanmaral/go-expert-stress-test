package stresstest

import (
	"fmt"
	"strings"
	"time"
)

type Report struct {
	totalTime float64
	requests  int64
	codes     map[int]int64
	errors    map[string]int64
	latency   Statistics
}

func newReport() Report {
	return Report{
		codes:   make(map[int]int64),
		errors:  make(map[string]int64),
		latency: newStatistics(),
	}
}

func (r *Report) String() string {
	elapsed := time.Duration(r.totalTime)
	rps := float64(r.requests) / elapsed.Seconds()
	min := time.Duration(r.latency.min)
	max := time.Duration(r.latency.max)
	mean := time.Duration(r.latency.Mean())
	stdDev := time.Duration(r.latency.StdDev())

	var sb strings.Builder
	for code, count := range r.codes {
		sb.WriteString(fmt.Sprintf("        %d:             %d\n", code, count))
	}

	if len(r.errors) > 0 {
		sb.WriteString("    Errors:\n")
		for msg, count := range r.errors {
			sb.WriteString(fmt.Sprintf("        %s: %d\n", msg, count))
		}
	}

	return fmt.Sprintf(`
Summary:
    Elapsed:             %s
    Count:               %d
%s    Requests per Second: %.2f

Statistics:
    Min:  %s
    Max:  %s
    Mean: %s +/- %s
`,
		elapsed,
		r.requests,
		sb.String(),
		rps,
		min.String(),
		max.String(),
		mean.String(),
		stdDev.String(),
	)
}
