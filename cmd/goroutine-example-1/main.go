package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	runApp(3)
}

func runApp(size int) {
	channel := make(chan string)
	var wait sync.WaitGroup
	for index := 0; index < size; index++ {
		wait.Add(1)
		worker := Worker{name: fmt.Sprintf("consumer-%d", index+1), wait: &wait}
		go worker.consume(&channel)
	}
	wait.Add(1)

	worker := Worker{name: "supplier", wait: &wait}
	go worker.supplier(&channel)

	wait.Wait()
}

type Worker struct {
	name string
	wait *sync.WaitGroup
}

func (wk Worker) consume(channel *chan string) {
	for {
		item, more := <-*channel
		if more {
			log.Println(wk.name, "receive", item)
		} else {
			log.Println(wk.name, "finish")
			wk.wait.Done()
			return
		}
	}
}

func (wk Worker) supplier(channel *chan string) {
	wk.supply(channel,
		"io.netty:netty-buffer",
		"io.projectreactor:reactor-core:3.2.1.RELEASE",
		"org.jctools:jctools-core",
		"org.slf4j:slf4j-api")
	log.Println(wk.name, "finish")
	close(*channel)
	wk.wait.Done()
}

func (wk Worker) supply(channel *chan string, items ...string) {
	var index int
	index = 0
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			item := items[index]
			log.Println(wk.name, "supply", item)
			*channel <- item
			index++
		}
		if index == len(items) {
			break
		}
	}
	ticker.Stop()
}
