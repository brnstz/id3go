package main

import "flag"
import "fmt"
import "os"

func readFile(filename string) (string) {
    buff := make([]byte, 128, 128)
    
    f, err := os.Open(filename, os.O_RDONLY, 0)
    if err != nil {
        return fmt.Sprintf("%v", err)
    }
    defer f.Close()

    // Read last 128 bytes of file to see ID3 tag
    f.Seek(-128, 2)
    f.Read(buff)

    fmt.Println(buff)

    return ""
}

func main() {
    //var myint *string = flag.String("album", "Transformer", "album name")

    flag.Parse()

    for _, filename := range(flag.Args()) {
        readFile(filename)
    }

    //fmt.Printf("Its value is: %v %v %v %v", *myint, flag.Args(), flag.NArg(), flag.NFlag())
}
