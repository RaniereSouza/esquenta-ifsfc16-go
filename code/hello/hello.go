package hello

import (
	"net/http"
)

func BuildHelloMessage(name string) string {
	return "Hello, " + name + "!"
}

func buildHelloInternetMessage(name string) string {
	return "Say hello to the Internet, " + name + "!"
}

func HelloHttpHandler(name string) func(resW http.ResponseWriter, req *http.Request) {
	return func(resW http.ResponseWriter, req *http.Request) {
		resW.Write([]byte(buildHelloInternetMessage(name)))
	}
}
