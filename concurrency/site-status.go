package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

func SiteMain() {
	sites := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.amazon.com",
		"http://www.golang.org",
	}

	c := make(chan string)
	for _, link := range sites {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, " is up :)")
	c <- link
}
