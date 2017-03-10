package main

import (
	"testing"
	"sort"
)

func TestPrimes(t *testing.T) {
	limit := 100
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	results := map[string][]int{"naive": naive(limit), "memoize": memoize(limit), "sieve": sieve(limit)}
	for algorithm, primes := range results {
		if len(primes) != len(expected) {
			t.Errorf("[%s] Number of primes don't match %d vs %d", algorithm, len(primes), len(expected))
		}
		sort.Ints(primes)
		for i := range primes {
			if primes[i] != expected[i] {
				t.Errorf("[%s] Primes don't match at %d: %d vs %d", algorithm, i, primes[i], expected[i])
			}
		}
	}
}
