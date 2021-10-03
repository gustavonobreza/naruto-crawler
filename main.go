package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	characters := parseFileData()
	paragraph := make(chan []string)

	println("Starting...")

	urlBase := "https://naruto.fandom.com/wiki/"
	println("Characters count:", len(characters))

	for _, name := range characters {
		url := urlBase + name
		go Fetch(url, name, paragraph)
	}

	for i := 0; true; i++ {
		if i == len(characters) {
			println("Finished!")
			return
		}
		<-paragraph
		// res := <-paragraph
		// println(strings.Join(res, "\n"))
	}
}

func parseFileData() []string {
	file, err := os.Open("characters.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	sanitized := strings.ReplaceAll(string(byteValue), "\n", "")
	splittedNames := strings.Split(sanitized, ",")

	for i, v := range splittedNames {
		splittedNames[i] = strings.ReplaceAll(v, " ", "_")
	}
	return splittedNames
}
