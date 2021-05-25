package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

const version = "1.0.0"

func main() {
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)
	http.HandleFunc("/api/quote", quoteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	fmt.Printf("Starting server at port " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/quote" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	authors, authErr := readLines("data/authors.txt")
	quotes, quoteErr := readLines("data/quotes.txt")

	if authErr == nil && quoteErr == nil {
		randomLine := rand.Intn(len(authors))

		json := "{\"quote\": \"" + quotes[randomLine] + "\", " +
			"\"author\": \"" + authors[randomLine] + "\", " +
			"\"appVersion\": \"" + version + "\"" +
			"}"

		fmt.Fprintf(w, json)
	} else {
		fmt.Fprintf(w, "Error")
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
