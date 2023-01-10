package goitertools

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccumulateWithInitAdd(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go AccumulateWithInit(c, []int{1, 2, 3, 4, 5}, add, 0)

	want := []int{0, 1, 3, 6, 10, 15}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func BenchmarkAccumulateWithInitAdd(b *testing.B) {
	c := make(chan int)
	add := func(a, b int) int { return a + b }
	go AccumulateWithInit(c, []int{1, 2, 3, 4, 5}, add, 0)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestAccumulateWithInitAdd2(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go AccumulateWithInit(c, []int{1, 2, 3, 4, 5}, add, 10)

	want := []int{10, 11, 13, 16, 20, 25}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateWithInitMax(t *testing.T) {
	c := make(chan float64)

	go AccumulateWithInit(c, []float64{1, 2, 3, 4, 5}, math.Max, 0)

	want := []float64{0, 1, 2, 3, 4, 5}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateWithInitMin(t *testing.T) {
	c := make(chan float64)
	go AccumulateWithInit(c, []float64{1, 2, 3, 4, 5}, math.Min, 3)

	want := []float64{3, 1, 1, 1, 1, 1}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateWithInitOneItem(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go AccumulateWithInit(c, []int{1}, add, 0)

	want := []int{0, 1}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateWithInitEmpty(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go AccumulateWithInit(c, []int{}, add, 0)

	want := []int{0}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func ExampleAccumulateWithInit() {
	c := make(chan int)

	// Using a non-cummutative function, i.e. a-b != b-a
	sub := func(a, b int) int { return a - b }
	data := []int{5, 4, 3, 2, 1}
	// The first item sent from AccumulateWithInit will be `initial`
	go AccumulateWithInit(c, data, sub, 0)

	for v := range c {
		fmt.Printf("%v\n", v)
	}
	//Output: 0
	// -5
	// -9
	// -12
	// -14
	// -15

}

func TestAccumulateAdd(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go Accumulate(c, []int{1, 2, 3, 4, 5}, add)

	want := []int{1, 3, 6, 10, 15}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func BenchmarkAccumulateAdd(b *testing.B) {
	c := make(chan int)
	add := func(a, b int) int { return a + b }
	go Accumulate(c, []int{1, 2, 3, 4, 5}, add)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestAccumulateAdd2(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go Accumulate(c, []int{1, 2, 3, 4, 5}, add)

	want := []int{1, 3, 6, 10, 15}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateMax(t *testing.T) {
	c := make(chan float64)

	go Accumulate(c, []float64{1, 2, 3, 4, 5}, math.Max)

	want := []float64{1, 2, 3, 4, 5}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateMin(t *testing.T) {
	c := make(chan float64)
	go Accumulate(c, []float64{1, 2, 3, 4, 5}, math.Min)

	want := []float64{1, 1, 1, 1, 1}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func ExampleAccumulate() {
	c := make(chan int)

	// Using a non-cummutative function, i.e. a-b != b-a
	sub := func(a, b int) int { return a - b }
	data := []int{5, 4, 3, 2, 1}
	// The first item sent from AccumulateWithInit will be `initial`
	go Accumulate(c, data, sub)

	for v := range c {
		fmt.Printf("%v\n", v)
	}
	//Output: 5
	// 1
	// -2
	// -4
	// -5

}

func TestAccumulateOneItem(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go Accumulate(c, []int{1}, add)

	want := []int{1}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateEmpty(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go Accumulate(c, []int{}, add)

	collected := make([]int, 0)
	for v := range c {
		collected = append(collected, v)
	}
	assert.Empty(t, collected)
}

func FuzzAccumulate(f *testing.F) {
	// For this fuzz test, input should be just a random seed. We'll use rand to generate
	// how many items should go into `data`, and then fill it up
	f.Add(int64(42))
	f.Fuzz(func(t *testing.T, rand_seed int64) {
		// Set the seed so that we always pick the same n for the inputs of this function
		rand.Seed(rand_seed)
		// Don't both checking more than 100 numbers
		n := rand.Intn(100)
		// Make sure n isn't 0
		n += 1

		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = rand.Intn(1000)
		}

		// Get the python output. Obviously firing up python every iteration is not ideal,
		// but haven't yet figured out a way to keep python up and running between calls
		want := pythonAccumulate(data)

		// Get the go output
		got := make([]int, n)
		c := make(chan int)
		add := func(a, b int) int { return a + b }

		go Accumulate(c, data, add)

		for i := 0; i < n; i++ {
			got[i] = <-c
		}

		// Check that the go output matches python
		assert.Equal(t, want, got)
	})
}

func BenchmarkStrArray10Items(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		str_array(data)
	}
}

func TestStrArray(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := str_array(data)
	assert.Equal(t, "[1,2,3,4,5,6,7,8,9,10]", got)
}

// collect receives all the items from a channel until the channel is closed. NB, if the
// sender does not close the channel, THIS WILL HANG. All items go into a slice
func collect[T any](c chan T) []T {
	res := make([]T, 0)

	for v := range c {
		res = append(res, v)
	}
	return res
}

func TestChainFromIterable(t *testing.T) {
	testCases := []struct {
		desc string
		in   [][]int
		c    chan int
		want []int
	}{
		{
			desc: "simple",
			in:   [][]int{{1, 2}, {3, 4}},
			c:    make(chan int),
			want: []int{1, 2, 3, 4},
		},
		{
			desc: "with empties",
			in:   [][]int{{}, {5, 6}, {}},
			c:    make(chan int),
			want: []int{5, 6},
		},

		{
			desc: "all empty",
			in:   [][]int{{}, {}, {}},
			c:    make(chan int),
			want: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			go ChainFromIterable(tC.c, tC.in)
			got := collect(tC.c)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestPairwise(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		c    chan [2]int
		want [][2]int
	}{
		{
			desc: "simple",
			in:   []int{1, 2, 3, 4, 5},
			c:    make(chan [2]int),
			want: [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
		{
			desc: "empty",
			in:   []int{},
			c:    make(chan [2]int),
			want: [][2]int{},
		},
		{
			desc: "one item",
			in:   []int{1},
			c:    make(chan [2]int),
			want: [][2]int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			go Pairwise(tC.c, tC.in)
			got := collect(tC.c)
			assert.Equal(t, tC.want, got)
		})
	}
}

func FuzzPairwise(f *testing.F) {
	// For this fuzz test, input should be just a random seed. We'll use rand to generate
	// how many items should go into `data`, and then fill it up
	f.Add(int64(42))
	f.Fuzz(func(t *testing.T, rand_seed int64) {
		// Set the seed so that we always pick the same n for the inputs of this function
		rand.Seed(rand_seed)
		// Don't both checking more than 100 items
		n := rand.Intn(100)
		fmt.Println(n)
		// Make sure n isn't 0
		n += 1

		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = rand.Intn(1000)
		}

		// Get the python output. Obviously firing up python every iteration is not ideal,
		// but haven't yet figured out a way to keep python up and running between calls
		want := pythonPairwise(data)

		// Get the go output
		got := make([][2]int, 0, n)
		c := make(chan [2]int)

		go Pairwise(c, data)

		for v := range c {
			got = append(got, v)
		}

		// Check that the go output matches python
		assert.Equal(t, want, got)
	})
}
