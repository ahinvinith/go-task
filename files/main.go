package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GOlang.")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylocgofile.txt") //Writing a file
	checkNilError(err)

	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	checkNilError(err)

	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Length is : ", length)
	defer file.Close()

	readFile("./mylocgofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Text data inside the file is : \n", string(dataByte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
