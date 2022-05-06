package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	t.Run("2+2", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("Got '%d', wanted '%d'", got, want)
		}
	})
	t.Run("0+0", func(t *testing.T) {
		got := Add(0, 0)
		want := 0

		if got != want {
			t.Errorf("Got '%d', wanted '%d'", got, want)
		}
	})
}
