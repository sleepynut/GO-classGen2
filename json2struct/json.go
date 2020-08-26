package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Birthday time.Time

func (bd *Birthday) UnmarshalJSON(b []byte) error {
	fmt.Println("Here")
	s, err := time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}

	*bd = Birthday(s)
	return nil
}
func (bd Birthday) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd)
}

type Todo struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Status string   `json:"status"`
	Time   Birthday `json:"time"`
}

func main() {
	data := []byte(`{
				"id": 1,
				"title": "pay credit card",
				"status": "active",
				"time":"2020-08-26"
			}`)
	var t Todo
	err := json.Unmarshal(data, &t)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", t.Time)
	fmt.Println("=================")
	{
		s, _ := time.Parse("2006-01-02", "2020-08-26")
		fmt.Println("HERE>", s)
	}

}
