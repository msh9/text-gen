package main

import (
    "github.com/msh9/ngrams/lib"
    "os"
    "text/scanner"
    "strconv"
)

func main() {
    args := os.Args
    ngrams := lib.InitMemory()
    ngramSize, err := strconv.ParseInt(args[2],10,32)
    if err == nil {
        scanner, _ := getScanner(args[1])
        ngrams.Consume(scanner, int(ngramSize))
        for i :=0; i < 3; i++ {
            ngram := ngrams.GetRandomBeginner()
            if ngram != nil {
                println(ngram.Values[0])
            }
        }
    }
}

func getScanner(path string) (*scanner.Scanner, error) {
    fh, err := os.Open(path)
    var reader scanner.Scanner
    if err == nil {
        reader.Init(fh)
        reader.Mode = scanner.ScanRawStrings | scanner.ScanStrings | scanner.ScanChars | scanner.ScanComments | scanner.ScanIdents
    }
    return &reader, err
}
