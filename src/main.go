package main

import (
    "flag"
    "fmt"
    "id3go"
    "os"
    "path/filepath"
    "log"
)

// File visitor for doing paths recursively
type id3Visitor int

func (v id3Visitor) VisitFile (path string, f *os.FileInfo) {
    printTag(path)
}

func (v id3Visitor) VisitDir (path string, f *os.FileInfo) bool {
    fmt.Println(path)
    return true;
}

func printTag(filename string) {
    fmt.Println(filename)

    res, err := id3go.ReadId3V1Tag(filename)

    if (err != nil) {
        log.Print(err)
    }

    for k, v := range(res) {
        fmt.Printf("%s => %s\n", k, v)
    }
    fmt.Println()
}

func main() {
    flag.Parse()

    for _, filename := range(flag.Args()) {
        finfo, err := os.Stat(filename)

        if (err != nil) {
            log.Print(err)
            continue
        }

        if (finfo.Mode().IsRegular()) {
            printTag(filename)

        } else if (finfo.IsDir()) {
            v := new(id3Visitor)
            filepath.Walk(filename, func(path string, fileInfo os.FileInfo, err error) error {
                if err != nil {
                    log.Print(err)
                } else if fileInfo.IsDir() {
                    v.VisitDir(path, &fileInfo)
                } else {
                    v.VisitFile(path, &fileInfo)
                }
                return nil
            })
        }
    }
}
