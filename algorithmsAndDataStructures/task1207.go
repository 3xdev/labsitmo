package main

import (
	"fmt"
)

type Dot struct {
	x int
	y int
}

type Plot struct {
	k          float64
	b          int
	isVertical bool
}

func main() {

	var N int
	var dots []Dot

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var tx, ty int
		fmt.Scan(&tx, &ty)
		dots = append(dots, Dot{
			x: tx,
			y: ty,
		})
	}

	type EastWest struct {
		east int
		west int
	}

	var c EastWest

	for i := 0; i < len(dots); i++ {
		for j := i + 1; j < len(dots); j++ {
			plot := buildPlot(dots[i], dots[j])
			if plot.isVertical {	
				fmt.Printf("\nPlot is vertical.\ndots: %d %d\nx == %d", i+1, j+1, plot.b)
			} else {
				fmt.Printf("\ndots: %d %d\nkx+b == %0.2fx + %d\n", i+1, j+1, plot.k, plot.b)
			}

			for k := 0; k < len(dots); k++ {
				if !(k == i || k == j) {
					result := calculateResult(plot, dots[k])
					if result > 0 {
						c.east++
					} else {
						c.west++
					}
				}
			}

			fmt.Println("Dots on East: ", c.east, " // Dots on West: ", c.west)
			if c.east == c.west {
				fmt.Println(i+1, j+1)
				break
			}
			c = EastWest{}
		}
		break

	}

}

func calculateResult(plot Plot, dot Dot) float64 {
	if plot.isVertical {
		if dot.x < plot.b {
			return 1
		} else {
			return -1
		}
	} else {
		return plot.k*float64(dot.x) + float64(plot.b)
	}
}

func buildPlot(dot1 Dot, dot2 Dot) Plot {

	var tk float64
	if dot1.x == dot2.x {
		return Plot{
			k:          0,
			b:          dot1.x,
			isVertical: true,
		}
	} else {
		tk = float64((dot1.y - dot2.y) / (dot1.x - dot2.x))
	}

	tb := float64(dot2.y) - tk*float64(dot2.x)
	return Plot{
		k:          tk,
		b:          int(tb),
		isVertical: false,
	}
}
