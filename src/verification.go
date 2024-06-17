package main

import (
	"fmt"
	"log"
	"strings"
)

/*
function which checks if the contents of the file are correct (if there is not an edge which connects to a node which does not exist,
or if two nodes are identical)
*/
func Verification(graph map[string][]Edge, endNode string) {
	var tabNode []string
	for i := range graph {
		tabNode = append(tabNode, i)
	}
	tabNode = append(tabNode, endNode)

	for _, edges := range graph {
		for _, value := range edges {
			edge := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%s", value), "{", ""), "}", "")
			if !IsinTab3(edge, tabNode) {
				log.Fatalln("ERROR TO PUT")
			}
		}
	}
}
