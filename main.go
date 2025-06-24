package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)
type createRequest struct {
	URL string `json:"url"`
}
type createResponse struct {
	URL string `json:"short_code"`
}

var urlMap = make(map[string]string)
var mutex = &sync.Mutex{}

func homePageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello! Welcome to my API.\n")
}

func validateUrl(shortCode, longURL string) {
	_ , err := http.Get(longURL)
	if err != nil {
		mutex.Lock()
		delete(urlMap, shortCode)
		mutex.Unlock()
		fmt.Println("Invalid Url")
		return
	}
}

func createLink(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Failed with error:", http.StatusMethodNotAllowed)
		
		return
	}
	var reqData createRequest
	decoder := json.NewDecoder(req.Body)
	
	err := decoder.Decode(&reqData)
	if err != nil {
		http.Error(w, "Bad request: could not decode JSON", http.StatusBadRequest)
		
		return
	}
	
	fmt.Printf("Successfully decoded the URL: %s\n", reqData.URL)
	shortCode := String(6)
	mutex.Lock()
	urlMap[shortCode] = reqData.URL
	mutex.Unlock()

	go validateUrl(shortCode, reqData.URL)
	
	responseData := &createResponse{
		URL: shortCode,
	}

	w.Header().Set(
		"Content-Type",
		"application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(responseData)

	
}

func redirectUser(w http.ResponseWriter, req *http.Request) {
	var urlCode string = strings.TrimPrefix(req.URL.Path, "/r/")
	mutex.Lock()
	value, ok := urlMap[urlCode]
	mutex.Unlock()
	if !ok {
		http.Error(w, "URL not found.", http.StatusNotFound)
		return
	}

	http.Redirect(w, req, value, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/hello", homePageHandler)
	http.HandleFunc("/create", createLink)
	http.HandleFunc("/r/", redirectUser)
	http.ListenAndServe(":8080", nil)
}