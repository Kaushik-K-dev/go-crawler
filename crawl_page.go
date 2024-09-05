package main
import "fmt"
import "net/url"

func (cfg *config) crawlPage(rawCurrentURL string){
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.pagesLen() >= cfg.maxPages {return}

	// baseURL, err := url.Parse(rawBaseURL)
	// if err != nil {
	// 	fmt.Printf("Error parsing URL: %v\n", err)
	// 	return
	// }
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing current URL: %v\n", err)
		return	
	}
	if cfg.baseURL.Host != currentURL.Host {return}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {fmt.Printf("Error normalizing URL: %v\n", rawCurrentURL)}

	exists := cfg.addPageVisit(normalizedURL); if exists == false {return}

	fmt.Printf("Currently crawling: %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting HTML: %v\n", err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error getting URLs from %s: %v\n", normalizedURL, err)
	}

	for _, nextURL := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
