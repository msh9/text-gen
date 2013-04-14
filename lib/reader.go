//Package lib contains funcs and structures to handle ngram generation
package lib

import (
    "text/scanner"
    "container/list"
    "unicode"
    "strings"
)


//Reader defines the funcs and structures
//necessary to split text into n-grams.

//Consume consumptions input from a reader and inserts n-size ngrams into a Ngrams struct
//
//Example input: "The quick brown fox jumped over the fence to avoid the zany doxen."
//To split into a two gram we would split on whitespace and get ['The','quick'] followed
//by ['quick','brown'], and so on. This suggests a blocking producer->consumer pattern
func (ngrams *Ngrams) Consume(reader *scanner.Scanner, n int) {
    window := list.New()
    stopList := [3]string {"?","!","."}
    //fill the initial window
    tok := reader.Scan()
    for i := 0; tok != scanner.EOF && i < n; i++ {
        window.PushBack(reader.TokenText())
        tok = reader.Scan()
    }
    ngrams.insert(window,n,&stopList)
    //now slide the window
    for tok != scanner.EOF {
        window.Remove(window.Front())
        window.PushBack(reader.TokenText())
        tok = reader.Scan()
        ngrams.insert(window,n,&stopList)
    }
}

//insert saves the current window as a n-sized ngram in the ngrams data structure
//
//We accept a list of strings and make a slice of strings to permanently store the
//strings in memory
func (ngrams *Ngrams) insert(window *list.List, n int, stopList *[3]string) *Ngram {
    values := make([]string,n)
    e := window.Front()
    for i := 0; i < n; i++ {
        if e == nil {
            break
        }
        values[i] = e.Value.(string)
        e = e.Next()
    }

    var ngram *Ngram
    if len(values) != 0 {
        ngram = new(Ngram)
        ngram.Values = values
        last := values[len(values) - 1]
        var hasSuffix bool
        for i := 0; i < len(stopList); i++ {
            hasSuffix = hasSuffix && strings.HasSuffix(last,stopList[i])
        }
        ngram.IsStop = hasSuffix
        ngram.IsBeginner = unicode.IsUpper(rune(values[0][0]))
       if ngrams.ngrams[values[0]] == nil {
           ngramList := list.New()
            ngramList.PushBack(ngram)
            ngrams.ngrams[values[0]] = ngramList
        } else {
            ngrams.ngrams[values[0]].PushBack(ngram)
        }
    }
    return ngram
}

