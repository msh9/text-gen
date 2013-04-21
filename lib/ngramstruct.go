package lib

import (
    "container/list"
    "text/scanner"
)

//Ngram holds the data for a single ngram to be stored
//
//Ngram exports a Values field which contains the actual
//ngram stored in a []string slice. Ngram also exports
//precomputed values indicating whether the ngram came
//at the beginning and\or end of an English sentence.
type Ngram struct {
    Values []string
    IsStop bool
    IsBeginner bool
}

//Ngrams is a wrapper struct that contains a simple in-memory store for ngrams
type Ngrams struct {
    ngrams map[string] *list.List
    beginnerKeys []string 
}

//NgramReader defines the methods that any object which returns 
//ngrams should support
type NgramReader interface {
    GetRandomBeginner() Ngram
    GetNext(string) (Ngram, bool)
}

//NgramBuilder defines method(s) that are expected for use in building a set of ngrams
type NgramBuilder interface {
    Consume(reader *scanner.Scanner, n int)
}

//InitMemory creates a basic in-memory ngram storage unit
func InitMemory() *Ngrams {
    ngrams := new(Ngrams)
    ngrams.ngrams = make(map[string] *list.List)
    return ngrams
}

