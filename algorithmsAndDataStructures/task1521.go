package main

import "fmt"

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	var h []int

	for i := 0; i < n; i++ {
		h = append(h, i+1)
	}

	for i, pos := 0, 1; i < n; i++ {
		switch {
		case pos%3 == 0:
			fmt.Print(h[i], " ")
			h[i] = 0
			pos++
		case h[i] == 0:
			continue
		case i == n-1:
			i = -1
			pos++
		default:
			pos++
		}

	}

}
