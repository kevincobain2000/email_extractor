# Email Extractor

<p align="center">
<b>With God speed</b> <br> Extract email addresses from entire website, by crawling URLS.
</p>

<p align="center">
<b>Just</b> <code>10 kilo bytes</code> binary, takes <b>just</b> <code>10 seconds</code> to extract emails from 1000 urls.
</p>

A free utility to extract email address by crawling a given url upto a given depth or number of urls to crawl provided by the user.
Web email extractor by url crawl using command line interface. Email addresses can be extracted from any url.
First it extracts all the number of urls provided by the user and at the same time extracts emails using simple Go routines. This simple application allows to crawl through a website with depth options.


**Quick Setup:** One command to install lighweight binary.

**Blazing speeds:** Extracts by crawling URLS in parallel with light weight cpu.

**Crawling capability:** Crawls entire page, finds the links and extracts email addresses.

**Beautiful:** Colorful output with write to file option.

**Dependency Free:** No need to install any dependencies from `pip`, `npm`. Just download and run.

# Install

```sh
curl -sL https://raw.githubusercontent.com/kevincobain2000/email_extractor/master/install.sh | sh
mv email_extractor /usr/local/bin/
```


# Usage

**Simple usage**

```sh
email_extractor -url=kevincobain2000.github.io
```

**Advanced usages**

Cli application

```sh
# Do not crawl urls, just this URL
email_extractor -depth=0 -url=kevincobain2000.github.io

# Crawl everything, including this URL
email_extractor -depth=-1 -url=kevincobain2000.github.io

# write emails to a file
email_extractor -out=emails.txt -url=kevincobain2000.github.io

#extract from 100 urls
email_extractor -limit-urls=100 -url=kevincobain2000.github.io
```

**All Options**

```sh
  -depth int
    	depth of urls to crawl.
    	-1 for url provided & all depths (both backward and forward)
    	0  for url provided (only this)
    	1  for url provided & until first level (forward)
    	2  for url provided & until second level (forward) (default -1)
  -ignore-queries
    	ignore query params in the url
    	Note: pagination links are usually query params
    	Set it to false, if you want to crawl such links
    	 (default true)
  -limit-emails int
    	limit of emails to crawl (default 1000)
  -limit-urls int
    	limit of urls to crawl (default 1000)
  -out string
    	file to write to (default "emails.txt")
  -parallel
    	crawl urls in parallel (default true)
  -sleep int
    	sleep in milliseconds before each request to avoid getting blocked
  -timeout int
    	timeout limit in milliseconds for each request (default 10000)
  -url string
    	url to crawl
  -version
    	prints version
```

# Samples

![Screenshot](https://imgur.com/fVulc7g.png)

# Performance

It crawled `1000 urls`, and found `300 email addresses` in about `11 seconds`.

# Development notes

```sh
go run main.go -h
```

# CHANGE LOG

- v1.0 - Python implementation to extract email addresses by crawling URLS. Installation using pip.
- v2.0 - 100x performance improvement by using goroutines
- v2.5 - 2x performance improvement by not opening the same url again
- v2.6 - Added depth of crawling urls
- v2.8 - Limit emails addresses, and possible fix on relative urls
- v2.10 - Adds hint for status code
- v2.12 - Option to do in parallel and better messaging
- v3.0 - Web UI
- v3.1 - Minor cli change
- v3.4 - Remove Web UI
