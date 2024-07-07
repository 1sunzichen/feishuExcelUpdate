package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Hot struct {
	Name  string `selector:"div.goods_info > div.goods_info_text>div.goods_name>a.color_00"`
	Price string `selector:"div.goods_info > div.goods_operate>div.price_promote>span#new"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.wdmcake.cn"),
	)
	var hots = []Hot{}

	c.OnHTML("div.goods", func(e *colly.HTMLElement) {
		hot := &Hot{}
		err := e.Unmarshal(hot)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		hots = append(hots, *hot)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})
	fmt.Printf("%d hots\n", len(hots))

	c.Visit("https://www.wdmcake.cn/category-1.html")
	for _, hot := range hots {

		fmt.Print("Name:", hot.Name)
		fmt.Println("Price:", hot.Price)

		// break
	}
}
