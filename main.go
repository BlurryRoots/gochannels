package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

var group_eater sync.WaitGroup
var group_feeder sync.WaitGroup

var stream = make (chan string)

func generate_interval (factor float32) time.Duration {
	interval := time.Duration (rand.Float32() * factor) * time.Second
	return interval
}

func feed (wobble string, k []int) {
	defer group_feeder.Done ()
	for _, i := range k {
		fmt.Printf ("%s: here take it! %v\n", wobble, i)
		stream <- fmt.Sprintf ("%v by %v", i, wobble)
		time.Sleep (generate_interval (0.618))
	}
}

func eat (gobble string) {
	defer group_eater.Done()
	for v := range stream {
		time.Sleep(generate_interval (1.618))
		fmt.Printf("%s: mhh yummi! %v\n", gobble, v)
	}
}

func goroutines () {
	rand.Seed (time.Now ().Unix ())

	var k = []int {1,3,3,7}
	var feeder = []string {"jason", "mona", "anton", "alesya"}
	var eater = []string {"molly", "karl", "hank", "lisa"}

	for _, v := range feeder {
		group_feeder.Add (1)
		go feed (v, k)
	}
	for _, name := range eater {
		group_eater.Add (1)
		go eat (name)
	}

	group_feeder.Wait ()
	close (stream)
	group_eater.Wait()	
}

func main () {
	goroutines ();
}