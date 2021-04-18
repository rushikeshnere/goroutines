package main

import (
	"fmt"
	"sync"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
)

var startNumber = 1

type num struct {
	Number int `json:"number"`
}

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go runServer(&wg)
	
	wg.Add(1)
	go fetchNumbers(&wg, 30, channel)

	wg.Add(1)
	go peformSquareOperation(&wg, channel)

	wg.Wait()
}

func fetchNumbers(wg *sync.WaitGroup, limit int, channel chan int) {
	defer wg.Done()
	time.Sleep(8 * time.Second)

	for i :=1; i<=20; i++ {
		response, err := http.Get("http://localhost:9000")
		if(err != nil) {
			fmt.Println(err)
			panic(err)
		}
		
		body, error := ioutil.ReadAll(response.Body)
   		if error != nil {
      		panic(error)
    	}
//		fmt.Println("Body = ", string(body))

		var n num
		json.Unmarshal(body, &n)
		channel <- n.Number
	}
}

func peformSquareOperation(wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	for i := range channel {
		fmt.Println("Square of ", i, " = ", i * i)
	}
}

func runServer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting server")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":9000", nil)
	fmt.Println("Ending server")
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	e := &num{Number : startNumber}	
	json.NewEncoder(writer).Encode(e)
	startNumber ++
}