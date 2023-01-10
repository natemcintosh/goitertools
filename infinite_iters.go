package goitertools

// Count will place numbers in the channel `c` to be received
// e.g. `go Count(c, 10, 1)` --> 10, 11, 12, 13, 14, ...
func Count(c chan int, start int, step int) {
	step_mul := 0
	for {
		c <- start + (step * step_mul)
		step_mul += 1
	}
}

// Cycle returns elements from `data`. When data is exhausted, go back to the start
// e.g. `go Cycle(c, []int{1, 2, 3, 4})` --> 1, 2, 3, 4, 1, 2, ...
//
// If you want to iterate over a string, you can convert it to a slice of runes with
// `[]rune(my_string)`
//
// Note that if T is a pointer (e.g. *MyType) or pointer-like (e.g. a slice or map), the
// channel will send the pointer that value. This means they are not threadsafe, and you
// should be aware that the items in `data` can change
func Cycle[S ~[]T, T any](c chan T, data S) {
	for {
		for _, item := range data {
			c <- item
		}
	}
}
