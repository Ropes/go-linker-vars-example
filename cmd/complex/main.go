package main

import "fmt"
import "github.com/ropes/go-linker-vars-example/src/version"

func main() {
	fmt.Printf("Git Tag:    %s\n", version.GitTag)
	fmt.Printf("Build User: %s\n", version.BuildUser)
	fmt.Printf("Version:    %s\n", version.Version)
}
