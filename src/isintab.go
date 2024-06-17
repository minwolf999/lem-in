package main

// this function check if the [][]string have 2 time the same value (else of the first and last)
func IsinTab1(tab [][]string) bool {
	for x1, i1 := range tab {
		for z1, y1 := range i1 {
			if z1 == 0 || z1 == len(i1)-1 {
				continue
			}
			for x2, i2 := range tab {
				if x1 == x2 {
					continue
				}

				for _, y2 := range i2 {
					if y1 == y2 {
						return true
					}
				}
			}
		}
	}
	return false
}

// this function check if on of the values of the []string is already in the [][]string
func IsinTab2(tab []string, tab2 [][]string) bool {
	for i1, i := range tab {
		for _, y := range tab2 {
			for _, x := range y {
				if i == x && i1 != 0 && i1 != len(tab)-1 {
					return true
				}
			}
		}
	}
	return false
}

// Function that checks if a string is present in an array of strings
func IsinTab3(str string, tab []string) bool {
	for _, value := range tab {
		if str == value {
			return true
		}
	}
	return false
}