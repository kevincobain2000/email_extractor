package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"github.com/gookit/color"
	"github.com/kevincobain2000/email_extractor/pkg"
)

var version = "dev"

type Flags struct {
	version     bool
	crawl       bool
	url         string
	writeToFile string
	limit       int
	timeout     int64
}

var f Flags

func main() {
	SetupFlags()
	if f.version {
		fmt.Println(version)
		return
	}
	options := []pkg.Option{
		func(opt *pkg.Options) error {
			opt.TimeoutMillisecond = f.timeout
			opt.Limit = f.limit
			opt.Crawl = f.crawl
			opt.WriteToFile = f.writeToFile
			return nil
		},
	}

	var wgC sync.WaitGroup
	wgC.Add(1)
	hc := pkg.NewHTTPChallenge(options...)
	hc.CrawlRecursive(f.url, &wgC)
	wgC.Wait()
	color.Success.Print("Crawling")
	color.Secondary.Print("....................")
	color.Success.Println("Complete")
	fmt.Println()

	hc.BrowseAndExtractEmails()
}

func SetupFlags() {
	flag.BoolVar(&f.version, "version", false, "prints version")
	flag.BoolVar(&f.crawl, "crawl", true, "crawl urls")
	flag.StringVar(&f.url, "url", "", "url to crawl")
	flag.StringVar(&f.writeToFile, "o", "emails.txt", "file to write to")
	flag.IntVar(&f.limit, "limit", 1000, "limit of urls to crawl")
	flag.Int64Var(&f.timeout, "timeout", 10000, "timeout in milliseconds")
	flag.Parse()

	if !strings.HasPrefix(f.url, "http") {
		f.url = "https://" + f.url
	}
}
