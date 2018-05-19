package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(0)
	var withinCircle int
	var withinSquare int
	var diff int
	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Printf(
				"\rApproximate Pi %f total points: %d rate: %d/s",
				(float64(withinCircle)/float64(withinSquare))*4,
				withinSquare,
				withinSquare-diff,
			)
			diff = withinSquare
		}
	}()
	for {
		randX := rand.Float64()
		randY := rand.Float64()
		if isInside(randX, randY) {
			withinCircle += 1
		}
		withinSquare += 1
	}
}

func isInside(x float64, y float64) bool {
	// Compare radius of circle with distance
	// of its center from given point
	if (x-0.5)*(x-0.5)+
		(y-0.5)*(y-0.5) < 0.5*0.5 {
		return true
	}
	return false
}
