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
	emails  []string
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
		color.Notice.Print("Crawling")
		color.Secondary.Print("....................")
		fmt.Println(u)
		mu.Lock()
		hc.urls = append(hc.urls, u)
		mu.Unlock()
		wg.Add(1)
		go hc.CrawlRecursive(u, wg)
	}
	return hc
}

func (hc *HTTPChallenge) BrowseAndExtractEmails() *HTTPChallenge {
	var wg sync.WaitGroup
	for _, u := range hc.urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			err := hc.browse.Open(url)
			if err != nil {
				return
			}

			rawBody := hc.browse.Body()

			emails := ExtractEmailsFromText(rawBody)
			emails = FilterOutCommonExtensions(emails)
			emails = UniqueStrings(emails)
			if len(emails) == 0 {
				return
			}
			hc.emails = append(hc.emails, emails...)
			color.Notice.Print("Emails")
			color.Secondary.Print("  ....................")
			color.Secondary.Println(fmt.Sprintf("(%d) %s", len(emails), url))
			for _, email := range emails {
				color.Secondary.Print("        ....................")
				color.Success.Println(email)
			}
		}(u)
	}
	wg.Wait()
	hc.emails = UniqueStrings(hc.emails)
	if hc.options.WriteToFile != "" {
		err := WriteToFile(hc.emails, hc.options.WriteToFile)
		if err != nil {
			color.Error.Println("Error writing emails to file", hc.options.WriteToFile)
		} else {
			color.Success.Println("Emails successfully written to", hc.options.WriteToFile)
		}
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
