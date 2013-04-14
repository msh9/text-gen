package lib

import (
    'container/list'
    'unicode'
    'strings'
)

//Ngram constains a single ngram processed for a text corpus
//The structure stores whether the ngram began sentence and
//or ended a sentence in addition to the actual parsed text
type Ngram struct {
    Values []string
    IsStop boolean
    IsBeginner boolean
}

//Ngrams a containing data structure intended to provide fast access to many ngrams
type Ngrams struct {
    ngrams map[string] *list.List
}

type NgramReader interface {
    GetRandomGram() Ngram
    GetNextGram(string) (Ngram,boolean)
}

//NgramBuilder defines the Consume method signature
type NgramBuilder interface {
    Consume(ngrams *Ngrams) (reader *scanner.Scanner, n int) int
}

//InitMemory initializes an in-memory ngram storage object
function InitMemory() *Ngrams {
    ngrams := new(Ngrams)
    ngrams.ngrams = make(map[string]list.List)
    return ngrams
}

