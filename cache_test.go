package itertools_test

import (
	"fmt"
	
	"github.com/Warashi/go-itertools"
)

func ExampleCache() {
	iter := itertools.Cache(func (yield func(int) bool) {
		for i := range 3 {
			fmt.Println(i)
			yield(i)
		}
	})

	for _ = range iter {
	}
	for _ = range iter {
	}

	// Output:
	// 0
	// 1
	// 2
}
