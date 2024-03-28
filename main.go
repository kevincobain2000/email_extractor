package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

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
	limitEmails   int
	depth         int
	timeout       int64
	sleep         int64
}

var f Flags

func main() {
	startTime := time.Now()

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
			opt.LimitEmails = f.limitEmails
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
	fmt.Println()
	color.Warn.Print("Crawling")
	color.Secondary.Print("....................")
	color.Note.Println("Complete!")
	color.Warn.Print("URLs")
	color.Secondary.Print("........................")
	ratio := (float64(hc.TotalURLsFound) / float64(hc.TotalURLsCrawled)) * 100
	fmt.Printf("%d urls crawled, %d urls with emails (%.2f﹪ hit rate)\n", hc.TotalURLsCrawled, hc.TotalURLsFound, ratio)

	hc.Emails = pkg.UniqueStrings(hc.Emails)

	color.Warn.Print("Unique emails")
	color.Secondary.Print("...............")
	fmt.Printf("%d addresses\n", len(hc.Emails))

	if f.writeToFile != "" {
		err := pkg.WriteToFile(hc.Emails, f.writeToFile)
		if err != nil {
			color.Danger.Print("Output file")
			color.Secondary.Print("・・・・・・・・")
			color.Danger.Println("Error writing emails to file", f.writeToFile)
		} else {
			color.Warn.Print("Output file")
			color.Secondary.Print(".................")
			color.Note.Println(f.writeToFile)
		}
	}
	endTime := time.Now()
	color.Warn.Print("Time taken")
	color.Secondary.Print("..................")
	fmt.Println(endTime.Sub(startTime))
}

func SetupFlags() {
	flag.StringVar(&f.url, "url", "", "url to crawl")
	flag.StringVar(&f.writeToFile, "out", "emails.txt", "file to write to")

	flag.IntVar(&f.limitUrls, "limit-urls", 1000, "limit of urls to crawl")
	flag.IntVar(&f.limitEmails, "limit-emails", 1000, "limit of emails to crawl")

	flag.IntVar(&f.depth, "depth", -1, `depth of urls to crawl.
-1 for url provided & all depths (both backward and forward)
0  for url provided (only this)
1  for url provided & until first level (forward)
2  for url provided & until second level (forward)`)

	flag.Int64Var(&f.timeout, "timeout", 10000, "timeout limit in milliseconds for each request")
	flag.Int64Var(&f.sleep, "sleep", 0, "sleep in milliseconds before each request to avoid getting blocked")

	flag.BoolVar(&f.version, "version", false, "prints version")
	flag.BoolVar(&f.ignoreQueries, "ignore-queries", true, `ignore query params in the url
Note: pagination links are usually query params
Set it to false, if you want to crawl such links
`)
	flag.Parse()

	if !strings.HasPrefix(f.url, "http") {
		f.url = "https://" + f.url
	}
}
