package stresstest

import (
	"time"
)

type Reporter struct {
	start  time.Time
	report Report
}

func NewReporter() *Reporter {
	return &Reporter{
		start:  time.Now(),
		report: newReport(),
	}
}

func (r *Reporter) Collect(results <-chan Result) Report {
	for {
		res, ok := <-results
		if !ok {
			break
		}

		r.report.requests += 1
		if res.Code != 0 {
			r.report.codes[res.Code] += 1
		}
		if res.Error != "" {
			r.report.errors[res.Error] += 1
		}
		r.report.latency.Add(float64(res.Duration))
	}

	r.report.totalTime = float64(time.Since(r.start))

	return r.report
}
