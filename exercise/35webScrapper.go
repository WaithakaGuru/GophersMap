/* Exercise Thirty Five
35. Write a concurrent web scraper that fetches titles from a list of URLs.

KEY CONCEPTS:
- Goroutines: Run functions concurrently (concurrent != parallel)
- Channels: Communicate between goroutines safely
- sync.WaitGroup: Wait for multiple goroutines to complete
- Context: Handle timeouts and cancellation
- Error handling in concurrent contexts
- HTTP requests and HTML parsing
- Efficient resource pooling and cleanup
*/

package exercise

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ============================================================================
// DATA STRUCTURES
// ============================================================================

// ScrapedData holds the result of scraping a URL
type ScrapedData struct {
	URL        string        // Original URL
	Title      string        // Extracted title
	StatusCode int           // HTTP status code
	FetchTime  time.Duration // Time taken to fetch
	Error      error         // Any error that occurred
}

// ScraperConfig holds configuration for the scraper
type ScraperConfig struct {
	MaxConcurrency int           // Maximum concurrent requests
	Timeout        time.Duration // Timeout for each request
	RetryCount     int           // Number of retries on failure
}

// ============================================================================
// BASIC WEB SCRAPER (Sequential)
// ============================================================================
// This shows the problem: slow sequential execution

func basicSequentialScraper(urls []string) []ScrapedData {
	fmt.Println("\n========== SEQUENTIAL SCRAPER ==========")
	fmt.Println("Fetching URLs one at a time (SLOW)...")

	var results []ScrapedData
	startTime := time.Now()

	for i, url := range urls {
		start := time.Now()

		// Fetch the URL
		data := fetchURL(url)
		data.FetchTime = time.Since(start)

		results = append(results, data)
		fmt.Printf("[%d/%d] %s - %s (%.2fs)\n",
			i+1, len(urls), url, data.Title, data.FetchTime.Seconds())
	}

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal time: %.2f seconds\n", totalTime.Seconds())
	return results
}

// ============================================================================
// SIMPLE CONCURRENT SCRAPER (Using Goroutines + Channels)
// ============================================================================
// This demonstrates basic concurrency

func simpleConcurrentScraper(urls []string) []ScrapedData {
	fmt.Println("\n========== SIMPLE CONCURRENT SCRAPER ==========")
	fmt.Println("Fetching URLs concurrently (FAST)...")

	// Create a channel to receive results
	// Buffered channel with capacity for all URLs
	resultsChan := make(chan ScrapedData, len(urls))

	startTime := time.Now()

	// Launch a goroutine for each URL
	for _, url := range urls {
		go func(u string) {
			start := time.Now()
			data := fetchURL(u)
			data.FetchTime = time.Since(start)
			resultsChan <- data // Send result to channel
		}(url)
	}

	// Collect all results from the channel
	var results []ScrapedData
	for i := 0; i < len(urls); i++ {
		result := <-resultsChan // Receive result from channel
		results = append(results, result)
		fmt.Printf("[%d/%d] %s - %s (%.2fs)\n",
			i+1, len(urls), result.URL, result.Title, result.FetchTime.Seconds())
	}

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal time: %.2f seconds (much faster!)\n", totalTime.Seconds())
	return results
}

// ============================================================================
// ADVANCED CONCURRENT SCRAPER (With WaitGroup)
// ============================================================================
// Better pattern using sync.WaitGroup for coordinating goroutines

func advancedConcurrentScraper(urls []string) []ScrapedData {
	fmt.Println("\n========== ADVANCED CONCURRENT SCRAPER (WaitGroup) ==========")
	fmt.Println("Using sync.WaitGroup for proper synchronization...")

	// sync.WaitGroup coordinates goroutines
	// Useful when you don't know exact count beforehand
	var wg sync.WaitGroup

	// Mutex protects the results slice from concurrent writes
	var mu sync.Mutex
	var results []ScrapedData

	startTime := time.Now()

	// Launch goroutine for each URL
	for _, url := range urls {
		wg.Add(1) // Increment counter for each goroutine

		go func(u string) {
			defer wg.Done() // Decrement counter when done

			start := time.Now()
			data := fetchURL(u)
			data.FetchTime = time.Since(start)

			// Lock mutex before writing to shared slice
			mu.Lock()
			results = append(results, data)
			mu.Unlock()

			fmt.Printf("✓ %s - %s (%.2fs)\n", data.URL, data.Title, data.FetchTime.Seconds())
		}(url)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal time: %.2f seconds\n", totalTime.Seconds())
	return results
}

// ============================================================================
// POOLED CONCURRENT SCRAPER (Limited Concurrency)
// ============================================================================
// Controls the number of concurrent requests (important for rate limiting)

func pooledConcurrentScraper(urls []string, maxConcurrency int) []ScrapedData {
	fmt.Println("\n========== POOLED CONCURRENT SCRAPER ==========")
	fmt.Printf("Max concurrent requests: %d\n\n", maxConcurrency)

	// Semaphore pattern: use buffered channel to limit concurrency
	semaphore := make(chan struct{}, maxConcurrency)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var results []ScrapedData

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			// Acquire semaphore slot
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release slot when done

			start := time.Now()
			data := fetchURL(u)
			data.FetchTime = time.Since(start)

			mu.Lock()
			results = append(results, data)
			mu.Unlock()

			fmt.Printf("✓ %s - %s (%.2fs)\n", data.URL, data.Title, data.FetchTime.Seconds())
		}(url)
	}

	wg.Wait()

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal time: %.2f seconds\n", totalTime.Seconds())
	return results
}

