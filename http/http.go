package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		todos = append(todos, t)
		fmt.Printf("% #v\n", todos)

		fmt.Fprintf(w, "hello %s created todos", req.Method)
		return
	}

	if req.Method == "GET" {
		b, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	if req.Method == "PUT" {
		fmt.Println("PUT>>", req.URL.String())
		fmt.Println("PUT-path>>", req.URL.EscapedPath())
		// todos[0].Status = "new status"
		// b, err := json.Marshal(todos)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	fmt.Fprintf(w, "error: %v", err)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.Write(b)
	}
	fmt.Fprintf(w, "hello %s todos", req.Method)
}

// func todosPUTHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "PUT" {
// 		fmt.Println("PUT>>", req.URL.String())
// 		fmt.Println("PUT-path>>", req.URL.EscapedPath())
// 	}
// }

func main() {
	http.HandleFunc("/todos", todosHandler)

	// how to handle xxx/todos/{x} case ??
	// http.HandleFunc("/todos/{1}", todosPUTHandler)
	fmt.Println("starting...")
	log.Fatal(http.ListenAndServe(":1234", nil))
	fmt.Println("end")
}
