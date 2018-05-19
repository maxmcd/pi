package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	rand.Seed(0)
	for n := 0; n < b.N; n++ {
		randX := rand.Float64()
		randY := rand.Float64()
		isInside(randX, randY)
	}
}
