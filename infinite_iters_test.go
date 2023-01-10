package goitertools

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount1(t *testing.T) {
	assert := assert.New(t)

	c := make(chan int)
	go Count(c, 0, 1)

	assert.Equal(0, <-c)
	assert.Equal(1, <-c)
	assert.Equal(2, <-c)
	assert.Equal(3, <-c)
	assert.Equal(4, <-c)
}

func BenchmarkCount1(b *testing.B) {
	c := make(chan int)
	go Count(c, 0, 1)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestCount2(t *testing.T) {
	assert := assert.New(t)

	c := make(chan int)
	go Count(c, 0, 4)

	assert.Equal(0, <-c)
	assert.Equal(4, <-c)
	assert.Equal(8, <-c)
	assert.Equal(12, <-c)
	assert.Equal(16, <-c)
}

func BenchmarkCount2(b *testing.B) {
	c := make(chan int)
	go Count(c, 0, 4)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestCycle1(t *testing.T) {
	assert := assert.New(t)

	c := make(chan float64)
	data := []float64{1.0, 2.0, 3.14, 5.0}
	go Cycle(c, data)

	assert.Equal(1.0, <-c)
	assert.Equal(2.0, <-c)
	assert.Equal(3.14, <-c)
	assert.Equal(5.0, <-c)
	assert.Equal(1.0, <-c)
	assert.Equal(2.0, <-c)
	assert.Equal(3.14, <-c)
	assert.Equal(5.0, <-c)
}

func BenchmarkCycle1(b *testing.B) {
	c := make(chan float64)
	data := []float64{1.0, 2.0, 3.14, 5.0}
	go Cycle(c, data)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestCycle2(t *testing.T) {
	assert := assert.New(t)

	c := make(chan []int)
	data := [][]int{{1, 2}, {3, 4}}
	go Cycle(c, data)

	assert.Equal([]int{1, 2}, <-c)
	assert.Equal([]int{3, 4}, <-c)
	assert.Equal([]int{1, 2}, <-c)
	assert.Equal([]int{3, 4}, <-c)
	assert.Equal([]int{1, 2}, <-c)
	assert.Equal([]int{3, 4}, <-c)
}

func BenchmarkCycle2(b *testing.B) {
	c := make(chan []int)
	data := [][]int{{1, 2}, {3, 4}}
	go Cycle(c, data)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestPyCount(t *testing.T) {
	n := 3
	c := make(chan int)
	go Count(c, 0, 1)

	got := make([]int, n)
	got[0] = <-c
	got[1] = <-c
	got[2] = <-c

	want := pythonCount(0, 1, n)

	assert.Equal(t, want, got)

}

func FuzzCount(f *testing.F) {
	// For this fuzz test, inputs need to be start, step, and a random seed
	f.Add(0, 1, int64(42))
	f.Fuzz(func(t *testing.T, start int, step int, rand_seed int64) {
		// Set the seed so that we always pick the same n for the inputs of this function
		rand.Seed(rand_seed)
		// Don't both checking more than 100 numbers
		n := rand.Intn(100)
		// Make sure n isn't 0
		n += 1

		// Get the python output. Obviously firing up python every iteration is not ideal,
		// but haven't yet figured out a way to keep python up and running between calls
		want := pythonCount(start, step, n)

		// Get the go output
		got := make([]int, n)
		c := make(chan int)
		go Count(c, start, step)
		for i := 0; i < n; i++ {
			got[i] = <-c
		}

		// Check that the go output matches python
		assert.Equal(t, want, got)
	})
}
