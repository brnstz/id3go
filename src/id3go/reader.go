package id3go

import (
    "fmt"
    "os"
    "bytes"
)

func byteString(b []byte) (string) {
    pos := bytes.IndexByte(b, 0)

    if pos == -1 {
        pos = len(b)
    }

    return string(b[0:pos])
}


func ReadId3V1Tag(filename string) (map[string] string, string) {
    buff_ := make([]byte, tagSize)

    f, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Sprintf("%v", err)
    }
    defer f.Close()

    // Read last 128 bytes of file to see ID3 tag
    f.Seek(-tagSize, 2)
    f.Read(buff_)

    // First 3 characters are static "TAG" FIXME: check to be sure it matches
    buff := buff_[tagStart:]

    id3tag := map[string] string {}

    id3tag["title"]    = byteString(buff[0         : titleEnd])
    id3tag["artist"]   = byteString(buff[titleEnd  : artistEnd])
    id3tag["album"]    = byteString(buff[artistEnd : albumEnd])
    id3tag["year"]     = byteString(buff[albumEnd  : yearEnd])
    id3tag["comment"]  = byteString(buff[yearEnd   : commentEnd])

    // Special case. If next-to-last comment byte is zero, then the last
    // comment byte is the track number
    if ( buff[commentEnd - 2] == 0) {
        id3tag["track"] = fmt.Sprintf("%d", buff[commentEnd - 1])
    }
    genre_code := buff[commentEnd]
    id3tag["genre"] = fmt.Sprintf("%d", genre_code)
    id3tag["genre_name"] = codeToName[genre_code]

    for k, v := range(id3tag) {
        fmt.Printf("%s => %s\n", k, v)
    }

    return id3tag, ""
}
