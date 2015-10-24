# stemmer
A library in golang that implements the PorterStemmer algorithm for stemming words.

## usage

```go
package main

import "github.com/caneroj1/stemmer"

func main() {
  str := "running"

  // stem a single word
  stem := stemmer.Stem(str)


  strings := []string{
    "playing",
    "skies",
    "singed",
  }

  // stem a list of words
  stems := stemmer.StemMultiple(strings)

  // stem a list of words in place, modifying the original slice
  stemmer.StemMultipleMutate(strings)

  // stem a list of words concurrently. this also stems in place, modifying
  // the original slice.
  // NOTE: the order of the strings is not guaranteed to be the same.
  stemmer.StemMultipleConcurrent(strings)

}
```
