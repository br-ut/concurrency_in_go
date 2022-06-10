package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	mut.Lock()
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
	mut.Unlock()
}

func main() {
	fmt.Println("Demo welcome...")
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <url1> <url2>... <urln> ")
	}

	for _, url := range os.Args[1:] {
		go sendRequest("https://" + url)

		wg.Add(1)
	}

	wg.Wait()
}

//First Benchmark
//4 Seconds

//Second Benchmark with go routine i.e, concurrency
//1.9

//Third Benchmark with Mutex lock and unlock
//2.4
