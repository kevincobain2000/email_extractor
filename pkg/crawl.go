package pkg

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type Options struct {
	TimeoutMillisecond int64
	Limit              int
	Crawl              bool
	WriteToFile        string
}

type Option func(*Options) error

type HTTPChallenge struct {
	browse *browser.Browser

	urls    []string
	Emails  []string
	options *Options
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
	if !hc.options.Crawl {
		urls = []string{url}
	}
	var mu sync.Mutex
	for _, u := range urls {
		if len(hc.urls) >= hc.options.Limit {
			break
		}
		if StringInSlice(u, hc.urls) {
			continue
		}

		mu.Lock()
		hc.urls = append(hc.urls, u)
		mu.Unlock()
		wg.Add(1)
		go hc.CrawlRecursive(u, wg)
	}
	return hc
}

func (hc *HTTPChallenge) Crawl(url string) []string {
	urls := []string{}
	urls = append(urls, url)
	err := hc.browse.Open(url)
	if err != nil {
		return urls
	}
	color.Secondary.Print("Crawling")
	color.Secondary.Print("....................")
	color.Secondary.Println(url)
	rawBody := hc.browse.Body()
	emails := ExtractEmailsFromText(rawBody)
	emails = FilterOutCommonExtensions(emails)
	emails = UniqueStrings(emails)
	if len(emails) > 0 {
		color.Note.Print("Emails")
		color.Secondary.Print("  ....................")
		color.Note.Println(fmt.Sprintf("(%d) %s", len(emails), url))
		for _, email := range emails {
			// color.Secondary.Print("        ....................")
			color.Secondary.Print("                            ")
			color.Success.Println(email)
		}
		fmt.Println()
	}
	if hc.options.WriteToFile != "" {
		hc.Emails = append(hc.Emails, emails...)
	}

	// crawl the page and print all links
	hc.browse.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		href = hc.relativeToAbsoluteURL(href)

		href = RemoveAnyQueryParam(href)
		href = RemoveAnyAnchors(href)
		isSubset := IsSameDomain(url, href)
		if isSubset {
			urls = append(urls, href)
		}
	})
	urls = UniqueStrings(urls)
	return urls
}

func (hc *HTTPChallenge) relativeToAbsoluteURL(href string) string {
	if !strings.HasPrefix(href, "http") && !strings.HasPrefix(href, "//") {
		href = fmt.Sprintf("%s://%s%s", hc.browse.Url().Scheme, hc.browse.Url().Host, href)
	}
	return href
}
