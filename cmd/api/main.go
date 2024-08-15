package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	// do we have our env args?
	fmt.Println("api port: ", os.Getenv("API_PORT"))
	fmt.Println("TZ: ", os.Getenv("TZ"))
}
