package pkg

import (
	"net/url"
	"regexp"
	"strings"
)

func ByteSliceToString(b []byte) string {
	return string(b)
}
func StringToByteSlice(s string) []byte {
	return []byte(s)
}

func IsSameDomain(u1, u2 string) bool {
	parsedURL1, err := url.Parse(u1)
	if err != nil {
		return false
	}

	parsedURL2, err := url.Parse(u2)
	if err != nil {
		return false
	}

	return parsedURL1.Host == parsedURL2.Host
}

func URLDepth(u, referenceURL string) int {
	refURL, err := url.Parse(referenceURL)
	if err != nil {
		return -1 // Error parsing reference URL
	}

	parsedURL, err := url.Parse(u)
	if err != nil {
		return -1 // Error parsing URL
	}

	refPath := strings.TrimSuffix(refURL.Path, "/")
	parsedPath := strings.TrimSuffix(parsedURL.Path, "/")

	if !strings.HasPrefix(parsedPath, refPath) {
		return 0 // No common path prefix, depth is 0
	}

	relPath := strings.TrimPrefix(parsedPath, refPath)
	depth := strings.Count(relPath, "/")

	if relPath == "" {
		return 0 // URL is the same as the reference URL, depth is 0
	}

	// Count the number of segments, including root ("/")
	// depth++
	return depth
}

func RemoveAnyQueryParam(u string) string {
	if strings.Contains(u, "?") {
		return strings.Split(u, "?")[0]
	}
	return u
}

func RemoveAnyAnchors(u string) string {
	if strings.Contains(u, "#") {
		return strings.Split(u, "#")[0]
	}
	return u
}

func GetBaseURL(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return ""
	}
	return parsedURL.Scheme + "://" + parsedURL.Host
}

func ExtractEmailsFromText(text string) []string {
	// Regular expression to match email addresses
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all email addresses in the text
	emails := re.FindAllString(text, -1)

	return emails
}

func RelativeToAbsoluteURL(href, currentURL, baseURL string) string {
	if strings.HasPrefix(href, "http") {
		return href
	}

	if strings.HasPrefix(href, "/") {
		return baseURL + href
	}

	return currentURL + href
}
