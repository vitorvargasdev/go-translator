package main

import (
	"fmt"
	"gotranslator/src/translator"
)

func main() {
	text, err := translator.Translate("Hello World", "en", "pt")
	if err != nil {
		panic(err)
	}

	fmt.Println(text)
}
