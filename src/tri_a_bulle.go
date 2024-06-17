package main

func Tri(tab [][]string) [][]string {
	for i := len(tab)-1; i > 0; i-- {
		for y := 0; y < i; y++ {
			if len(tab[y]) > len(tab[y+1]) {
				tab[y], tab[y+1] = tab[y+1], tab[y]
			}
		}
	}
	return tab
}