package main

import (
	"fmt"
	"math/rand"
	"time"
)

//show servers OMIT
var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2") // HL
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2") // HL
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2") // HL
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

func SendQuery(query string) (results []Result) {
	//show timeout OMIT
	c := make(chan Result)
	go func() { c <- QueryReplicas(query, Web1, Web2) }()     // HL
	go func() { c <- QueryReplicas(query, Image1, Image2) }() // HL
	go func() { c <- QueryReplicas(query, Video1, Video2) }() // HL

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
	//end show timeout OMIT
}

func QueryReplicas(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}
