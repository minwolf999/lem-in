package main

func SearchStr(input string, tofind string) int {
	var output int = 0
	for i := range input {
		if i+len(tofind)-1 > len(input)-1 {
			break
		}
		if input[i:i+len(tofind)] == tofind {
			output++
		}
	}
	return output
}
