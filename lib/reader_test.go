package lib
//
//import (
//    "testing"
//    "text/scanner"
//    "bytes"
//)
//
//const simpleTestString = "Hello world, my name is Michael Hughes. This is a simple test string. Goodbye!"
//const testNgramSize = 2
//const numOfSimpleTestHelloOpeners = 1
//
//func getSimpleTestConsumer() *scanner.Scanner {
//    strBuffer := bytes.NewBufferString(simpleTestString)
//    var reader scanner.Scanner
//    reader.Init(strBuffer)
//    reader.Mode = scanner.ScanRawStrings | scanner.ScanStrings | scanner.ScanChars | scanner.ScanComments | scanner.ScanIdents
//    return &reader
//}
//
//func TestConsume(t *testing.T) {
//    //run test
//    ngrams := InitMemory()
//    consumer := getSimpleTestConsumer()
//    ngrams.Consume(consumer, testNgramSize)
//
//    //verify with spot checks
//    if ngrams.ngrams["Hello"] == nil || ngrams.ngrams["string"] == nil {
//        t.Fatal("Hello world or string. Goodbye! ngrams were not present in memory\n")
//    }
//
//    if ngrams.ngrams["Hello"].Len() != numOfSimpleTestHelloOpeners {
//        t.Errorf("Ngrams not of size %d\n", numOfSimpleTestHelloOpeners)
//    }
//
//    ngram := ngrams.ngrams["Hello"].Front().Value.(*Ngram)
//    if !ngram.IsBeginner {
//        t.Error("Beginning of sentence should have been marked beginner")
//    }
//    if ngram.IsStop {
//        t.Error("Beginning of setence should have not been marked as a stop")
//    }
//}
//
//func BenchmarkConsume(b *testing.B) {
//    n := 60
//    b.StopTimer()
//    b.ResetTimer()
//    ngrams := InitMemory()
//    consumer := getSimpleTestConsumer()
//    for i := 0; i < n; i++ {
//        b.StartTimer()
//        ngrams.Consume(consumer, testNgramSize)
//        b.StopTimer()
//    }
//    b.N = n
//}
