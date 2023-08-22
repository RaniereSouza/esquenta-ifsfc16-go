package course

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Name        string `json:"course"`
	Description string `json:"description"`
	Price       int    `json:"price"` // cents
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf(
		"Course:\n"+
			"  Name: %s, Description: %s, Price: %d",
		c.Name, c.Description, c.Price,
	)
}

func CourseJsonHttpHandler(course Course) func(resW http.ResponseWriter, req *http.Request) {
	return func(resW http.ResponseWriter, req *http.Request) {
		json.NewEncoder(resW).Encode(course)
	}
}
