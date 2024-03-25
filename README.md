# About

Web email extractor by url crawl using command line interface. A free utility to extract email address by crawling a given url upto a given depth or number of urls to crawl provided by the user. Email addresses can be extracted from any url. Emails Extractor can be used to extract emails from a given url. First it extracts all the number of urls by the depth provided by the user and then extracts emails using simple Python libraries. This simple application allows to crawl through a website with a max depth of crawling 5000 urls and extracts email addresses and saves to a file.

# Install


`python setup.py install`_


# Usage


**usage::**

```
email_extractor --url= --limit=
```


```sh
  -limit int
    	limit of urls to crawl (default 20)
  -timeout int
    	timeout in milliseconds (default 1000)
  -url string
    	url to crawl
```