# go-linker-vars-example
This project is a simple example of using `go build` linker flags to set package variables deeper than just main which is well covered. I created this repo because I didn't find a good example on the topic of deeper variables which could be imported into multiple main packages. Full documentation for the [go command](https://golang.org/cmd/go/).

**NOTE: This repo example repo is built using Go 1.5.2, I know for a fact the flags changed since 1.4.x. However the same concepts should are available just check the documentation**

## Simple Configuration of `main` package variables
```
cd cmd/
go build -ldflags "-X main.MainVar=hihi"

$ ./cmd
MainVar: hihi
Version: 
```

## Setting pathed package variables with linker flags
The `github.com/ropes/go-linker-vars-example/src/version` can then be imported into any Go package to access its exported variables.
```
cd cmd/
go build -ldflags "-X github.com/ropes/go-linker-vars-example/src/version.Version=hihi"
./cmd

Output:

MainVar: 
Version: hihi
```

# Getting fancy with it!
Import the `version` package to where it will be exposed or logged:
```go
package main

import "fmt"
import "github.com/ropes/go-linker-vars-example/src/version"

func main() {
	fmt.Printf("Git Tag:   %s\n", version.GitTag)
	fmt.Printf("Build User: %s\n", version.BuildUser)
	fmt.Printf("Version:   %s\n", version.Version)
}
```

Inject VCS and build information into the an importable package which can be referenced by multiple binaries via the `go ``build` or `install` commands. 
```
cd cmd/complex
gittag=$(git describe --tags)

go build -ldflags "-X github.com/ropes/go-linker-vars-example/src/version.GitTag=${gittag} 
-X github.com/ropes/go-linker-vars-example/src/version.BuildUser=${USER} 
-X github.com/ropes/go-linker-vars-example/src/version.Version=v0.0.1"

Output:
$./complex 
Git Tag:    v0.0.1-1-gd75e8db
Build User: josh
Version:    v0.0.1
```

