package main

import (
    "fmt"
    "rsc.io/quote"

    p "github.com/kwmt/go-print"
)

func main() {
    fmt.Println(quote.Hello())

    fmt.Println(p.Print("そうだ京都へ行こう！"))
}
