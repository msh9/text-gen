package main

import (
    "github.com/msh9/ngrams/lib"
    "os"
    "io"
    "strconv"
)

func main() {
    args := os.Args
    ngrams := lib.InitMemory()
    ngramSize, err := strconv.ParseInt(args[2],10,32)
    if err == nil {
        reader, _ := getReader(args[1])
        ngrams.Consume(reader, int(ngramSize))
        for i :=0; i < 3; i++ {
            ngram := ngrams.GetRandomBeginner()
            if ngram != nil {
                println(ngram.Values[0])
            }
        }
    }
}

func getReader(path string) (io.Reader, error) {
    fh, err := os.Open(path)
    return fh, err
}
