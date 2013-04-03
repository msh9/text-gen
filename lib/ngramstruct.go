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
    Insert(*[]string) *Ngram
}

function InitMemory() *Ngrams {
    ngrams := new(Ngrams)
    ngrams.ngrams = make(map[string]list.List)
    return ngrams
}

function (ngrams *Ngrams) Insert(values []string) *Ngram {
    var ngram *Ngram
    if len(values) != 0 {
        ngram = new(Ngram)
        ngram.Values = values
        last := values[len(values) - 1]
        var hasSuffix boolean
        for i := 0; i < len(StopList); i++ {
            hasSuffix = hasSuffix && strings.HasSuffix(last,StopList)
        }
        ngram.IsStop = hasSuffix
        ngram.IsBeginner = unicode.IsUpper(rune(values[0][0]))
       if ngrams.ngrams[values[0]] == nil {
            ngramList = list.New()
            ngramList.PushBack(ngram)
            ngrams.ngrams[values[0]] = ngramList
        } else {
            ngrams.ngrams[values[0]].PushBack(ngram)
        }
    }
    return ngram
}

function(ngrams *Ngrams) GetRandomGram() *Ngram {
    
