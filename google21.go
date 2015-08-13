package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := SendQuery("golang") // HL
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

type Result string

//show fakeSearch OMIT
type Search func(query string) Result // HL

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result %q\n", kind, query))
	}
}

//end show fakeSearch OMIT

func SendQuery(query string) (results []Result) {
	//show timeout OMIT
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond) // HL
	for i := 0; i < 3; i++ {
		select {
		case result := <-c: // HL
			results = append(results, result)
		case <-timeout: // HL
			fmt.Println("timed out")
			return
		}
	}
	return
	//end show timeout OMIT
}
