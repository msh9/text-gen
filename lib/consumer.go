//Package lib contains funcs and structures to handle ngram generation
package lib

import (
    "text/scanner"
    "container/list"
    "unicode"
    "strings"
)

//Consume consumptions input from a reader and inserts n-size ngrams into a Ngrams struct
//
//Example input: "The quick brown fox jumped over the fence to avoid the zany doxen."
//To split into a two gram we would split on whitespace and get ['The','quick'] followed
//by ['quick','brown'], and so on. 
func (ngrams *Ngrams) Consume(reader *scanner.Scanner, n int) {
    window := list.New()
    stopList := [3]string {"?","!","."}
    beginnerSet := make(map[string] bool)
    //fill the initial window
    tok := reader.Scan()
    for i := 0; tok != scanner.EOF && i < n; i++ {
        window.PushBack(reader.TokenText())
        tok = reader.Scan()
    }
    ngrams.insert(window,n,&stopList,beginnerSet)
    //now slide the window
    for tok != scanner.EOF {
        window.Remove(window.Front())
        window.PushBack(reader.TokenText())
        tok = reader.Scan()
        ngrams.insert(window,n,&stopList,beginnerSet)
    }

    totalBeginners := len(beginnerSet)
    beginnerKeys := make([]string,totalBeginners)
    i := 0
    for k := range beginnerSet {
        beginnerKeys[i] = k
        i++
    }
    ngrams.beginnerKeys = beginnerKeys
}

//insert saves the current window as a n-sized ngram in the ngrams data structure
//
//We accept a list of strings and make a slice of strings to permanently store the
//strings in memory. We also check to see if the ngram passed in starts with a upper
//case letter in which case it's marked as a beginner. If the ngram passed in ends
//with a character in the stopList argument then it is marked as an ender. Finally,
//for all beginner the beginning value is added to the beginnerSet map.
func (ngrams *Ngrams) insert(window *list.List, n int, stopList *[3]string, beginnerSet map[string] bool) *Ngram {
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
        hasSuffix := false
        for i := 0; i < len(stopList); i++ {
            hasSuffix = hasSuffix || strings.HasSuffix(last,stopList[i])
        }
        ngram.IsStop = hasSuffix
        ngram.IsBeginner = unicode.IsUpper(rune(values[0][0]))
        if ngram.IsBeginner {
            beginnerSet[values[0]] = true
        }
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

