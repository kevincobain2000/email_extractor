package main

import (
	"flag"
	"fmt"

	"github.com/gookit/color"
	"github.com/kevincobain2000/email_extractor/pkg"
)

type Flags struct {
	url     string
	limit   int
	timeout int64
}

var f Flags

func main() {
	SetupFlags()
	options := []pkg.Option{
		func(opt *pkg.Options) error {
			opt.TimeoutMillisecond = f.timeout
			opt.Limit = f.limit
			return nil
		},
	}
	hc := pkg.NewHTTPChallenge(options...)
	hc.CrawlRecursive(f.url)
	color.Success.Print("Crawling")
	color.Secondary.Print("....................")
	color.Success.Println("Complete")
	fmt.Println()
	hc.BrowseAndExtractEmails()
}

func SetupFlags() {
	flag.StringVar(&f.url, "url", "", "url to crawl")
	flag.IntVar(&f.limit, "limit", 100, "limit of urls to crawl")
	flag.Int64Var(&f.timeout, "timeout", 10000, "timeout in milliseconds")
	flag.Parse()
}
