//Package lib contains funcs and structures to handle ngram generation
package lib

import (
    "bufio"
    "io"
    "container/list"
    "unicode"
    "strings"
)

//Consume consumptions input from a reader and inserts n-size ngrams into a Ngrams struct
//
//Example input: "The quick brown fox jumped over the fence to avoid the zany doxen."
//To split into a two gram we would split on whitespace and get ['The','quick'] followed
//by ['quick','brown'], and so on. 
func (ngrams *Ngrams) Consume(reader io.Reader, n int) {
    window := list.New()
    stopList := [3]string {"?","!","."}
    beginnerSet := make(map[string] bool)
    //fill & slide the window
    bufferedReader := bufio.NewReader(reader)
    curLine, err := bufferedReader.ReadString('\n')

    //first set of loops fill the window and then insert the initial data
    i := 0
    fieldsRead := 0
    var fields []string
    for err == nil {
        fields = strings.Fields(curLine)
        for fieldsRead := 0; fieldsRead < len(fields); fieldsRead++ {
            window.PushBack(fields[fieldsRead])
            i++
            if i == n {
                ngrams.insert(window,n,&stopList,beginnerSet)
                break
            }
        }
        if i == n {
            break
        }
        curLine, err = bufferedReader.ReadString('\n')
    }

    //second set of loops run till we run out of content
    for err == nil {
        fields := strings.Fields(curLine)
        for j := fieldsRead; j < len(fields); j++ {
            window.PushBack(fields[j])
            window.Remove(window.Front())
            ngrams.insert(window,n,&stopList,beginnerSet)
        }
        curLine, err = bufferedReader.ReadString('\n')
        //bit of a hack here because after during the first iteration of the 
        //above loop we want to pick up where the initial window filling loop
        //left off--so we use the fieldsRead variable. After this loop's first
        //iteration though we want to always start as zero.
        fieldsRead = 0 
    }

    totalBeginners := len(beginnerSet)
    beginnerKeys := make([]string,totalBeginners)
    i = 0
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
        println("Made ngram with "+values[0])
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

