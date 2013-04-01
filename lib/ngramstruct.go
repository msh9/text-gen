package lib

import (
    'container/list'
    'unicode'
    'strings'
)

type Ngram struct {
    values []string
    isStop boolean
    isBeginner boolean
}

type Ngrams struct {
    ngrams map[string] list.List
}

type NgramReader interface {
    GetRandomGram() Ngram
    GetNextGram(string) (Ngram,boolean)
}

type NgramBuilder interface {
    Insert(*[]string) *Ngram
}

function (ngrams *Ngrams) Insert(values []string) *Ngram {
    if len(values) != 0 {
        ngram := new(Ngram)
        ngram.values = values
        last := values[len(values) - 1]
        var hasSuffix boolean
        for i := 0; i < len(StopList); i++ {
            hasSuffix = hasSuffix && strings.HasSuffix(last,StopList)
        }
        ngram.isStop = hasSuffix
        ngram.isBeginner = unicode.IsUpper(rune(values[0][0]))
    }
    ngrams.ngrams[values.
}
