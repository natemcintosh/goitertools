package goitertools

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccumulateAdd(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go Accumulate([]int{1, 2, 3, 4, 5}, add, 0, c)

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
	go Accumulate([]int{1, 2, 3, 4, 5}, add, 0, c)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestAccumulateAdd2(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go Accumulate([]int{1, 2, 3, 4, 5}, add, 10, c)

	want := []int{11, 13, 16, 20, 25}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateMax(t *testing.T) {
	c := make(chan float64)

	go Accumulate([]float64{1, 2, 3, 4, 5}, math.Max, 0, c)

	want := []float64{1, 2, 3, 4, 5}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateMin(t *testing.T) {
	c := make(chan float64)
	go Accumulate([]float64{1, 2, 3, 4, 5}, math.Min, 1, c)

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
	// The first item sent from Accumulate will be `initial - data[0]`, in this case
	// 0 - 5 = -5
	go Accumulate(data, sub, 0, c)

	for v := range c {
		fmt.Printf("%v\n", v)
	}
	//Output: -5
	// -9
	// -12
	// -14
	// -15

}
