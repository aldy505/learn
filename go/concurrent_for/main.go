package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	arr := make([]byte, 500)
	io.ReadFull(rand.Reader, arr)
	fmt.Println(arr)
	forSearch(arr, 250)
	goSearch(arr, 250)
}

func forSearch(arr []byte, query byte) {
	fmt.Println("starting for search")
	for i := 0; i < len(arr); i++ {
		if arr[i] == query {
			fmt.Println("forsearch:", i)
		}
	}
}

func goSearch(arr []byte, query byte) {
	fmt.Println("starting go search")
	var chunks [][]byte
	chunks = append(chunks, arr[0:99], arr[100:199], arr[200:299], arr[300:399], arr[400:])
	go goSearchUtil(chunks[0], query)
	goSearchUtil(chunks[1], query)
	goSearchUtil(chunks[2], query)
	goSearchUtil(chunks[3], query)
	goSearchUtil(chunks[4], query)
}

func goSearchUtil(chunk []byte, q byte) {
	for i := 0; i < len(chunk); i++ {
		if chunk[i] == q {
			fmt.Println("gosearch:", i)
			break
		}
	}
}
