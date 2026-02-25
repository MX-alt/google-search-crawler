package main

import (
    "fmt"
    "time"
)

// Search function for a single query
func search(query string) {
    fmt.Printf("Searching for: %s\n", query)
    // Simulate search process
    time.Sleep(1 * time.Second)
    fmt.Println("Search results for:", query)
}

// Batch search function
func batchSearch(queries []string) {
    for _, query := range queries {
        search(query)
    }
}

// Cache struct to store search results
type Cache struct {
    results map[string]string
}

// NewCache function to create a new cache
func NewCache() *Cache {
    return &Cache{results: make(map[string]string)}
}

// Cache search results
func (c *Cache) cacheSearch(query string, result string) {
    c.results[query] = result
}

// Get cached result
func (c *Cache) getCache(query string) (string, bool) {
    result, found := c.results[query]
    return result, found
}

func main() {
    // Example usage
    fmt.Println("--- Single Search ---")
    search("golang tutorials")

    fmt.Println("--- Batch Search ---")
    batchQueries := []string{"golang tutorials", "python tutorials", "java tutorials"}
    batchSearch(batchQueries)

    fmt.Println("--- Cache Functionality ---")
    cache := NewCache()
    cache.cacheSearch("golang tutorials", "Result for golang tutorials")
    if result, found := cache.getCache("golang tutorials"); found {
        fmt.Println("Cached Result:", result)
    } else {
        fmt.Println("No cache found for the query")
    }
}