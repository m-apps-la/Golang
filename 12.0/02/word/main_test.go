package word

import (
	"fmt"
	"testing"

	"github.com/Golang/12.0/02/quote"
)

func TestUseCount(t *testing.T) {
	x := UseCount(`one two three three three`)
	for k, v := range x {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("Got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("Got", v, "want", 1)
			}
		}
	}
}
func TestCount(t *testing.T) {
	x := Count(`one two three`)
	if x != 3 {
		t.Error("Got", x, "want", 3)
	}
}

func ExampleCount() {
	fmt.Println(Count(`one two three`))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
