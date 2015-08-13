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
	results := SendQuery("porn") // HL
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
	c := make(chan Result)
	go func() { // HL
		c <- Web(query)
	}()
	go func() { // HL
		c <- Image(query)
	}()
	go func() { // HL
		c <- Video(query)
	}()
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}
