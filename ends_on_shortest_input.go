package goitertools

// Accumulate sends accumulated results. E.g. if `fn` adds the two inputs, then
// the channel sends the cumulative sum:
// `Accumulate([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }, c) --> 1, 3, 6, 10`
//
// It is not recommended to use data where T is a pointer or pointer-like (e.g. slice)
// As items in `data`, or `initial` may be changed. It also is not threadsafe
//
// Note that that the order of arguments passed into `fn` will always be the accumulated
// value first, and then the next item from data. E.g. for subtraction with
// `data := []int{3, 2, 1}`, Accumulate sends 3, 1, 0. If `data := []int{1, 2, 3}`,
// Accumulate sends 1, -1, -4
//
// Also note that the first sent by `c` will be the first item in data. See
// `AccumulateWithInit` for a version of this function that starts with `initial`
//
// If `data` is empty, the channel is immediately closed
func Accumulate[S ~[]T, T any](c chan T, data S, fn func(T, T) T) {
	// If there's no data, close the channel and return
	if len(data) == 0 {
		close(c)
		return
	}

	// Send the first item
	res := data[0]
	c <- res

	for _, d := range data[1:] {
		res = fn(res, d)
		c <- res
	}
	close(c)
}

// AccumulateWithInit behaves in the same manner has Accumulate, but the first item
// sent by `c` is `initial`
//
// If `data` is empty, this channel will send `initial` and then close.
// Note the difference from `Accumulate` which immediately closes the channel having
// sent nothing
func AccumulateWithInit[S ~[]T, T any](c chan T, data S, fn func(T, T) T, initial T) {
	res := initial
	c <- res

	for _, d := range data {
		res = fn(res, d)
		c <- res
	}
	close(c)
}

func Chain[S ~[]T, T any](c chan T, iterables ...S) {

}
