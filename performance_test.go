package main

import (
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("")
	}
	limit := 50000
	start := time.Now()
	naive(limit)
	t.Logf("[naive]        limit %d took %s", limit, time.Since(start))
	start = time.Now()
	sieve(limit)
	t.Logf("[sieve]        limit %d took %s", limit, time.Since(start))
	start = time.Now()
	parallelSieve(limit)
	t.Logf("[paralleSieve] limit %d took %s", limit, time.Since(start))
	start = time.Now()
	memoize(limit)
	t.Logf("[memoize]      limit %d took %s", limit, time.Since(start))
}
