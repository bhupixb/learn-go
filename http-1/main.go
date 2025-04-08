// A simple HTTP client and server
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	port    = 9876
	maxLen  = 8
	charset = "abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateRandomChars() string {
	var res string
	for i := 0; i < maxLen; i++ {
		randChar := charset[rand.Intn(len(charset))]
		res += string(randChar)
	}
	return res
}

func httpClientSimple() {
	client := &http.Client{
		Timeout: 10 * time.Second,
		// Timeout: 10 * time.Millisecond,
	}

	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	// one way
	// bodyContent, _ := io.ReadAll(resp.Body)
	// log.Println("resp", resp.StatusCode, string(bodyContent))

	// other way
	var data struct {
		UserId   int    `json:"userId"`
		Id       int    `json:"id"`
		Title    string `json:"title"`
		Complete bool   `json:"complete"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", data)
}

func httpClientWithRequestContext() {
	client := &http.Client{
		Timeout: 10 * time.Second,
		// Timeout: 10 * time.Millisecond,
	}

	url := "https://jsonplaceholder.typicode.com/todos/1"
	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)

	resp, _ := client.Do(request)
	defer resp.Body.Close()
	bodyContent, _ := io.ReadAll(resp.Body)
	log.Println("resp", resp.StatusCode, string(bodyContent))
}

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(1000 * time.Millisecond)
	log.Println(r.Method)
	log.Println(w.Header())
	w.WriteHeader(201)
	log.Println(w.Header())
	w.Write([]byte("Hello World, from server! " + generateRandomChars()))
}

func httpServer() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		Handler:      HelloHandler{},
	}

	server.ListenAndServe()
}

func main() {
	log.Println("starting server at port:", port)

	httpClientSimple()
	httpClientWithRequestContext()

	for i := 0; i < 5; i++ {
		log.Println(generateRandomChars())
	}

	httpServer()
}
