package main

import (
	"fmt"
	"strings"
)

func GetWays(graph map[string][]Edge, node string, end string, visited map[string]bool, path []string, paths *[][]string) {
	// if we have already passed into the current node we go back to search an other path
	if visited[node] {
		return
	}

	// we write the current node as already passed and we add the current node in the path array of string
	visited[node] = true
	path = append(path, node)

	// if the current node is the last node we reset the path array and write the path found in the array of array of string
	if node == end {
		// Si nous avons atteint la destination, ajoutons le chemin actuel aux chemins trouvés
		newPath := make([]string, len(path))
		copy(newPath, path)
		*paths = append(*paths, newPath)

	// else we make a recursive with the edges off the current node
	} else {
		// Sinon, poursuivons la recherche en profondeur
		neighbors := graph[node]
		for _, i := range neighbors {
			neighbor := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%s", i), "{", ""), "}", "")
			GetWays(graph, neighbor, end, visited, path, paths)
		}
	}

	// we reset the map off visited node and we remove the last node in the path array
	visited[node] = false
	path = path[:len(path)-1]
}

// this function sorts the paths obtained to return only the useful paths
func TriWays(ways [][]string, graph map[string][]Edge, startNode string, endNode string, connard []string) [][]string {
	var tri [][]string

	// if the paths obtained is already good we return ways
	if (len(ways) == len(graph[startNode]) || len(ways) == len(graph[ways[0][len(ways[0])-1]])) && !IsinTab1(ways) {
		return ways
	}

	// if we can only have 1 path we return the shortest path
	if len(graph[endNode]) == 1 || len(graph[startNode]) == 1 {
		min := ways[0]
		for _, i := range ways {
			if len(i) < len(min) {
				min = i
			}
		}
		return [][]string{min}
	}

	/* we append the ways if it's possible in the new [][]string and we return it if the number of way in the new 
	  [][]string is equal to the number of way connected to the end node
	*/
	index := 0
	for len(tri) != len(graph[startNode]) && len(tri) != len(graph[endNode]) {
		tri = [][]string{}
		ways[0], ways[index] = ways[index], ways[0]
		for i := range ways {
			if !IsinTab2(ways[i], tri) {
				tri = append(tri, ways[i])
			}
		}
		index++
	}

	return tri
}
