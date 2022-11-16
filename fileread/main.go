package main

import (
	"fmt"
	"log"
	"os"
)

func fileRead(filePath string) string {
	oFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	ss
	defer oFile.Close()

	buff := make([]byte, 100)
	content := ""
	for no, err := oFile.Read(buff); err == nil; no, err = oFile.Read(buff) {
		if no > 0 {

			os.Stdout.Write(buff[0:no])
			content = fmt.Sprintf("%s%s", content, buff[0:no])

		}
	}
	return content

}
func main() {
	temp := fileRead("sample.txt")
	log.Println(temp)
}
