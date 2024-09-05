package main
import "net/url"
import "fmt"
import "strings"

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {return "", fmt.Errorf("Error parsing URL: %v", err)}

	normalPath := parsedURL.Host+parsedURL.Path
	normalPath = strings.TrimSuffix(strings.ToLower(normalPath), "/")
	return normalPath, nil
}