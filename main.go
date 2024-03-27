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
	version       bool
	ignoreQueries bool
	url           string
	writeToFile   string
	limitUrls     int
	depth         int
	timeout       int64
	sleep         int64
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
			opt.SleepMillisecond = f.sleep
			opt.LimitUrls = f.limitUrls
			opt.WriteToFile = f.writeToFile
			opt.URL = f.url
			opt.Depth = f.depth
			opt.IgnoreQueries = f.ignoreQueries
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
	color.Success.Println("Complete!")

	hc.Emails = pkg.UniqueStrings(hc.Emails)
	if f.writeToFile != "" {
		err := pkg.WriteToFile(hc.Emails, f.writeToFile)
		if err != nil {
			color.Danger.Print("Emails")
			color.Secondary.Print("  ....................")
			color.Danger.Println("Error writing emails to file", f.writeToFile)
		} else {
			color.Success.Print("Emails")
			color.Secondary.Print("  ....................")
			color.Success.Println("Success writing emails to file", f.writeToFile)
		}
	}
}

func SetupFlags() {
	flag.StringVar(&f.url, "url", "", "url to crawl")
	flag.StringVar(&f.writeToFile, "out", "emails.txt", "file to write to")

	flag.IntVar(&f.limitUrls, "limit-urls", 1000, "limit of urls to crawl")

	flag.IntVar(&f.depth, "depth", -1, `depth of urls to crawl.
-1 for infinite depth
0 for no depth, only the url provided
1 for only the url provided and links from the url provided until the first level
2 for only the url provided and links from the url provided until the second level`)

	flag.Int64Var(&f.timeout, "timeout", 10000, "timeout limit in milliseconds for each request")
	flag.Int64Var(&f.sleep, "sleep", 0, "sleep in milliseconds before each request to avoid getting blocked")

	flag.BoolVar(&f.version, "version", false, "prints version")
	flag.BoolVar(&f.ignoreQueries, "ignore-queries", true, "ignore query params in the url")
	flag.Parse()

	if !strings.HasPrefix(f.url, "http") {
		f.url = "https://" + f.url
	}
}
