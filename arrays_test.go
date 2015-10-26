package hellomath

import "testing"

func TestArraySum(t *testing.T) {

	cases := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{50, 50}, 100},
		{[]int{-10, 10}, 0},
		{[]int{1}, 1},
		{[]int{0, 0, 0, 0, 0, 0, 1}, 1},
	}

	for _, c := range cases {
		got := SumArray(c.in)
		if got != c.want {
			t.Errorf("SumArray(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
