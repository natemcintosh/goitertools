package goitertools

// Accumulate sends accumulated results. E.g. if `fn` adds the two inputs, then the
// channel sends the cumulative sum:
// `Accumulate([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }, 0, c) --> 1, 3, 6, 10`
//
// It is not recommended to use data where T is a pointer or pointer-like (e.g. slice)
// As items in `data`, or `initial` may be changed. It also is not threadsafe
//
// Note that that the order of arguments passed into `fn` will always be the accumulated
// value first, and then the next item from data. E.g. for subtraction with
// `data := []int{3, 2, 1}`, Accumulate sends 3, 1, 0. If `data := []int{1, 2, 3}`,
// Accumulate sends 1, -1, -4
func Accumulate[T any](data []T, fn func(T, T) T, initial T, c chan T) {
	res := initial
	for _, d := range data {
		res = fn(res, d)
		c <- res
	}
	close(c)
}
