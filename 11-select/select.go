package main

import (
	"time"
	"net/http"
	"fmt"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winnter string, error error) {
	return ConfigurableRacer(a, b, 10 * tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winnter string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
