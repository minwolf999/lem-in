package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Edge struct {
	End string
}

var AntNbr = 0

func main() {
	// we get the data of the text file
	graphText, err := HandleFile()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(graphText, "\n\n")

	// we take the number of ant
	temp := strings.Fields(graphText)
	AntNbr, err = strconv.Atoi(temp[0])
	if err != nil {
		log.Fatalln("ERROR TO PUT")
	}

	if AntNbr <= 0 {
		log.Fatalln("ERROR TO PUT")
	}

	// we change the data of the text file in a graph and we take the start/end node off our graph
	graph, startNode, endNode := parseGraph(graphText)

	// we search if there is an error in the text file
	Verification(graph, endNode)

	/* printGraph(graph, startNode, endNode)

	fmt.Println("Quantity of Ant:", AntNbr)

	fmt.Println() */

	// we look for all possible paths to go from the start node to the end node
	var ways [][]string
	GetWays(graph, startNode, endNode, make(map[string]bool), []string{}, &ways)

	if len(ways) == 0 {
		log.Fatalln("ERROR TO PUT")
	}
	// we sort the paths obtained to keep only the paths which will be useful for print
	ways = Tri(ways)
	ways = TriWays(ways, graph, startNode, endNode, ways[0])

	// we change our array of array of string in an array of string
	var fways []string
	for _, i := range ways {
		fways = append(fways, strings.Join(i, " -> "))
	}

	// we print the travel of the ant
	PrintHantTravelling(fways, endNode)
}
