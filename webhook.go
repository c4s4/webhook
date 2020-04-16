package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var Conf Configuration

type Call struct {
	Repository struct {
		Name string `json:"full_name"`
	}
}

func Run(action Action) {
	
}

func handle(response http.ResponseWriter, request *http.Request) {
	event := request.Header.Get("X-GitHub-Event")
	var call Call
	err := json.NewDecoder(request.Body).Decode(&call)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	name := call.Repository.Name
	if repository, ok := Conf[name]; ok {
		if action, ok := action[event]; ok {
			Run(action)
		}
	} else {
		http.Error(response, "Repository "+name+" not found", http.StatusNotFound)
		return
	}
	fmt.Printf("> %s on %s\n", event, repository)
}

func setupHandler() *http.ServeMux {
	handler := http.NewServeMux()
	handler.HandleFunc("/", handle)
	return handler
}

func main() {
	if len(os.Args) != 2 {
		println("You must pass configuration file on command line")
		os.Exit(1)
	}
	configuration, err := LoadConfiguration(os.Args[1])
	if err != nil {
		println("Error loading configuration:", err.Error())
		os.Exit(2)
	}
	Conf = *configuration
	handler := setupHandler()
	http.ListenAndServe(":8001", handler)
}
