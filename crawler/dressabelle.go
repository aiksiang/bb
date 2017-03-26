package crawler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

var dressabelleSeeds = []string{
	// "https://www.dressabelle.com.sg",
	"https://www.dressabelle.com.sg/dsb/new-arrivals/asymmetrical-top-navy-detail.html",
}

var dressabelleBlacklists = []string{
	"https://www.dressabelle.com.sg/create-an-account.html",
	"https://www.dressabelle.com.sg/userarea.html",
	"https://www.dressabelle.com.sg/cart.html",
}

// BBDressabelleExtender .
type BBDressabelleExtender struct {
	gocrawl.DefaultExtender
}

// Visit .
func (x *BBDressabelleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// Use the goquery document or res.Body to manipulate the data
	fmt.Println(doc.Url)
	name := parseName(doc)
	fmt.Println(name)
	price := parsePrice(doc)
	fmt.Println(price)

	return nil, true
}

// Filter .
func (x *BBDressabelleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}

// CrawlDressabelle .
func CrawlDressabelle(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Crawling Dressabelle...")

	opts := gocrawl.NewOptions(new(BBDressabelleExtender))
	opts.CrawlDelay = 100 * time.Millisecond
	crawler := gocrawl.NewCrawlerWithOptions(opts)
	crawler.Run(dressabelleSeeds)
}

func parseName(doc *goquery.Document) string {
	s := doc.Find("h1").FilterFunction(
		func(_ int, s *goquery.Selection) bool {
			propName, exists := s.Attr("itemprop")
			return exists && propName == "name"
		})
	return s.Text()
}

func parsePrice(doc *goquery.Document) string {
	s := doc.Find("meta").FilterFunction(
		func(_ int, s *goquery.Selection) bool {
			propName, exists := s.Attr("itemprop")
			return exists && propName == "price"
		})
	price, exists := s.Attr("content")
	if !exists {
		return ""
	}
	return price
}
