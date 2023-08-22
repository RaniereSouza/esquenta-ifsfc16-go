package main

import (
	"fmt"
	"main/counter"
	"main/course"
	"main/hello"
	"net/http"
	"os"
)

func main() {
	//==================== Hello World with system argument ====================//
	args, stdiName := os.Args[1:], "World"
	if len(args) > 0 {
		stdiName = args[0]
	}

	fmt.Println(hello.BuildHelloMessage(stdiName))

	//=========================== struct and methods ===========================//
	courseInstance := course.Course{
		Name:        "Go 101",
		Description: "Golang introduction course",
		Price:       10000,
	}

	fmt.Println("\n" + courseInstance.GetFullInfo())

	//==================== HTTP endpoints and simple server ====================//
	http.HandleFunc("/", course.CourseJsonHttpHandler(courseInstance))
	http.HandleFunc("/hello", hello.HelloHttpHandler(stdiName))

	go http.ListenAndServe(":5000", nil) // running into a goroutine to unblock the rest of the code

	//========================= channels and goroutines ========================//
	countLimit := 10
	numGoroutines := 3
	countingChannel := counter.SetupParallelCounters(countLimit, numGoroutines)

	for res := range countingChannel {
		fmt.Printf("value %d from counter #%d\n", res[0], res[1])
	}
}
