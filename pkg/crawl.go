package pkg

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type Options struct {
	TimeoutMillisecond int64
	SleepMillisecond   int64
	URL                string
	IgnoreQueries      bool
	Depth              int
	LimitUrls          int
	LimitEmails        int
	WriteToFile        string
}

type Option func(*Options) error

type HTTPChallenge struct {
	browse *browser.Browser

	urls             []string
	Emails           []string
	TotalURLsCrawled int
	TotalURLsFound   int
	options          *Options
}

func NewHTTPChallenge(opts ...Option) *HTTPChallenge {
	opt := &Options{}
	for _, o := range opts {
		err := o(opt)
		if err != nil {
			panic(err)
		}
	}
	b := surf.NewBrowser()
	b.SetUserAgent("Go/email_extractor")
	b.SetTimeout(time.Duration(opt.TimeoutMillisecond) * time.Millisecond)

	return &HTTPChallenge{
		browse:  b,
		options: opt,
	}
}

func (hc *HTTPChallenge) CrawlRecursive(url string, wg *sync.WaitGroup) *HTTPChallenge {
	defer wg.Done()
	urls := hc.Crawl(url)

	var mu sync.Mutex
	for _, u := range urls {
		if len(hc.urls) >= hc.options.LimitUrls {
			break
		}
		if len(hc.Emails) >= hc.options.LimitEmails {
			hc.Emails = hc.Emails[:hc.options.LimitEmails]
			break
		}
		if StringInSlice(u, hc.urls) {
			continue
		}

		mu.Lock()
		hc.urls = append(hc.urls, u)
		mu.Unlock()

		if runtime.NumGoroutine() > 10000 {
			color.Warn.Print("Sleeping")
			color.Secondary.Print("....................")
			color.Secondary.Println(fmt.Sprintf("%ds (goroutines %d, exceeded limit)", 10, runtime.NumGoroutine()))
			time.Sleep(10 * time.Second)
			wg.Add(1)
			go hc.CrawlRecursive(u, wg)
		} else {
			wg.Add(1)
			go hc.CrawlRecursive(u, wg)
		}
	}
	return hc
}

func (hc *HTTPChallenge) Crawl(url string) []string {
	if hc.options.SleepMillisecond > 0 {
		color.Secondary.Print("Sleeping")
		color.Secondary.Print("....................")
		color.Secondary.Println(fmt.Sprintf("%dms (sleeping before request)", hc.options.SleepMillisecond))
		time.Sleep(time.Duration(hc.options.SleepMillisecond) * time.Millisecond)
	}
	urls := []string{}
	err := hc.browse.Open(url)
	if err != nil {
		return urls
	}
	hc.TotalURLsCrawled++

	color.Secondary.Print("Crawling")
	color.Secondary.Print("....................")
	color.Secondary.Println(url)
	rawBody := hc.browse.Body()
	emails := ExtractEmailsFromText(rawBody)
	emails = FilterOutCommonExtensions(emails)
	emails = UniqueStrings(emails)
	if len(emails) > 0 {
		hc.TotalURLsFound++
		color.Note.Print("Emails")
		color.Secondary.Print("......................")
		color.Note.Println(fmt.Sprintf("(%d) %s", len(emails), url))
		for _, email := range emails {
			color.Secondary.Print("                            ")
			color.Success.Println(email)
		}
		fmt.Println()
	}
	if hc.options.WriteToFile != "" {
		hc.Emails = append(hc.Emails, emails...)
		hc.Emails = UniqueStrings(hc.Emails)
	}

	// crawl the page and print all links
	hc.browse.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		href = RelativeToAbsoluteURL(href, url, GetBaseURL(url))

		if hc.options.IgnoreQueries {
			href = RemoveAnyQueryParam(href)
		}
		href = RemoveAnyAnchors(href)
		isSubset := IsSameDomain(hc.options.URL, href)
		if !isSubset {
			return
		}
		if hc.options.Depth != -1 {
			depth := URLDepth(href, hc.options.URL)
			if depth == -1 {
				return
			}
			if depth == 0 {
				return
			}
			if depth > hc.options.Depth {
				return
			}
		}
		urls = append(urls, href)
	})
	urls = UniqueStrings(urls)
	return urls
}
