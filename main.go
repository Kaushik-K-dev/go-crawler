package main
import "fmt"
import "os"
import "strconv"

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Arguments needed - <Website> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error in max Concurrency: %v\n", err)
		return
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error in max Pages: %v\n", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error in configuration: %v", err)
		return
	}
	
	fmt.Printf("Starting crawl of: %s...\n", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	//for page, count := range cfg.pages{fmt.Printf("Page: %v visited %d times\n", page, count)}
	//https://www.wagslane.dev/

	printReport(cfg.pages, rawBaseURL)
}