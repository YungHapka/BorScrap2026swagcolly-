package main

import (
	"fmt"
	"log"
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
		DomainGlob: "*",
		Paralelist: 2,
		Deley:      1 * time.Second,
	})

	var products []Product
	c.OnHTML("div.product-card", func(e *colly.HTMLElement) {
		product := Product{}
		product.Name = e.ChildText("h4.title")
		product.Price = e.ChildText("div.price-whapper")
		product.URL = e.Request.AbsoluteURL(e.ChildAttr("a.card-header", "href"))
		products = append(products, product)

	})
	c.OnScraped(func(r *colly.Response) {
		for _, p := range products {
			fmt.Printf("Name: %s, Price: %s, URL:  %s\n", p.Name, p.Price, p.URL)
		}
	})

	err := c.Visit("https://tv.yandex.ru/search?text=%D0%94%D0%B6%D0%BE%D0%BD+%D0%A3%D0%B8%D0%BA")
	if err != nil {
		log.Fatal(err)
	}
}
