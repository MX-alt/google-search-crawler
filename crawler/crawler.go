package crawler

import (
	"encoding/json"
	"net/http"
	"sync"
)

// SearchResult holds the details of a single search result.
type SearchResult struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

// SearchResponse holds the structure of the response from the search API.
type SearchResponse struct {
	Results []SearchResult `json:"results"`
}

// Crawler struct which holds the necessary information for the crawler.
type Crawler struct {
	Client      *http.Client
	Cache       map[string]*SearchResult
	CacheMutex  sync.RWMutex
}

// NewCrawler creates an instance of Crawler.
func NewCrawler() *Crawler {
	return &Crawler{
		Client:     &http.Client{},
		Cache:      make(map[string]*SearchResult),
	}
}

// Search queries the search API with the given query string.
func (c *Crawler) Search(query string) ([]SearchResult, error) {
	if result, found := c.getFromCache(query); found {
		return []SearchResult{*result}, nil
	}
	
	response, err := c.performSearch(query)
	if err != nil {
		return nil, err
	}
	
	c.addToCache(query, response.Results)
	return response.Results, nil
}

// SearchBatch performs searches for multiple queries in parallel.
func (c *Crawler) SearchBatch(queries []string) (map[string][]SearchResult, error) {
	results := make(map[string][]SearchResult)
	var wg sync.WaitGroup

	for _, query := range queries {
		wg.Add(1)
		go func(q string) {
			defer wg.Done()
			result, _ := c.Search(q)
			results[q] = result
		}(query)
	}
	
	wg.Wait()
	return results, nil
}

// performSearch makes the actual API call to the search service.
func (c *Crawler) performSearch(query string) (*SearchResponse, error) {
	// Imagine a real API call here and unmarshal response
	resp, err := c.Client.Get("https://api.example.com/search?q=" + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// getFromCache returns a cached search result if available.
func (c *Crawler) getFromCache(query string) (*SearchResult, bool) {
	c.CacheMutex.RLock()
	defer c.CacheMutex.RUnlock()
	result, found := c.Cache[query]
	return result, found
}

// addToCache adds search results to the cache.
func (c *Crawler) addToCache(query string, results []SearchResult) {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	
	for _, result := range results {
		if _, exists := c.Cache[result.Link]; !exists {
			c.Cache[result.Link] = &result
		}
	}
}