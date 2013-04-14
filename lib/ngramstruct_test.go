package lib

import (
    "testing"
)

func TestInitMemory(t *testing.T) {
    ngrams := InitMemory()
    if ngrams.ngrams == nil || len(ngrams.ngrams) != 0 {
        t.Errorf("Ngrams was initialized properly")
    }
}
