package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
1. Modify our sequential-letter frequency program to produce a list of word frequencies rather than
letter frequencies. You can use the same URLs for the RFC web pages as were used in listing 3.3.
Once it's finished, the program should output a list of words with the frequency with which each
word appears in the web page. Here's some sample output:

$ go run wordfrequency.go
the -> 5
a -> 8
car -> 1
program -> 3
*/

func countLetters(url string, frequency map[string]int) {
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		panic("Server returning error status code: " + resp.Status)
	}
	defer resp.Body.Close() //nolint:errcheck
	body, _ := io.ReadAll(resp.Body)
	words := strings.Fields(string(body))
	for _, w := range words {
		w = strings.ToLower(w)
		i := frequency[w]
		frequency[w] = i + 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
	}
	for w, i := range frequency {
		fmt.Printf("%s-%d ", w, i)
	}
}
