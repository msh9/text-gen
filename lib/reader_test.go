package lib

import (
    "testing"
)

//getSimpleTestConsumer and testNgramSize appear in
//reader_test.go
func getBuiltNgrams() *Ngrams {
    ngrams := InitMemory()
    consumer := getSimpleTestConsumer()
    ngrams.Consume(consumer, testNgramSize)
    return ngrams
}

func TestGetRandomBeginner(t *testing.T) {
    ngrams := getBuiltNgrams()
    ngram := ngrams.GetRandomBeginner()
    if !ngram.IsBeginner {
        t.Error("No ngram returned from GetRandomBeginner or it wasn't marked as a beginner")
    }
    t.Logf("Ngram opening word was: %s", ngram.Values[0])
}

//TestGetNext tests the basic case of GetNext, looking for the first opener
//in the simple test case string
func TestGetNext(t *testing.T) {
    ngrams := getBuiltNgrams()
    ngram,_ := ngrams.GetNext("Hello")
    if !ngram.IsBeginner {
        t.Error("No ngram returned from GetNext or it wasn't marked as a beginner")
    }
    t.Logf("Ngram opening word was: %s", ngram.Values[0])
}

//TestGetNext tests another basic case of GetNext, looking for a none beginner ngram
func TestGetNextNonBeginner(t *testing.T) {
    ngrams := getBuiltNgrams()
    ngram,_ := ngrams.GetNext("world")
    if ngram.IsBeginner {
        t.Error("No ngram returned from GetNext or it was marked as a beginner")
    }
    t.Logf("Ngram opening word was: %s", ngram.Values[0])
}
//TestGetNext tests another basic case of GetNext, test for a value that doesn't
//exist as an opener in the ngram set
func TestGetNextEmpty(t *testing.T) {
    ngrams := getBuiltNgrams()
    _,existed := ngrams.GetNext("Somebigwordthatdoesnotexist")
    if existed {
        t.Error("GetNext should not have returned true for a word not in the ngram memory")
    }
}
