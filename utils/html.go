package utils

import (
	"ManGo/websites"
	"github.com/gocolly/colly"
	"fmt"
)

func get_images(imagesUrl string, website websites.Website){
	c := colly.NewCollector(
		colly.AllowedDomains(website.Domain),
	)
	c.OnXML("", func(e *colly.XMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	c.Visit(imagesUrl)
}