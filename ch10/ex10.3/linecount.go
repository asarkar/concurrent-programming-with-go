package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

/*
3. The following listing downloads 30 web pages and counts the total number of lines on all the
documents sequentially. Convert this program to use concurrent programming, using a concurrency
pattern explained in this chapter.

package main

import (

	"fmt"
	"io"
	"net/http"
	"strings"

)

	func main() {
		const pagesToDownload = 30
		totalLines := 0
		for i := 1000; i < 1000 + pagesToDownload; i++ {
			url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
			fmt.Println("Downloading", url)
			resp, _ := http.Get(url)
			if resp.StatusCode != 200 {
				panic("Server’s error: " + resp.Status)
			}
			bodyBytes, _ := io.ReadAll(resp.Body)
			totalLines += strings.Count(string(bodyBytes), "\n")
			resp.Body.Close()
		}
		fmt.Println("Total lines:", totalLines)
	}

ANSWER: fork/join.
*/
func main() {
	const pagesToDownload = 30
	var wg sync.WaitGroup
	counts := make(chan int)

	fork(1000, 1000+pagesToDownload, counts, &wg)
	totalLines := join(counts)

	wg.Wait()
	close(counts)

	fmt.Println("Total lines:", <-totalLines)
}

func fork(start int, end int, counts chan int, wg *sync.WaitGroup) {
	for i := start; i < end; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		wg.Add(1)
		go countLines(url, counts, wg)
	}
}

func join(counts chan int) chan int {
	var numLines int
	totalLines := make(chan int)
	go func() {
		for i := range counts {
			numLines += i
		}
		totalLines <- numLines
	}()
	return totalLines
}

func countLines(url string, counts chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Downloading:", url)
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		panic("Server's error: " + resp.Status)
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	err := resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	counts <- strings.Count(string(bodyBytes), "\n")
}
