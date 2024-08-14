package itertools_test

import (
	"fmt"
	"iter"

	"github.com/Warashi/go-itertools"
)

func ExampleCache() {
	iter := itertools.Cache(func(yield func(int) bool) {
		for i := range 3 {
			fmt.Println("yield", i)
			yield(i)
		}
	})

	for i := range iter {
		fmt.Println("receive", i)
	}
	for i := range iter {
		fmt.Println("receive", i)
	}

	// Output:
	// yield 0
	// receive 0
	// yield 1
	// receive 1
	// yield 2
	// receive 2
	// receive 0
	// receive 1
	// receive 2
}

func Example() {
	var (
		isPrime func(int) bool
		primes  iter.Seq[int]
	)

	isPrime = func(n int) bool {
		allNonDivisable := itertools.ForAll(func(x int) bool {
			return n%x != 0
		})
		takeWhile := itertools.TakeWhile(func(x int) bool {
			return x*x <= n
		})
		return allNonDivisable(takeWhile(primes))
	}

	primes = itertools.Cache(func(yield func(int) bool) {
		if !yield(2) {
			return
		}
		for i := 3; ; i += 2 {
			if isPrime(i) {
				if !yield(i) {
					return
				}
			}
		}
	})

	for p := range itertools.Take[int](5)(primes) {
		fmt.Println(p)
	}

	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
}
