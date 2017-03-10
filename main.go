package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

func naive(limit int) []int {
	primes := []int{2}
	for i := 3; i <= limit; i++ {
		n := float64(i)
		found := false
		for i := 2.0; i < n; i++ {
			if math.Mod(n, i) == 0 {
				found = true
				break
			}
		}
		if !found {
			primes = append(primes, i)
		}
	}
	return primes
}

func sieve(limit int) []int {
	thres := math.Sqrt(float64(limit))
	sieb := map[int]bool{0: true, 1: true}
	for i := 2; i<limit;i++{
		sieb[i] = false
	}
	for i := 0; float64(i) <= thres; i++ {
		if !sieb[i] {
			for j:=i*i; j <= limit; j+=i {
				if !sieb[j] {
					sieb[j] = true
				}
			}
		}
	}
	primes := []int{}
	for k, v := range sieb {
		if !v {
			primes = append(primes, k)
		}
	}
	return primes
}

func memoize(limit int) []int {
	primes := []int{2}
	for i := 3; i <= limit; i++ {
		n := float64(i)
		found := false
		thres := math.Sqrt(n)
		for _, p := range primes {
			if float64(p) > thres {
				break
			}
			if math.Mod(n, float64(p)) == 0 {
				found = true
				break
			}
		}
		if !found {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var limit int
	var verbose bool
	var algorithm string
	flag.IntVar(&limit, "limit", 0, "Last number to check.")
	flag.BoolVar(&verbose, "verbose", false, "Wether or not to print the primes.")
	flag.StringVar(&algorithm, "algorithm", "naive", "Which algorithm to use.")
	flag.Parse()
	primes := []int{}
	switch algorithm {
	case "naive":
		primes = naive(limit)
	case "sieve":
		primes = sieve(limit)
	case "memoize":
		primes = memoize(limit)
	}
	if verbose {
		sort.Ints(primes)
		for _, p := range primes {
			fmt.Println(p)
		}
	}
}
