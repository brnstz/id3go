package main

import "flag"
import "fmt"
import "os"

func readFile(filename string) {
    var buff []byte
    
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    // Read last 128 bytes of file to see ID3 tag
    f.Seek(128, 2)
    f.Read(buff, 128)

    fmt.Println(buff)
}

func main() {
    var myint *string = flag.String("album", "Transformer", "album name")

    flag.Parse()

    for _, filename := range(flag.Args()) {
    
    }

    fmt.Println(*myint)


    //fmt.Printf("Its value is: %v %v %v %v", *myint, flag.Args(), flag.NArg(), flag.NFlag())
}
