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
