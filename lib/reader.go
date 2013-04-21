package lib

import (
    "math/rand"
)

//GetRandomGram returns a random beginner gram
func (ngrams *Ngrams) GetRandomGram() Ngram {
    choice := rand.Intn(len(ngrams.keys))
    return ngrams.ngrams[ngrams.keys[choice]]
}

//GetNextGram attempts to return the next ngram that begins with firstWord
//
//Returns a Ngram and bool. The bool will be set to false if no ngram is
//found that begins with firstWord, additionally if no ngram is found the
//return value of Ngram will nil.
func (ngrams *Ngrams) GetNextGram(string firstWord) (Ngram, bool) {
    foundNgram := false
    var ngram Ngram
    ngram = nil
    ngramList := ngrams.ngrams[firstWord]
    if ngramList != nil && ngramList.Len() > 0 {
        choice := rand.Intn(gramList.Len())
        pos := 0
        for e := ngramList.Front(); e != nil; e = e.Next() {
            if pos == choice {
                foundNgram = true
                ngram = e.Value.(Ngram)
            }
        }
    }
    return ngram,foundNgram
}
