package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {

	t.Run("Collection (slice) of any size", func(t * testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want { 
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
	

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TesSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum emptu slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0,9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
