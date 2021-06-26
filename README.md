# DGOC
a simple command handler for [discordgo](https://github.com/bwmarrin/discordgo)

## installation
```
go get github.com/iAtomPlaza/dgoc
```

## usage

- initialization
```go
package main

import (
    "github.com/iAtomPlaza/dgoc"
    "github.com/bwmarrin/discordgo"
)

func main() {

    // create a new instance of dgoc
    // 'session' being created with 'discordgo.New()'
    c := dgoc.New(session)
}
```