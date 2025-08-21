package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

/*
4. In the previous chapter, in exercise 3.1, we developed a program to output the frequencies of
words from downloaded web pages. If you used a shared memory map to store the word frequencies,
access to the shared map would need to be protected. Can you use a mutex to guarantee exclusive
access to the map?

Note: In this program we have a timer at the end which you might need to adjust
depending on how fast your internet connection is.
In later chapters we cover how to wait for threads to complete their work.
*/
func countLetters(url string, frequency map[string]int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	if resp.StatusCode != 200 {
		panic("Server's error: " + resp.Status)
	}
	defer resp.Body.Close() //nolint:errcheck
	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	mutex.Lock()
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower] += 1
	}
	mutex.Unlock()
	fmt.Println("Completed:", url)
}

func main() {
	mutex := sync.Mutex{}
	var frequency = make(map[string]int)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	time.Sleep(10 * time.Second)
	mutex.Lock()
	for k, v := range frequency {
		fmt.Println(k, "->", v)
	}
	mutex.Unlock()
}
