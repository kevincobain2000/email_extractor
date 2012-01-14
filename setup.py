import os
from setuptools import setup

# Utility function to read the README file.
# Used for the long_description.  It's nice, because now 1) we have a top level
# README file and 2) it's easier to type in the README file than to put a raw
# string in below ...
def read(fname):
    return open(os.path.join(os.path.dirname(__file__), fname)).read()

setup(
    scripts = ['scripts/email_extractor'],
    name = "email_extractor",
    version = "v1",
    author = "KATHURIA Pulkit",
    author_email = "kevincobain2000@gmail.com",
    description = ("Extract Emails by using webcrawl "
                                   "Extract Url links from a site"),
    license = "BSD",
    keywords = "http://www.jaist.ac.jp/~s1010205/email_extractor/html/index.html",
    url = "http://www.jaist.ac.jp/~s1010205/email_extractor/html/index.html",
    packages=[''],
    long_description=read('README'),
    classifiers=[
        "Development Status :: 6 - Mature",
        "Topic :: Utilities",
    ],
)
