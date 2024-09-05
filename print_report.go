package main
import "fmt"
import "sort"

type Page struct {
	URL   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`=============================
  REPORT for %v
=============================`, baseURL)

	sortedpages := sortPages(pages)
	for _, page := range sortedpages {
		url := page.URL
		count := page.count
		fmt.Printf("Found %d internal links to %v\n", count, url)
	}
}


func sortPages(pages map[string]int) []Page {
	sortedpages := []Page{}
	for url, count_val := range pages {
		sortedpages = append(sortedpages, Page{URL: url, count: count_val})
	}
	sort.Slice(sortedpages, func(i, j int) bool {
		if sortedpages[i].count == sortedpages[j].count {
			return sortedpages[i].URL < sortedpages[j].URL
		}
		return sortedpages[i].count > sortedpages[j].count
	})
	return sortedpages
}