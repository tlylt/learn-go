package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.github.com",
		"http://www.amazon.com",
		"http://www.twitter.com",
		"http://www.linkedin.com",
		"http://www.youtube.com",
		"http://www.reddit.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Might be up I think"
}