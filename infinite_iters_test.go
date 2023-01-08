package goitertools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount1(t *testing.T) {
	assert := assert.New(t)

	c := make(chan int)
	go Count(0, 1, c)

	assert.Equal(0, <-c)
	assert.Equal(1, <-c)
	assert.Equal(2, <-c)
	assert.Equal(3, <-c)
	assert.Equal(4, <-c)
}

func BenchmarkCount1(b *testing.B) {
	c := make(chan int)
	go Count(0, 1, c)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestCount2(t *testing.T) {
	assert := assert.New(t)

	c := make(chan int)
	go Count(0, 4, c)

	assert.Equal(0, <-c)
	assert.Equal(4, <-c)
	assert.Equal(8, <-c)
	assert.Equal(12, <-c)
	assert.Equal(16, <-c)
}

func BenchmarkCount2(b *testing.B) {
	c := make(chan int)
	go Count(0, 4, c)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestCycle1(t *testing.T) {
	assert := assert.New(t)

	c := make(chan float64)
	data := []float64{1.0, 2.0, 3.14, 5.0}
	go Cycle(data, c)

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
	go Cycle(data, c)

	for i := 0; i < b.N; i++ {
		<-c
	}
}

func TestCycle2(t *testing.T) {
	assert := assert.New(t)

	c := make(chan []int)
	data := [][]int{{1, 2}, {3, 4}}
	go Cycle(data, c)

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
	go Cycle(data, c)

	for i := 0; i < b.N; i++ {
		<-c
	}
}
