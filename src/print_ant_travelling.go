package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintHantTravelling(ways []string, endNode string) {
	// we create a map who gonna take the number of an ant and his way
	ant := make(map[int]string)

	// we take the len of the longest and the shortest way in the array of strings
	max := SearchStr(ways[0], " -> ")
	min := max
	minWay := ways[0]

	for i := range ways {
		if SearchStr(ways[i], " -> ") > max {
			max = SearchStr(ways[i], " -> ")
		} else if SearchStr(ways[i], " -> ") < min {
			min = SearchStr(ways[i], " -> ")
			minWay = ways[i]
		}
	}

	max--

	for _, i := range ways {
		fmt.Println(i)
	}
	fmt.Println()

	/*
	  we initialize 1 counter and 1 index to 0 and increment them so that we
	  can assign the correct path to each ant so that the path is made in the most optimized way possible
	*/
	index := 0
	count := 0
	for i := 1; i <= AntNbr; i++ {
		if i == AntNbr {
			ant[i] = minWay
			break
		}

		ant[i] = ways[index]
		count += SearchStr(ways[index], " -> ")

		if count >= max-min {
			if index == len(ways)-1 {
				index = 0
			} else {
				index++
			}
			count = 0
		}
	}

	// we loop while all the ant aren't arrive to the end node
	for AntEndTravel := 0; AntEndTravel < AntNbr; {
		var tab []string
		// we loop for all ant
		for i := 1; i <= AntNbr; i++ {
			antCurrentWay := strings.Split(ant[i], " -> ")
			if len(antCurrentWay) != 1 && !IsinTab3(antCurrentWay[0]+" -> "+antCurrentWay[1], tab) {
				tab = append(tab, antCurrentWay[0]+" -> "+antCurrentWay[1])

				fmt.Print("L" + strconv.Itoa(i) + "-" + antCurrentWay[1] + " ")
				ant[i] = strings.Join(antCurrentWay[1:], " -> ")
				if SearchStr(ant[i], " -> ") == 0 {
					AntEndTravel++
				}
			}
		}
		fmt.Println()
	}
}