// ============================================================================
// SCRAPER WITH CONTEXT AND TIMEOUT
// ============================================================================
// Production-ready scraper with context cancellation and timeouts

func contextualScraper(ctx context.Context, urls []string, config ScraperConfig) []ScrapedData {
	fmt.Println("\n========== CONTEXTUAL SCRAPER (Production Ready) ==========")
	fmt.Printf("Config: MaxConcurrency=%d, Timeout=%v, Retries=%d\n\n",
		config.MaxConcurrency, config.Timeout, config.RetryCount)

	// Semaphore for concurrency control
	semaphore := make(chan struct{}, config.MaxConcurrency)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var results []ScrapedData

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			// Check if context already cancelled
			select {
			case <-ctx.Done():
				mu.Lock()
				results = append(results, ScrapedData{
					URL:   u,
					Error: ctx.Err(),
				})
				mu.Unlock()
				return
			default:
			}

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Create context with timeout for this specific request
			fetchCtx, cancel := context.WithTimeout(ctx, config.Timeout)
			defer cancel()

			data := fetchURLWithContext(fetchCtx, u, config.RetryCount)

			mu.Lock()
			results = append(results, data)
			mu.Unlock()

			statusStr := "✓"
			if data.Error != nil {
				statusStr = "✗"
			}
			fmt.Printf("%s %s - %s\n", statusStr, data.URL, data.Title)
		}(url)
	}

	wg.Wait()

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal time: %.2f seconds\n", totalTime.Seconds())
	return results
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// fetchURL fetches a URL and extracts the title
func fetchURL(url string) ScrapedData {
	start := time.Now()

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return ScrapedData{
			URL:   url,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// Extract title from HTML
	title := extractTitle(resp.Body)

	return ScrapedData{
		URL:        url,
		Title:      title,
		StatusCode: resp.StatusCode,
		FetchTime:  time.Since(start),
		Error:      nil,
	}
}

// fetchURLWithContext fetches a URL with context and retry logic
func fetchURLWithContext(ctx context.Context, url string, maxRetries int) ScrapedData {
	start := time.Now()

	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Check context cancellation
		select {
		case <-ctx.Done():
			return ScrapedData{
				URL:       url,
				Error:     ctx.Err(),
				FetchTime: time.Since(start),
			}
		default:
		}

		// Create request with context
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			lastErr = err
			continue
		}

		// Execute request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			if attempt < maxRetries {
				time.Sleep(time.Millisecond * 100 * time.Duration(attempt))
				continue
			}
		}

		if resp != nil {
			defer resp.Body.Close()

			title := extractTitle(resp.Body)
			return ScrapedData{
				URL:        url,
				Title:      title,
				StatusCode: resp.StatusCode,
				FetchTime:  time.Since(start),
				Error:      nil,
			}
		}
	}

	return ScrapedData{
		URL:       url,
		Error:     lastErr,
		FetchTime: time.Since(start),
	}
}

// extractTitle extracts the <title> tag from HTML
func extractTitle(body io.Reader) string {
	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		line := scanner.Text()

		// Look for <title> tag
		if strings.Contains(line, "<title>") && strings.Contains(line, "</title>") {
			// Extract content between tags
			startIdx := strings.Index(line, "<title>") + 7
			endIdx := strings.Index(line, "</title>")

			if startIdx < endIdx {
				return strings.TrimSpace(line[startIdx:endIdx])
			}
		}

		// Also check for title in separate lines
		if strings.Contains(line, "<title>") {
			// Find opening tag
			startIdx := strings.Index(line, "<title>") + 7
			if startIdx < len(line) {
				// Collect content until closing tag
				content := line[startIdx:]

				// Continue reading if closing tag not found
				for scanner.Scan() {
					nextLine := scanner.Text()
					content += " " + nextLine

					if strings.Contains(nextLine, "</title>") {
						endIdx := strings.Index(content, "</title>")
						if endIdx > 0 {
							return strings.TrimSpace(content[:endIdx])
						}
						break
					}
				}
				return strings.TrimSpace(content)
			}
		}
	}

	return "No title found"
}

// ============================================================================
// COMPARISON DEMONSTRATION
// ============================================================================

