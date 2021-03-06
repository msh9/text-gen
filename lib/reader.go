package lib

import (
    "math/rand"
)

//GetRandomBeginner returns a random beginner gram
func (ngrams *Ngrams) GetRandomBeginner() *Ngram {
    var ngram *Ngram
    choice := rand.Intn(len(ngrams.beginnerKeys))
    ngramList := ngrams.ngrams[ngrams.beginnerKeys[choice]]
    choice = rand.Intn(ngramList.Len())
    i := 0
    for e := ngramList.Front(); e != nil; e = e.Next() {
        if i == choice {
            ngram = e.Value.(*Ngram)
        }
        i++
    }
    return ngram
}

//GetNext attempts to return the next ngram that begins with firstWord
//
//Returns a Ngram and bool. The bool will be set to false if no ngram is
//found that begins with firstWord, additionally if no ngram is found the
//return value of Ngram will be invalid.
func (ngrams *Ngrams) GetNext(firstWord string) (*Ngram, bool) {
    foundNgram := false
    var ngram *Ngram
    ngramList := ngrams.ngrams[firstWord]
    if ngramList != nil && ngramList.Len() > 0 {
        choice := rand.Intn(ngramList.Len())
        pos := 0
        for e := ngramList.Front(); e != nil; e = e.Next() {
            if pos == choice {
                foundNgram = true
                ngram = e.Value.(*Ngram)
            }
            pos++
        }
    }
    return ngram,foundNgram
}

//GetNextIgnorePunc attempts to return the next ngram that begins with firstWord ignore any leading or end non-word characters

//Returns a Ngram and bool. The bool will be set to false if no ngram is
//found that begins with firstWord, additionally if no ngram is found the
//return value of Ngram will be invalid.
//
//N.B. Currently not implemented--uses logic of regular GetNext
func (ngrams *Ngrams) GetNextIgnorePunc(firstWord string) (*Ngram, bool) {
    return ngrams.GetNext(firstWord)
}

