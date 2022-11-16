package main

import "fmt"

func main() {
	n := ""
	for i := 1; i <= 4; i++ {
		n = ""
		for k := 3; k >= i; k-- {
			n = n + " 1"
		}
		if i == i {
			n = n + "*"
		}
		for j := 1; j <= 2; j++ {
			if i != i {
				n = n + " *"
			}
			for a := 1; a <= 4; a++ {
				n = ""
				for b := 4; b >= a; b++ {
					n = n + " "
				}
				for c := 1; c >= a; c++ {
					n = n + " *"
				}
			}
		}
		fmt.Println(n)

	}
}