func DemoAllScrapers() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║     WEB SCRAPER CONCURRENCY COMPARISON                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	// Sample URLs (using small pages for demo)
	urls := []string{
		"https://www.example.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.wikipedia.org",
		"https://www.stackoverflow.com",
	}

	// 1. Sequential scraper (slow)
	sequentialResults := basicSequentialScraper(urls)
	printResults("Sequential", sequentialResults)

	// 2. Simple concurrent (fast)
	concurrentResults := simpleConcurrentScraper(urls)
	printResults("Simple Concurrent", concurrentResults)

	// 3. Advanced with WaitGroup
	advancedResults := advancedConcurrentScraper(urls)
	printResults("Advanced (WaitGroup)", advancedResults)

	// 4. Pooled (limited concurrency)
	pooledResults := pooledConcurrentScraper(urls, 2)
	printResults("Pooled (Max 2)", pooledResults)

	// 5. With context and timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config := ScraperConfig{
		MaxConcurrency: 3,
		Timeout:        10 * time.Second,
		RetryCount:     2,
	}
	contextualResults := contextualScraper(ctx, urls, config)
	printResults("Contextual (Production)", contextualResults)
}

// printResults displays scraping results
func printResults(name string, results []ScrapedData) {
	fmt.Printf("\n%s Results Summary:\n", name)
	fmt.Println(strings.Repeat("-", 60))

	successCount := 0
	for _, result := range results {
		if result.Error == nil {
			successCount++
		}
	}

	fmt.Printf("Success: %d/%d\n", successCount, len(results))
	fmt.Println(strings.Repeat("-", 60))
}

// ============================================================================
// ADVANCED PATTERNS
// ============================================================================

// WorkerPool pattern for efficient resource management
type WorkerPool struct {
	workers int
	jobs    chan string
	results chan ScrapedData
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers int, timeout time.Duration) *WorkerPool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	return &WorkerPool{
		workers: numWorkers,
		jobs:    make(chan string, numWorkers*2), // Buffered channel
		results: make(chan ScrapedData, numWorkers*2),
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Start initializes worker goroutines
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)

		go func(workerID int) {
			defer wp.wg.Done()

			for url := range wp.jobs {
				// Process job
				data := fetchURLWithContext(wp.ctx, url, 1)
				wp.results <- data
			}
		}(i)
	}
}

// Submit adds a job to the pool
func (wp *WorkerPool) Submit(url string) error {
	select {
	case wp.jobs <- url:
		return nil
	case <-wp.ctx.Done():
		return errors.New("pool context cancelled")
	default:
		return errors.New("job queue full")
	}
}

// Wait closes the job channel and waits for workers to finish
func (wp *WorkerPool) Wait() []ScrapedData {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)

	var results []ScrapedData
	for result := range wp.results {
		results = append(results, result)
	}

	wp.cancel()
	return results
}

// DemoWorkerPool demonstrates the worker pool pattern
func DemoWorkerPool() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║           WORKER POOL PATTERN DEMONSTRATION            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	urls := []string{
		"https://www.example.com",
		"https://www.golang.org",
		"https://www.github.com",
	}

	pool := NewWorkerPool(2, 30*time.Second)
	pool.Start()

	for _, url := range urls {
		pool.Submit(url)
	}

	results := pool.Wait()

	fmt.Printf("\nWorkerPool processed %d URLs\n", len(results))
	for _, result := range results {
		if result.Error == nil {
			fmt.Printf("✓ %s - %s\n", result.URL, result.Title)
		} else {
			fmt.Printf("✗ %s - Error: %v\n", result.URL, result.Error)
		}
	}
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunWebScraperExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    WEB SCRAPER - CONCURRENT PROGRAMMING CONCEPTS       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	DemoAllScrapers()
	DemoWorkerPool()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║             KEY TAKEAWAYS - WEB SCRAPER                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	takeaways := `
PATTERNS DEMONSTRATED:
1. Sequential: Simple but slow
2. Goroutines + Channels: Concurrent but uncontrolled
3. WaitGroup: Better synchronization
4. Semaphore/Pool: Rate limiting and resource control
5. Context: Cancellation and timeouts
6. Worker Pool: Production-ready pattern

WHY CONCURRENCY MATTERS:
- I/O-bound operations (like HTTP requests) benefit greatly
- Sequential: Must wait for each response
- Concurrent: Can fetch multiple URLs simultaneously
- With proper pooling: Limits load on servers (good practice)

IMPORTANT PATTERNS:
- Always use sync.Mutex to protect shared data
- Always check context cancellation
- Always close channels when done
- Always defer resource cleanup
- Use buffered channels when you know the size
- Use worker pools for production applications
`
	fmt.Println(takeaways)

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                 ALL EXAMPLES COMPLETE                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}

// WebScraper is the legacy function name
func WebScraper() {
	RunWebScraperExamples()
}
