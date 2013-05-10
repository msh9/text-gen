package lib

import (
    "testing"
    "bytes"
    "io"
)

const simpleTestString = "Hello world, my name is Michael Hughes. This is a simple test string. Goodbye!\n"
const testNgramSize = 2
const numOfSimpleTestHelloOpeners = 1

func getSimpleTestConsumer() io.Reader {
    reader := bytes.NewReader([]byte(simpleTestString))
    return reader
}

func TestConsume(t *testing.T) {
    //run test
    ngrams := InitMemory()
    consumer := getSimpleTestConsumer()
    ngrams.Consume(consumer, testNgramSize)

    //verify with spot checks
    if ngrams.ngrams["Hello"] == nil || ngrams.ngrams["string"] == nil {
        t.Fatal("Hello world or string. Goodbye! ngrams were not present in memory\n")
    }

    if ngrams.ngrams["Hello"].Len() != numOfSimpleTestHelloOpeners {
        t.Errorf("Ngrams not of size %d\n", numOfSimpleTestHelloOpeners)
    }

    ngram := ngrams.ngrams["Hello"].Front().Value.(*Ngram)
    if !ngram.IsBeginner {
        t.Error("Beginning of sentence should have been marked beginner")
    }
    if ngram.IsStop {
        t.Error("Beginning of setence should have not been marked as a stop")
    }

    beginnerKeys := ngrams.beginnerKeys
    if len(beginnerKeys) < 1 {
        t.Error("List of beginner keys should contain some value")
    }

    foundHello := false
    for i := 0; i < len(beginnerKeys); i++ {
        foundHello = foundHello || beginnerKeys[i] == "Hello"
    }
    if !foundHello {
        t.Errorf("%s should have been in beginner key list\n","Hello")
    }
}

func BenchmarkConsume(b *testing.B) {
    b.StopTimer()
    b.ResetTimer()
    ngrams := InitMemory()
    consumer := getSimpleTestConsumer()
    for i := 0; i < b.N; i++ {
        b.StartTimer()
        ngrams.Consume(consumer, testNgramSize)
        b.StopTimer()
    }
}
