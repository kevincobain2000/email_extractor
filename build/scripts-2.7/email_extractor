# -*- coding: utf-8 -*-
#! /usr/bin/env python
__thisfile__ = "http://www.jaist.ac.jp/~s1010205/email_extractor/email_extractor.py"
"""
    Web Data Extractor, extract emails by sitecrawl
    Copyright (C) 2011 KATHURIA Pulkit
    Contact: pulkit@jaist.ac.jp

    Contributors:
        Open Source Sitemap Generator sitemap_gen by Vladimir Toncar
        http://toncar.cz/opensource/sitemap_gen.html

    This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation; either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
"""
import sys
import re
import commands
from urllib import urlopen
from collections import defaultdict
import argparse
import string
import urllib2
import urlparse
from HTMLParser import HTMLParser
from HTMLParser import HTMLParseError
import robotparser
import httplib

def getPage(url):
    try:
        f = urllib2.urlopen(url)
        page = ""
        for i in f.readlines():
            page += i
        date = f.info().getdate('Last-Modified')
        if date == None:
            date = (0, 0, 0)
        else:
            date = date[:3]
        f.close()
        return (page, date, f.url)
    except urllib2.URLError, detail:
        pass
        return (None, (0,0,0), "")

def joinUrls(baseUrl, newUrl):
	helpUrl, fragment = urlparse.urldefrag(newUrl)
        return urlparse.urljoin(baseUrl, helpUrl)

def getRobotParser(startUrl):
	rp = robotparser.RobotFileParser()
	robotUrl = urlparse.urljoin(startUrl, "/robots.txt")
	page, date, url = getPage(robotUrl)

	if page == None:
	    return None
	rp.parse(page)
	return rp

class MyHTMLParser(HTMLParser):
    def __init__(self, pageMap, redirects, baseUrl, maxUrls, blockExtensions, robotParser):
        HTMLParser.__init__(self)
        self.pageMap = pageMap
	self.redirects = redirects
        self.baseUrl = baseUrl
        self.server = urlparse.urlsplit(baseUrl)[1] # netloc in python 2.5
        self.maxUrls = maxUrls
        self.blockExtensions = blockExtensions
	self.robotParser = robotParser
    def hasBlockedExtension(self, url):
        p = urlparse.urlparse(url)
        path = p[2].upper() # path attribute
        for i in self.blockExtensions:
            if path.endswith(i):
                return 1
        return 0
    def handle_starttag(self, tag, attrs):
        if len(self.pageMap) >= self.maxUrls:
            return
        if (tag.upper() == "BASE"):
	    if (attrs[0][0].upper() == "HREF"):
		self.baseUrl = joinUrls(self.baseUrl, attrs[0][1])
        if (tag.upper() == "A"):
            url = ""
	    for attr in attrs:
                if (attr[0].upper() == "REL") and (attr[1].upper().find('NOFOLLOW') != -1):
                    return  
                elif (attr[0].upper() == "HREF") and (attr[1].upper().find('MAILTO:') == -1):
                    url = joinUrls(self.baseUrl, attr[1])
            if url == "": return
            if urlparse.urlsplit(url)[1] <> self.server:
                return
            if self.hasBlockedExtension(url) or self.redirects.count(url) > 0:
                return
            if (self.robotParser <> None) and not(self.robotParser.can_fetch("*", url)):
                return
            if not(self.pageMap.has_key(url)):
                self.pageMap[url] = ()

def getUrlToProcess(pageMap):
    for i in pageMap.keys():
        if pageMap[i] == ():
            return i
    return None

def parsePages(startUrl, maxUrls, blockExtensions):
    pageMap = {}
    pageMap[startUrl] = ()
    redirects = []
    robotParser = getRobotParser(startUrl)
    while True:
        url = getUrlToProcess(pageMap)
        if url == None:
            break
        print " ", url
        page, date, newUrl = getPage(url)
        if page == None:
            del pageMap[url]
	elif url != newUrl:
	    print newUrl
            del pageMap[url]
	    pageMap[newUrl] = ()
	    redirects.append(url)
        else:
            pageMap[url] = date
            parser = MyHTMLParser(pageMap, redirects, url, maxUrls, blockExtensions, robotParser)
            try:
                parser.feed(page)
                parser.close()
            except HTMLParseError:
                pass
            except UnicodeDecodeError:
                pass
    return pageMap

def grab_email(text):
    found = []
    mailsrch = re.compile(r'[\w\-][\w\-\.]+@[\w\-][\w\-\.]+[a-zA-Z]{1,4}')
    for line in text:
        found.extend(mailsrch.findall(line))    
    u = {}
    for item in found:
        u[item] = 1
    return u.keys()

def urltext(url):
    viewsource = urlopen(url).readlines()
    return viewsource

def crawl_site(url, limit):
    return parsePages(url, limit, 'None')

if __name__ == '__main__':
    parser = argparse.ArgumentParser(add_help = True)
    parser = argparse.ArgumentParser(description= 'Web Email Extractor')
    parser.add_argument('-l','--limit', action="store", default=100, dest= "limit", type= int, help='-l numUrlsToCrawl')
    parser.add_argument('-u','--url', action="store" ,dest= "url", help='-u http://sitename.com')
    myarguments = parser.parse_args()
    emails = defaultdict(int)
    for url in crawl_site(myarguments.url, myarguments.limit):
        for email in grab_email(urltext(url)):
            if not emails.has_key(email): print email
            emails[email] += 1

        




