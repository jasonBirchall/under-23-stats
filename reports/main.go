package reports

import (
	"time"

	"github.com/gocolly/colly"
)

type player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (player *player) setPlayerAge() {
	c := colly.NewCollector(
		colly.AllowedDomains("footystats.org"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
		Delay:       5 * time.Second,
	})

	// Find and visit all links on the parent page.
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

}
