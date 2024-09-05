package main
import "fmt"
import "net/url"
import "golang.org/x/net/html"
import "strings"

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error){
	htmlReader := strings.NewReader(htmlBody)
	parsedHTML, err := html.Parse(htmlReader) // creates a tree of HTML nodes
	if err != nil {return nil, fmt.Errorf("Error parsing HTML: %v", err)}

	var urls []string
	var extractURLs func(*html.Node)
	extractURLs = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href := attr.Val
					parsedURL, err := baseURL.Parse(href)
					if err == nil {urls = append(urls, parsedURL.String())}
					break
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			extractURLs(child)
		}
	}
	extractURLs(parsedHTML)
	return urls, nil
}

