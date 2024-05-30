package itertools_test

import (
	"testing"

	"github.com/Warashi/go-itertools"
)

func TestCount(t *testing.T) {
	var i int
	for v := range itertools.Count(3, 5) {
		if v != 3+5*i {
			t.Errorf("Count(3, 5) = %d; want %d", v, 3+5*i)
		}
		if i >= 10 {
			break
		}
		i++
	}
}
