package stresstest

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Tester struct {
	url         string
	requests    int
	concurrency int
	resultChan  chan Result
	client      http.Client
	closeOnce   sync.Once
}

func NewTester(url string, requests, concurrency int) *Tester {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Tester{
		url:         url,
		requests:    requests,
		concurrency: concurrency,
		client:      http.Client{Transport: tr},
		resultChan:  make(chan Result, concurrency),
	}
}

func (t *Tester) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	remaining := int32(t.requests)
	var wg sync.WaitGroup
	for range t.concurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				// Check if the context have not been closed
				select {
				case <-ctx.Done():
					return
				default:
				}

				if atomic.AddInt32(&remaining, -1) < 0 {
					return
				}

				result := t.doRequest()
				t.resultChan <- *result
			}
		}()
	}

	wg.Wait()
	t.closeResultChannel()
}

func (t *Tester) Results() <-chan Result {
	return t.resultChan
}

func (t *Tester) doRequest() *Result {
	start := time.Now()

	r, err := t.client.Get(t.url)
	if err != nil {
		res := Result{
			Duration: time.Since(start),
			Error:    err.Error(),
		}

		if r != nil {
			res.Code = r.StatusCode
		}

		return &res
	}

	return &Result{
		Code:     r.StatusCode,
		Duration: time.Since(start),
	}
}

func (t *Tester) closeResultChannel() {
	t.closeOnce.Do(func() {
		close(t.resultChan)
	})
}
