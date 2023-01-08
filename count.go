package goitertools

// Count will place numbers in the channel `c` to be received
// e.g. Count(10, 1, c) --> 10, 11, 12, 13, 14, ...
func Count(start int, step int, c chan int) {
	step_mul := 0
	for {
		c <- start + (step * step_mul)
		step_mul += 1
	}
}
