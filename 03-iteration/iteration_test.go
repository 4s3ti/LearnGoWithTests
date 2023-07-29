package main

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func TestRepeatStrings(t *testing.T) {
	repeated := RepeatStrings("b", 10)
	expected := "bbbbbbbbbb"
	
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}



func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// output: aaaaa
}
