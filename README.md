# About

Web email extractor by url crawl using command line interface. A free utility to extract email address by crawling a given url upto a given depth or number of urls to crawl provided by the user. Email addresses can be extracted from any url. Emails Extractor can be used to extract emails from a given url. First it extracts all the number of urls by the depth provided by the user and then extracts emails using simple Python libraries. This simple application allows to crawl through a website with a max depth of crawling 5000 urls and extracts email addresses and saves to a file.

# Install

```sh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/email_extractor/master/install.sh | sh
```


# Usage

```
./email_extractor --url=
```


```sh
  -limit int
    	limit of urls to crawl (default 20)
  -timeout int
    	timeout in milliseconds (default 1000)
  -url string
    	url to crawl
```

# CHANGE LOG

- v1 - Python implementation to extract email addresses by crawling URLS. Installation using pip.
- v2 - 100x performance improvement by using goroutines