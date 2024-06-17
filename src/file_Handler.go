package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// HandleFile is a function that reads a file and returns its contents as a string.
//
// It expects the name of the file as a command-line argument and returns an error
// if the file cannot be read or if the command-line argument is missing.
//
// Returns:
// - string: the contents of the file.
// - error: an error if the file cannot be read or if the command-line argument is missing.
func HandleFile() (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("no fileName in option")
	}
	file, err := os.ReadFile(os.Args[1])
	return string(file), err
}

func parseGraph(text string) (map[string][]Edge, string, string) {
	graph := make(map[string][]Edge)
	var startNode, endNode string

	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "##start") {
			// Save the name of the node after "##start"
			scanner.Scan()
			startNode = strings.Fields(scanner.Text())[0]
		} else if strings.HasPrefix(line, "##end") {
			// Save the name of the node after "##end"
			scanner.Scan()
			endNode = strings.Fields(scanner.Text())[0]
		} else if strings.Contains(line, "-") {
			// It's an edge definition
			edge := strings.Split(line, "-")
			start := edge[0]
			end := edge[1]

			if start == end {
				log.Fatalln("ERROR TO PUT")
			}
			graph[start] = append(graph[start], Edge{End: end})
			graph[end] = append(graph[end], Edge{End:start})
		}
	}

	if startNode == "" || endNode == "" {
		log.Fatalln("ERROR TO PUT")
	}

	return graph, startNode, endNode
}

func printGraph(graph map[string][]Edge, startNode, endNode string) {
	fmt.Println("Graph:")
	for node, edges := range graph {
		fmt.Printf("%s -> %v\n", node, edges)
	}

	fmt.Printf("\nStart Node: %s\n", startNode)
	fmt.Printf("End Node: %s\n", endNode)
}
