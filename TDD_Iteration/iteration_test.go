package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	output := "aaaaa"
	recieved := Repeat("a")
	expected := output

	if recieved != expected {
		t.Errorf("expected %q but got %q", expected, recieved)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// func ExampleRepeat() {
// 	output := Repeat("a")
// 	fmt.Println(output)
// 	// Output: "aaaaa"
// }
