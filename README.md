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

**Dependecy Free:** No need to install any dependencies from `pip`, `npm`. Just download and run.

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
mv email_extractor /usr/local/bin/
```


# Usage

**Simple usage**

```sh
email_extractor -url=kevincobain2000.github.io
```

**Advanced usages**


```sh
# Do not crawl urls
email_extractor -depth=0 -url=kevincobain2000.github.io

# write emails to a file
email_extractor -out=marketing.txt -url=kevincobain2000.github.io

#extract from 100 urls
email_extractor -limit-urls=100 -url=kevincobain2000.github.io
```

**All Options**

```sh
  -depth int
    	depth of urls to crawl.
    	-1 for the url provided & all depths (default)
    	0  for the url provided only
    	1  for the url provided & until the first level
    	2  for the url provided & until the second level (default -1)
  -ignore-queries
    	ignore query params in the url
    	Note: pagination links are usually query params
    	Set it to false, if you want to crawl such links
    	 (default true)
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

It crawled `1000 urls`, and found `300 email addresses` in about `11 seconds`.

<p align="center">
  <img alt="bar chart" src='https://instachart.coveritup.app/bar?title=Performance&subtitle=Email+Extractor&output=svg&metric=sec&theme=dark&data={%20%22x%22:%20[%22100%20URLS%22,%20%22500%20URLS%22,%20%221000%20URLS%22],%20%22y%22:%20[[1,6,11]],%20%22names%22:%20[%22Time%20to%20Extract%22]%20}'>
</p>

# CHANGE LOG

- v1.0 - Python implementation to extract email addresses by crawling URLS. Installation using pip.
- v2.0 - 100x performance improvement by using goroutines
- v2.5 - 2x performance improvement by not opening the same url again
- v2.6 - Added depth of crawling urls