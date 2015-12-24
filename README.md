# go-linker-vars-example
This project is a simple example of using Go build linker flags to set package variables deeper than just main. I created this repo because I didn't find a good example on the topic of deeper variables which could be imported into multiple main packages. Full documentation for the [go command](https://golang.org/cmd/go/).

## Simple Configuration of `main` package variables
```
cd cmd/
go build -ldflags "-X main.MainVar=hihi"

Output:

MainVar: hihi
Version: 
```

## Setting nested package variabls with linker flags
```
cd cmd/
go build -ldflags "-X github.com/ropes/go-linker-vars-example/src/version.Version=hihi"

Output:

MainVar: 
Version: hihi
```


