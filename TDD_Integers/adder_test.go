package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	recieved := Add(2, 2)
	expected := 4

	if recieved != expected {
		t.Errorf("expected '%d' but got '%d'", expected, recieved)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
