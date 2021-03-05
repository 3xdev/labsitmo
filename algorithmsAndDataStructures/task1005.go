package main

import (
	"fmt"
	"math"
)

type bin struct {
	slice []int
}

func (b *bin) plusOne() {
	for i := len(b.slice) - 1; i >= 0; i-- {
		if b.slice[i] == 0 {
			b.slice[i] = 1
			for j := i + 1; j < len(b.slice); j++ {
				b.slice[j] = 0
			}
			break
		}
	}
}

func main() {

	var n int
	var w []int
	var a []int
	var difference = 100000

	var bs bin

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		w = append(w, 0)
		a = append(a, 0)
		fmt.Scan(&w[i])
	}
	bs.slice = a

	for i := 0; i < 1<<len(bs.slice); i++ {
		bs.plusOne()
		w1, w2 := 0, 0
		for j := 0; j < len(bs.slice); j++ {
			if bs.slice[j] == 1 {
				w1 += w[j]
			} else {
				w2 += w[j]
			}

		}
		difference = int(math.Min(float64(difference), math.Abs(float64(w1-w2))))
	}
	fmt.Println(difference)
}
