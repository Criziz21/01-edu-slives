package main

import(
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if(len(arguments) > 0) {
		fmt.Print(arguments[0])
	} else {
		fmt.Print("doesn't work")
	}
}