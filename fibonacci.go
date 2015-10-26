// Sample package containing dummy math functions
package hellomath

// fibonacci function returns an anonymous function that returns an integer itself.
func Fibonacci() func() int {
	// fmt.Println("This code is executed only the first time")
	// Notice that this vars will be modified by the following anonymous function every time.
	prev, next := 0, 1

	return func() int {
		// fmt.Println("This code is executed subsequent times")
		prev, next = next, prev+next
		return prev
	}

}
