package main

import (
	"flag"
	"fmt"

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
	fmt.Println("")
	hc.BrowseAndExtractEmails()
}

func SetupFlags() {
	flag.StringVar(&f.url, "url", "", "url to crawl")
	flag.IntVar(&f.limit, "limit", 20, "limit of urls to crawl")
	flag.Int64Var(&f.timeout, "timeout", 1000, "timeout in milliseconds")
	flag.Parse()
}
