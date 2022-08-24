package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fizzywhizbang/gorot"
)

func main() {

	encode := flag.Bool("e", false, "true to encode else decode")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	var output string
	if !*encode {
		output = gorot.Decode(text)
	} else {
		output = gorot.Encode(text)
	}

	fmt.Println(output)
}
