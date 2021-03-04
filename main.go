package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	createFile()
	createSha256()
}

func createFile() {
	var f, err = os.OpenFile("./file.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var _, errC = f.WriteString("Insert text here")
	if errC != nil {
		log.Fatal(errC)
	}
}

func createSha256() {
	var f, err = os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var h = sha256.New()
	io.Copy(h, f)

	var sum = h.Sum(nil)

	fmt.Printf("%x", sum)
}
