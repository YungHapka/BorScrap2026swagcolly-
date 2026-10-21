package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

type Product struct {
	Name  string
	Price string
	URL   string
}

func main() {

	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response from", r.Request.URL.String())
	})

	urls := []string{
		"https://tv.yandex.ru/search?text=%D0%94%D0%B6%D0%BE%D0%BD+%D0%A3%D0%B8%D0%BA/1",
		"https://tv.yandex.ru/search?text=%D0%94%D0%B6%D0%BE%D0%BD+%D0%A3%D0%B8%D0%BA/2",
		"https://tv.yandex.ru/search?text=%D0%94%D0%B6%D0%BE%D0%BD+%D0%A3%D0%B8%D0%BA/3",
		"https://tv.yandex.ru/search?text=%D0%94%D0%B6%D0%BE%D0%BD+%D0%A3%D0%B8%D0%BA/4",
	}

	for _, url := range urls {
		c.Visit(url)
	}
	c.Wait()
}
