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
	"https://www.dressabelle.com.sg",
}

// BBDressabelleExtender .
type BBDressabelleExtender struct {
	gocrawl.DefaultExtender
}

func (x *BBDressabelleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// Use the goquery document or res.Body to manipulate the data
	fmt.Println(doc.Url)
	return nil, true
}

func (x *BBDressabelleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}

func CrawlDressabelle(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Crawling Dressabelle...")

	opts := gocrawl.NewOptions(new(BBDressabelleExtender))
	opts.CrawlDelay = 100 * time.Millisecond
	crawler := gocrawl.NewCrawlerWithOptions(opts)
	crawler.Run(dressabelleSeeds)
}
