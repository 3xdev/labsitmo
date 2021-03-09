package main

import (
	"fmt"
	"strconv"
)

func main() {

	type Wall struct {
		N int
		K int
	}

	var numTests int
	var tests []Wall

	_, err := fmt.Scan(&numTests)
	if err != nil {
		fmt.Print("Error in number of tests: " + err.Error())
	}

	for i := 0; i < numTests; i++ {
		tests = append(tests, Wall{})

		_, err := fmt.Scan(&tests[i].N)
		if err != nil {
			fmt.Print("Error in number of fighters" + err.Error())
		}

		_, err = fmt.Scan(&tests[i].K)
		if err != nil {
			fmt.Print("Error in count of teams #" + strconv.Itoa(i) + err.Error())
		}
	}

	for _, f := range tests {

		var teams []int
		for i := 0; i < f.K; i++ {
			teams = append(teams, 0)
		}

		fighters := f.N
		for i := 0; i < len(teams); i++ {
			if fighters > 0 {
				teams[i]++
				fighters--
			} else {
				break
			}
			if i == len(teams)-1 && fighters > 0 {
				i = -1
			}
		}
		sum := 0
		for i := 0; i < len(teams); i++ {
			for j := i + 1; j < len(teams); j++ {
				sum += teams[i] * teams[j]
			}
		}
		fmt.Println(sum)

	}
}
