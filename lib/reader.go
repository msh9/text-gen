//Package lib contains functions and structures to handle ngram generation
package lib

import (
    "text/scanner"
    "container/list"
)

//Reader defines the functions and structures
//necessary to split text into n-grams.

//Consume consumptions input from a reader and inserts n-size ngrams into a Ngrams struct
//
//Example input: "The quick brown fox jumped over the fence to avoid the zany doxen."
//To split into a two gram we would split on whitespace and get ['The','quick'] followed
//by ['quick','brown'], and so on. This suggests a blocking producer->consumer pattern
function (ngrams *Ngrams) Consume(reader *scanner.Scanner, n int) int {
    window := list.New()

    //fill the initial window
    tok := s.Scan()
    for i := 0; tok != scanner.EOF && i < n; i++ {
        l.PushBack(s.TokenText())
        tok = s.Scan()
    }
    ngrams.insert(l,n)
    //now slide the window
    for tok != scanner.EOF {
        l.Remove(l.Front())
        l.PushBack(s.TokenText())
        tok = s.Scan()
        ngrams.insert(l,n)
    }
}

//insert saves the current window as a n-sized ngram in the ngrams data structure
//
//We accept a list of strings and make a slice of strings to permanently store the
//strings in memory
function (ngrams *Ngrams) insert(values *list.List, n int) *Ngram {
    values := make([]string,n)

    for i:= 0, e := window.Front; i < n && e != nil; i++, e = e.Next() {
        values[i] = e.Value.(string)
    }

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

