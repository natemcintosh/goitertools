package goitertools

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccumulateWithInitAdd(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go AccumulateWithInit([]int{1, 2, 3, 4, 5}, add, 0, c)

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
	go AccumulateWithInit([]int{1, 2, 3, 4, 5}, add, 0, c)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestAccumulateWithInitAdd2(t *testing.T) {
	c := make(chan int)

	add := func(a, b int) int { return a + b }
	go AccumulateWithInit([]int{1, 2, 3, 4, 5}, add, 10, c)

	want := []int{10, 11, 13, 16, 20, 25}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateWithInitMax(t *testing.T) {
	c := make(chan float64)

	go AccumulateWithInit([]float64{1, 2, 3, 4, 5}, math.Max, 0, c)

	want := []float64{0, 1, 2, 3, 4, 5}

	idx := 0
	for v := range c {
		assert.Equal(t, want[idx], v)
		idx += 1
	}
}

func TestAccumulateWithInitMin(t *testing.T) {
	c := make(chan float64)
	go AccumulateWithInit([]float64{1, 2, 3, 4, 5}, math.Min, 3, c)

	want := []float64{3, 1, 1, 1, 1, 1}

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
	go AccumulateWithInit(data, sub, 0, c)

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
