package lib

import (
    'container/list'
    'unicode'
    'strings'
)

type Ngram struct {
    Values []string
    IsStop boolean
    IsBeginner boolean
}

type Ngrams struct {
    ngrams map[string] *list.List
}

type NgramReader interface {
    GetRandomGram() Ngram
    GetNextGram(string) (Ngram,boolean)
}

type NgramBuilder interface {
    Consume(reader *scanner.Scanner, n int) int
}

function InitMemory() *Ngrams {
    ngrams := new(Ngrams)
    ngrams.ngrams = make(map[string]list.List)
    return ngrams
}

