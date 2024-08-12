package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

/*
Exercise: Concurrent Web Scraper

Write a program in Go that concurrently fetches the HTML content of a list of URLs and counts the number of words in each page.
The program should print the URL, the word count, and any errors encountered during the fetch.

Requirements:

1. Define a function fetch(url string, ch chan<- string) that fetches the HTML content of a URL, counts the number of words, and sends the result on a channel.
   If there's an error, it should also send an error message on the channel.

2. Define a function wordCount(url string) that counts the number of words in the HTML content.

3. In the main() function, create a list of URLs and a channel to communicate results.
   Start a goroutine for each URL to concurrently fetch and process the content.

4. Print the URL, word count, and any errors received from the channel.

Note: You can use the http.Get function to fetch the HTML content.
To count words, you can use a simple algorithm that splits the content into words based on spaces.

This exercise will help you practice creating concurrent programs in Go using goroutines and channels, as well as error handling.
It also involves working with HTTP requests and string manipulation.

*/

func Fetcher() {

	pipe := make(chan string)

	urls := []string{
		"https://pkg.go.dev/sync/atomic",
		"https://devwizard.me",
	}

	for _, url := range urls {
		go fetch(url, pipe)
		defer close(pipe)
	}

	// Receive and print messages from the channel
	for msg := range pipe {
		fmt.Println(msg)
	}

	// Close the channel after all goroutines are done
	close(pipe)

}

func fetch(url string, ch chan<- string) {
	count, err := wordCount(url)
	ch <- fmt.Sprintf("Words counted %d - Error: %v", count, err)
}

func wordCount(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, err
	}

	var count int
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.TextNode {
			text := strings.TrimSpace(n.Data)
			words := strings.Fields(text)
			count += len(words)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)

	return count, nil
}
