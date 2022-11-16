package main

import "fmt"

func main() {
	emp := make(map[string]string)

	emp["id1"] = "Thottavaram"
	fmt.Println(emp)

	_, val := emp["id1"]
	fmt.Println(val)

}
