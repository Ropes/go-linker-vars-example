package main

import "fmt"
import "github.com/ropes/go-linker-vars-example/src/version"

var MainVar string

func main() {

	fmt.Printf("MainVar: %s\n", MainVar)

	fmt.Printf("Version: %s\n", version.Version)

}
