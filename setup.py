import os
from setuptools import setup

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
    keywords = "http://www.jaist.ac.jp/~s1010205/pybits/pybits.html.LyXconv/pybits.html#x1-20001",
    url = "http://www.jaist.ac.jp/~s1010205/home",
    packages=[''],
    long_description=read('README'),
    classifiers=[
        "Development Status :: 6 - Mature",
        "Topic :: Utilities",
    ],
)
