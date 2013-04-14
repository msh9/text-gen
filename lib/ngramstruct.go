package lib

import (
    "container/list"
    "text/scanner"
)

type Ngram struct {
    Values []string
    IsStop bool
    IsBeginner bool
}

type Ngrams struct {
    ngrams map[string] *list.List
}

type NgramReader interface {
    GetRandomGram() Ngram
    GetNextGram(string) (Ngram, bool)
}

type NgramBuilder interface {
    Consume(reader *scanner.Scanner, n int) 
}

func InitMemory() *Ngrams {
    ngrams := new(Ngrams)
    ngrams.ngrams = make(map[string] *list.List)
    return ngrams
}

