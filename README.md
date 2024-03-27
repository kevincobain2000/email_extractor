# Email Extractor

**With God speed - Extract email addresses from entire website, by crawling URLS.**

**Just 10KB** Binary size

Web email extractor by url crawl using command line interface. A free utility to extract email address by crawling a given url upto a given depth or number of urls to crawl provided by the user. Email addresses can be extracted from any url. Emails Extractor can be used to extract emails from a given url.

First it extracts all the number of urls provided by the user and then extracts emails using simple Python libraries. This simple application allows to crawl through a website with a max depth of crawling 5000 urls and extracts email addresses and saves to a file.

# C.I

![go-build-time](https://coveritup.app/badge?org=kevincobain2000&repo=email_extractor&type=go-build-time&branch=master)
![go-test-runtime](https://coveritup.app/badge?org=kevincobain2000&repo=email_extractor&type=go-test-runtime&branch=master)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=email_extractor&type=coverage&branch=master)
![go-binary-size](https://coveritup.app/badge?org=kevincobain2000&repo=email_extractor&type=go-binary-size&branch=master)
![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=email_extractor&type=go-mod-dependencies&branch=master)

![go-build-time](https://coveritup.app/chart?org=kevincobain2000&repo=email_extractor&type=go-build-time&output=svg&width=160&height=160&branch=master&line=fill)
![go-test-runtime](https://coveritup.app/chart?org=kevincobain2000&repo=email_extractor&type=go-test-runtime&output=svg&width=160&height=160&branch=master&line=fill)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=email_extractor&type=coverage&output=svg&width=160&height=160&branch=master)
![go-binary-size](https://coveritup.app/chart?org=kevincobain2000&repo=email_extractor&type=go-binary-size&output=svg&width=160&height=160&branch=master)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=email_extractor&type=go-mod-dependencies&output=svg&width=160&height=160&branch=master&line=fill)


# Install

```sh
curl -sL https://raw.githubusercontent.com/kevincobain2000/email_extractor/master/install.sh | sh
mv email_extractor /usr/local/bin
```


# Usage

Simple usage

```sh
email_extractor -url=kevincobain2000.github.io
```

Alternative usages


```sh
# Do not crawl urls
email_extractor -depth=0 -url=kevincobain2000.github.io

# write emails to a file
email_extractor -out=marketing.txt -url=kevincobain2000.github.io

#extract from 100 urls
email_extractor -limit-urls=100 -url=kevincobain2000.github.io
```

Options

```sh
  -depth int
    	depth of urls to crawl.
    	-1 for infinite depth
    	0 for no depth, only the url provided
    	1 for only the url provided and links from the url provided until the first level
    	2 for only the url provided and links from the url provided until the second level (default -1)
  -ignore-queries
    	ignore query params in the url (default true)
  -limit-urls int
    	limit of urls to crawl (default 1000)
  -out string
    	file to write to (default "emails.txt")
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

![Screenshot](https://imgur.com/efJfwB3.png)

# Performance

It crawled `1000 urls`, and found `300 email addresses` in about `10 seconds`.

```sh
╰─$ time email_extractor -url=https://medium.com
    6.24s user 2.42s system 72% cpu 11.938 total

╰─$ wc -l emails.txt
    314 emails.txt
```

# CHANGE LOG

- v1.0 - Python implementation to extract email addresses by crawling URLS. Installation using pip.
- v2.0 - 100x performance improvement by using goroutines
- v2.5 - 2x performance improvement by not opening the same url again
- v2.6 - Added depth of crawling urls