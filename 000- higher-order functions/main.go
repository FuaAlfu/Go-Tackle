package main

import (
	"errors"
	"fmt"
	"math"
	"net"
	"net/http"
	"strings"
	"time"
)

/*
TODO:
- Implement a higher-order function that retries a given operation.
- The function should take the following parameters:
	- fn: Operation to retry
*/

// Retry is a higher-order function that retries a given operation.
func Retry(
	fn func() error, // Operation to retry
	retryable func(error) bool, // Determines if an error is retryable
	retries int, // Max retry attempts
	backoff func(int) time.Duration, // Backoff strategy
) error {
	var err error
	for attempt := 0; attempt < retries; attempt++ {
		if err = fn(); err == nil {
			return nil // Success
		}
		if !retryable(err) {
			break // Non-retryable error
		}
		time.Sleep(backoff(attempt)) // Wait before retrying
	}
	return fmt.Errorf("failed after %d retries: %w", retries, err)
}

// Example: HTTP Request with Retry
func MakeRequest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err // Network error
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}
	return nil // Success
}

func main() {
	url := "https://api.example.com/data"

	err := Retry(
		func() error {
			return MakeRequest(url)
		},
		func(err error) bool {
			// Retry on network errors or 5xx server errors
			var netErr net.Error
			if errors.As(err, &netErr) {
				return true
			}
			return strings.Contains(err.Error(), "HTTP error: 5")
		},
		3, // Max retries
		func(attempt int) time.Duration {
			// Exponential backoff: 1s, 2s, 4s, etc.
			return time.Duration(math.Pow(2, float64(attempt))) * time.Second
		},
	)

	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Println("Request succeeded!")
	}
}
