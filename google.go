package main

import (
	"fmt"
	"math/rand"
	"time"
)

//show servers OMIT
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

//end show servers OMIT

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

func SendQuery(query string) (results []Result) { // HL
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}
