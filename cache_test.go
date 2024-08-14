package itertools_test

import (
	"fmt"
	
	"github.com/Warashi/go-itertools"
)

func ExampleCache() {
	iter := itertools.Cache(func (yield func(int) bool) {
		for i := range 3 {
			fmt.Println("yield", i)
			yield(i)
		}
	})

	for i := range iter {
		fmt.Println("receive" , i)
	}
	for i := range iter {
		fmt.Println("receive" , i)
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
