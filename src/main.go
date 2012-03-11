package main

import (
    "flag"
    "fmt"
    "id3go"
)

func main() {
    flag.Parse()

    for _, filename := range(flag.Args()) {
        res, err := id3go.ReadId3V1Tag(filename)
        if (err != "") {
            fmt.Println(err)
        } else {
            fmt.Println(res)
        }
    }
}
