package main

import (
	"log"

	"github.com/DanielTakeyama/eagleEyes/internal/core"
)

func main() {
	//fmt.Println("Hello, world..")
	wordpress := []string{"wp-admin"}
	wordlist := core.Wordlist{
		Wordpress: wordpress,
		Default:   wordpress,
	}

	err := wordlist.ScanWordpress("https://www.codeflakes.com.br")
	if err != nil{
		log.Printf("%v", err)
	}
}
