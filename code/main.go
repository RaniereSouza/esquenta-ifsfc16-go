package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func buildHelloMessage(name string) string {
	return "Hello, " + name + "!"
}

type Course struct {
	Name        string `json:"course"`
	Description string `json:"description"`
	Price       int    `json:"price"` // cents
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf(
		"Course:\n  Name: %s, Description: %s, Price: %d",
		c.Name, c.Description, c.Price,
	)
}

func courseJsonHttpHandler(course Course) func(resW http.ResponseWriter, req *http.Request) {
	return func(resW http.ResponseWriter, req *http.Request) {
		json.NewEncoder(resW).Encode(course)
	}
}

func helloHttpHandler(name string) func(resW http.ResponseWriter, req *http.Request) {
	return func(resW http.ResponseWriter, req *http.Request) {
		resW.Write([]byte(buildHelloMessage(name)))
	}
}

func countTo(limit int, channel chan [2]int, counterId int) {
	for i := 0; i < limit; i++ {
		channel <- [2]int{i, counterId}
		time.Sleep(time.Second)
	}
}

func main() {
	//==================== Hello World with system argument ====================//
	args, name := os.Args[1:], "World"

	if len(args) > 0 {
		name = args[0]
	}

	fmt.Println(buildHelloMessage(name))

	//=========================== struct and methods ===========================//
	course := Course{
		Name:        "Go 101",
		Description: "Golang introduction course",
		Price:       10000,
	}

	fmt.Println("\n" + course.GetFullInfo())

	//==================== HTTP endpoints and simple server ====================//
	http.HandleFunc("/", courseJsonHttpHandler(course))
	http.HandleFunc("/hello", helloHttpHandler(name))
	go http.ListenAndServe(":5000", nil) // running into a goroutine to unblock the rest of the code

	//========================= channels and goroutines ========================//
	countLimit := 10
	countingChannel := make(chan [2]int)
	numGoroutines := 3

	fmt.Printf("\nParallel counters (%d threads):\n", numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		counterId := i + 1
		go countTo(countLimit, countingChannel, counterId)
	}

	for res := range countingChannel {
		fmt.Printf("value %d from counter #%d\n", res[0], res[1])
	}
}
