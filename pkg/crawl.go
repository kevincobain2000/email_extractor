package pkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type Options struct {
	TimeoutMillisecond int64
	Limit              int
}

type Option func(*Options) error

type HTTPChallenge struct {
	browse         *browser.Browser
	alreadyCrawled []string
	options        *Options
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

func (hc *HTTPChallenge) CrawlRecursive(url string) *HTTPChallenge {
	urls := hc.Crawl(url)
	for _, u := range urls {
		if StringInSlice(u, hc.alreadyCrawled) {
			continue
		}
		color.Notice.Println("Crawling ....................", u)
		hc.alreadyCrawled = append(hc.alreadyCrawled, u)
		hc.CrawlRecursive(u)
		if len(hc.alreadyCrawled) > hc.options.Limit {
			break
		}
	}
	return hc
}

func (hc *HTTPChallenge) BrowseAndExtractEmails() *HTTPChallenge {
	for _, u := range hc.alreadyCrawled {
		emails := []string{}
		err := hc.browse.Open(u)
		if err != nil {
			continue
		}

		rawBody := hc.browse.Body()

		emails = append(emails, ExtractEmailsFromText(rawBody)...)
		emails = UniqueStrings(emails)
		if len(emails) == 0 {
			continue
		}
		color.Notice.Println("Emails   ....................", "(", len(emails), ")", u)
		for _, email := range emails {
			color.Success.Println(email)
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
