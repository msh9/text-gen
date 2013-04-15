package lib

import (
    "testing"
    "text/scanner"
    "bytes"
)

const simpleTestString = "Hello world, my name is Michael Hughes. This is a simple test string. Goodbye!"
const testNgramSize = 2

func getSimpleTestReader() *scanner.Scanner {
    strBuffer := bytes.NewBufferString(simpleTestString)
    var reader scanner.Scanner
    reader.Init(strBuffer)
    return &reader
}

func TestConsume(t *testing.T) {
    //run test
    ngrams := InitMemory()
    reader := getSimpleTestReader()
    ngrams.Consume(reader, testNgramSize)

    //verify with spot checks
    if ngrams.ngrams["Hello"] == nil || ngrams.ngrams["string"] == nil {
        t.Fatal("Hello world or string. Goodbye! ngrams were not present in memory\n")
    }

    if ngrams.ngrams["Hello"].Len() != testNgramSize {
        t.Errorf("Ngrams not of size %d\n", testNgramSize)
    }

    ngram := ngrams.ngrams["string"].Front().Value.(Ngram)
    if !ngram.IsBeginner {
        t.Error("Beginning of sentence should have been marked beginner")
    }
    if ngram.IsStop {
        t.Error("Beginning of setence should have not been marked as a stop")
    }
}
