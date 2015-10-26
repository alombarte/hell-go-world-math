package hellomath

import "testing"

func TestFibonacci(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
	}
	fib := Fibonacci()
	for _, c := range cases {
		got := fib()
		if got != c.want {
			t.Errorf("Fibonacci(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
