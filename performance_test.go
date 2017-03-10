package main

import (
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("")
	}
	limit := 5000
	start := time.Now()
	naive(limit)
	t.Logf("[naive]        limit %6d took %s", limit, time.Since(start))
	limit = 500000
	start = time.Now()
	memoize(limit)
	t.Logf("[memoize]      limit %6d took %s", limit, time.Since(start))
	start = time.Now()
	sieve(limit)
	t.Logf("[sieve]        limit %6d took %s", limit, time.Since(start))
}
