package concepts

import (
	_ "bufio"
	"fmt"
	_ "io"
	_ "net/http"
	"net/url"
	_ "net/url"
	_ "strings"
	_ "time"
)

// ============================================================================
// HTTP AND REST BASICS CONCEPTS
// ============================================================================
// This file covers:
// - HTTP server fundamentals
// - HTTP client operations
// - Request/Response handling
// - Headers and status codes
// - URL parsing and query parameters
// - REST API principles
// ============================================================================

// HTTPServerBasicsDemo demonstrates creating a simple HTTP server
func HTTPServerBasicsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         HTTP SERVER BASICS DEMO                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	// --- SIMPLE HANDLER FUNCTION ---
	fmt.Println("1. SIMPLE HANDLER FUNCTION:")
	fmt.Println(`
// Define a handler function - receives ResponseWriter and Request
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // w: writes the response
    // r: contains request information
    
    fmt.Fprintf(w, "Hello, World!")
}

// Register the handler
http.HandleFunc("/hello", HelloHandler)

// Start the server on port 8080
http.ListenAndServe(":8080", nil)
`)

	// --- HANDLER WITH DIFFERENT METHODS ---
	fmt.Println("\n2. HANDLING DIFFERENT HTTP METHODS:")
	fmt.Println(`
func MethodHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintf(w, "GET request received")
    case http.MethodPost:
        fmt.Fprintf(w, "POST request received")
    case http.MethodPut:
        fmt.Fprintf(w, "PUT request received")
    case http.MethodDelete:
        fmt.Fprintf(w, "DELETE request received")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// HTTP Methods explain WHAT action to perform:
// GET    - Retrieve data (safe, idempotent)
// POST   - Create new data (not idempotent)
// PUT    - Replace entire resource (idempotent)
// PATCH  - Partial update
// DELETE - Remove resource (idempotent)
// HEAD   - Like GET but no body
// OPTIONS- Describe communication options
`)

	// --- STATUS CODES EXPLAINED ---
	fmt.Println("\n3. HTTP STATUS CODES:")
	fmt.Println(`
// 2xx - Success (Request successful)
http.StatusOK                  // 200 - Request succeeded
http.StatusCreated             // 201 - Resource created
http.StatusNoContent           // 204 - Successful, no response body

// 3xx - Redirection (Further action needed)
http.StatusMovedPermanently    // 301 - Resource moved permanently
http.StatusFound               // 302 - Temporary redirect
http.StatusNotModified         // 304 - Use cached version

// 4xx - Client Error (Bad request)
http.StatusBadRequest          // 400 - Malformed request
http.StatusUnauthorized        // 401 - Authentication required
http.StatusForbidden           // 403 - Access denied
http.StatusNotFound            // 404 - Resource not found
http.StatusConflict            // 409 - Request conflicts with server state

// 5xx - Server Error (Server failed)
http.StatusInternalServerError // 500 - Server error
http.StatusNotImplemented      // 501 - Feature not implemented
http.StatusServiceUnavailable  // 503 - Server temporarily unavailable

// Writing status codes:
w.WriteHeader(http.StatusCreated) // Must be called before writing body
fmt.Fprintf(w, "Resource created") // Response body
`)

	fmt.Println("✓ HTTP Server basics demonstrated")
}

// HTTPClientBasicsDemo demonstrates making HTTP requests
func HTTPClientBasicsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         HTTP CLIENT BASICS DEMO                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. BASIC GET REQUEST:")
	fmt.Println(`
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    // Network error, timeout, etc.
    return err
}
defer resp.Body.Close() // CRITICAL: Always close the body

// Check status code
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("server returned %d", resp.StatusCode)
}

// Read response body
body, err := io.ReadAll(resp.Body)
if err != nil {
    return err
}

data := string(body)
`)

	fmt.Println("\n2. MAKING A POST REQUEST:")
	fmt.Println(`
payload := strings.NewReader(` + "`" + `{"name": "John", "age": 30}` + "`" + `)

req, err := http.NewRequest(http.MethodPost, 
    "https://api.example.com/users", payload)
if err != nil {
    return err
}

// Set headers
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token123")

// Execute request
client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()

// Process response...
`)

	fmt.Println("\n3. HTTP CLIENT WITH TIMEOUT:")
	fmt.Println(`
// Good practice: Always set timeouts to avoid hanging
client := &http.Client{
    Timeout: 5 * time.Second,
}

req, _ := http.NewRequest(http.MethodGet, url, nil)
resp, err := client.Do(req)
if err != nil {
    // Could be timeout error
    return err
}
defer resp.Body.Close()

// Without timeout, requests can hang indefinitely!
`)

	fmt.Println("\n✓ HTTP Client basics demonstrated")
}

// RequestResponseHandlingDemo shows request/response details
func RequestResponseHandlingDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║     REQUEST AND RESPONSE HANDLING DEMO                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. REQUEST PROPERTIES:")
	fmt.Println(`
func DetailedHandler(w http.ResponseWriter, r *http.Request) {
    // Request Method
    fmt.Println("Method:", r.Method)  // "GET", "POST", etc.
    
    // Request URL
    fmt.Println("Path:", r.URL.Path)
    fmt.Println("Query:", r.URL.RawQuery)
    
    // Request Headers
    fmt.Println("User-Agent:", r.Header.Get("User-Agent"))
    fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
    
    // Request Body (for POST/PUT)
    body, err := io.ReadAll(r.Body)
    defer r.Body.Close()
    
    // Remote address
    fmt.Println("Remote Address:", r.RemoteAddr)
    
    // Protocol
    fmt.Println("Protocol:", r.Proto) // "HTTP/1.1"
}
`)

	fmt.Println("\n2. RESPONSE WRITING:")
	fmt.Println(`
func ResponseExample(w http.ResponseWriter, r *http.Request) {
    // Set headers BEFORE writing status or body
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "value")
    
    // Write status code (default is 200)
    w.WriteHeader(http.StatusOK)
    
    // Write response body
    fmt.Fprintf(w, ` + "`{\"status\": \"success\"}" + `)
    
    // After WriteHeader, you cannot change headers!
}
`)

	fmt.Println("\n3. RESPONSE WRITER INTERFACE:")
	fmt.Println(`
// ResponseWriter interface:
type ResponseWriter interface {
    // Write data to response body
    Write([]byte) (int, error)
    
    // Set response header values
    Header() Header
    
    // Send status code (can only be called once!)
    WriteHeader(statusCode int)
}

// Common pattern:
w.Header().Set("Content-Type", "text/plain")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "Response body")
`)

	fmt.Println("✓ Request/Response handling demonstrated")
}

// HeadersAndCookiesDemo shows header and cookie handling
func HeadersAndCookiesDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       HEADERS AND COOKIES DEMO                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. WORKING WITH HEADERS:")
	fmt.Println(`
func HeaderExample(w http.ResponseWriter, r *http.Request) {
    // REQUEST HEADERS
    // Get single header value
    userAgent := r.Header.Get("User-Agent")
    
    // Get all values for a header (some can have multiple)
    acceptLanguages := r.Header.Values("Accept-Language")
    
    // Check if header exists
    if _, exists := r.Header["Authorization"]; exists {
        // Authorization header is present
    }
    
    // RESPONSE HEADERS
    // Set response headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Cache-Control", "max-age=3600")
    w.Header().Set("X-Request-ID", generateRequestID())
    
    // Add multiple header values
    w.Header().Add("Set-Cookie", "session=abc123")
    w.Header().Add("Set-Cookie", "theme=dark")
    
    // Delete a header
    w.Header().Del("X-Powered-By")
}
`)

	fmt.Println("\n2. COOKIES:")
	fmt.Println(`
func CookieExample(w http.ResponseWriter, r *http.Request) {
    // GET A COOKIE
    cookie, err := r.Cookie("sessionID")
    if err == http.ErrNoCookie {
        // Cookie not found
        fmt.Println("No session cookie")
    } else if err == nil {
        sessionID := cookie.Value
        fmt.Println("Session ID:", sessionID)
    }
    
    // SET A COOKIE
    newCookie := &http.Cookie{
        Name:     "sessionID",
        Value:    "abc123def456",
        Path:     "/",
        MaxAge:   86400 * 7,  // 7 days in seconds
        HttpOnly: true,        // Not accessible from JavaScript
        Secure:   true,        // Only HTTPS
        SameSite: http.SameSiteLaxMode,
    }
    http.SetCookie(w, newCookie)
    
    // GET ALL COOKIES
    allCookies := r.Cookies()
    for _, cookie := range allCookies {
        fmt.Println(cookie.Name, "=", cookie.Value)
    }
}

// Cookie attributes explained:
// Name:     Cookie name
// Value:    Cookie value
// Path:     URL path where cookie is valid
// Domain:   Domain where cookie is valid
// MaxAge:   Lifetime in seconds (0 = delete, -1 = session)
// Expires:  Expiration time (older format)
// HttpOnly: Not accessible from JavaScript (security)
// Secure:   Only sent over HTTPS (security)
// SameSite: Prevent CSRF attacks (Strict, Lax, None)
`)

	fmt.Println("✓ Headers and cookies demonstrated")
}

// URLParsingDemo shows URL and query parameter handling
func URLParsingDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       URL PARSING AND QUERY PARAMS DEMO                ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. PARSING URLS:")
	urlString := "https://api.example.com:8080/users/123?role=admin&active=true#section"

	parsedURL, err := url.Parse(urlString)
	if err == nil {
		fmt.Println("\nURL Components:")
		fmt.Printf("  Scheme:   %s\n", parsedURL.Scheme)     // https
		fmt.Printf("  Host:     %s\n", parsedURL.Host)       // api.example.com:8080
		fmt.Printf("  Hostname: %s\n", parsedURL.Hostname()) // api.example.com
		fmt.Printf("  Port:     %s\n", parsedURL.Port())     // 8080
		fmt.Printf("  Path:     %s\n", parsedURL.Path)       // /users/123
		fmt.Printf("  Query:    %s\n", parsedURL.RawQuery)   // role=admin&active=true
		fmt.Printf("  Fragment: %s\n", parsedURL.Fragment)   // section
	}

	fmt.Println("\n2. QUERY PARAMETERS IN HANDLER:")
	fmt.Println(`
func QueryParamHandler(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters
    r.ParseForm() // Must call this first
    
    // Get single query parameter
    page := r.FormValue("page")  // Returns empty string if not present
    
    // Get multiple values for same parameter
    tags := r.Form["tag"]  // Returns []string
    
    // Iterate over all parameters
    for key, values := range r.Form {
        for _, value := range values {
            fmt.Printf("%s = %s\n", key, value)
        }
    }
    
    // URL: /search?q=golang&sort=date&sort=popularity
    q := r.FormValue("q")           // "golang"
    sorts := r.Form["sort"]         // ["date", "popularity"]
}
`)

	fmt.Println("\n3. BUILDING URLS WITH QUERY PARAMETERS:")
	fmt.Println(`
// Build URL with query parameters
params := url.Values{}
params.Add("search", "golang")
params.Add("sort", "date")
params.Add("limit", "10")

baseURL, _ := url.Parse("https://api.example.com/search")
baseURL.RawQuery = params.Encode()

fullURL := baseURL.String()
// Result: https://api.example.com/search?search=golang&sort=date&limit=10
`)

	fmt.Println("✓ URL parsing and query parameters demonstrated")
}

// RESTBasicsDemo covers REST API principles
func RESTBasicsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         REST API BASICS DEMO                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. REST PRINCIPLES:")
	fmt.Println(`
// REST = Representational State Transfer
// RESTful design guidelines:

1. RESOURCE-ORIENTED:
   - Think in terms of resources (users, posts, comments)
   - Every resource has a unique URL (identifier)
   - /users/123 identifies user with ID 123

2. USE HTTP METHODS FOR OPERATIONS:
   GET     /users          - List all users
   GET     /users/123      - Get specific user
   POST    /users          - Create new user
   PUT     /users/123      - Replace user 123 completely
   PATCH   /users/123      - Partially update user 123
   DELETE  /users/123      - Delete user 123

3. STATELESS:
   - Each request contains all needed information
   - Server doesn't store client context
   - Same request always produces same result

4. REPRESENTATION:
   - Resource can be represented in different formats
   - Usually JSON, but could be XML, CSV, etc.
   - Content-Type header specifies format

5. PROPER USE OF STATUS CODES:
   - 200 OK - Request succeeded
   - 201 Created - Resource successfully created
   - 204 No Content - Successful but no response body
   - 400 Bad Request - Invalid request data
   - 401 Unauthorized - Authentication required
   - 403 Forbidden - Access denied
   - 404 Not Found - Resource doesn't exist
   - 500 Internal Server Error - Server error
`)

	fmt.Println("\n2. REST API DESIGN PATTERNS:")
	fmt.Println(`
// Good REST API structure:

// Collections (plural)
GET    /api/v1/users
POST   /api/v1/users

// Individual resources (with ID)
GET    /api/v1/users/123
PUT    /api/v1/users/123
PATCH  /api/v1/users/123
DELETE /api/v1/users/123

// Nested resources
GET    /api/v1/users/123/posts        - Posts by user 123
GET    /api/v1/users/123/posts/456    - Specific post
POST   /api/v1/users/123/posts        - Create post for user

// Query parameters for filtering/pagination
GET    /api/v1/users?role=admin&limit=10&offset=20

// API versioning (important!)
/api/v1/users      - Version 1 API
/api/v2/users      - Version 2 API (can be different)
`)

	fmt.Println("\n3. REQUEST/RESPONSE EXAMPLES:")
	fmt.Println(`
// REQUEST to create user:
POST /api/v1/users HTTP/1.1
Host: api.example.com
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user"
}

// RESPONSE (201 Created):
HTTP/1.1 201 Created
Content-Type: application/json
Location: /api/v1/users/123

{
  "id": 123,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "created_at": "2024-01-15T10:30:00Z"
}
`)

	fmt.Println("✓ REST API basics demonstrated")
}

// PracticalHTTPServerDemo shows a practical example
func PracticalHTTPServerDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║     PRACTICAL HTTP SERVER EXAMPLE                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("COMPLETE EXAMPLE - Simple API Server:")
	fmt.Println(`
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

type User struct {
    ID    int    ` + "`" + `json:"id"` + "`" + `
    Name  string ` + "`" + `json:"name"` + "`" + `
    Email string ` + "`" + `json:"email"` + "`" + `
}

var users = map[int]User{
    1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
    2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // List all users
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
        
    case http.MethodPost:
        // Create new user
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        user.ID = len(users) + 1
        users[user.ID] = user
        
        w.WriteHeader(http.StatusCreated)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
        
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL
    idStr := r.URL.Path[len("/users/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    user, exists := users[id]
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    switch r.Method {
    case http.MethodGet:
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
        
    case http.MethodDelete:
        delete(users, id)
        w.WriteHeader(http.StatusNoContent)
        
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler)
    
    fmt.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
`)

	fmt.Println("✓ Practical HTTP server demonstrated")
}

// RunHTTPAndRESTExamples executes all HTTP and REST demos
func RunHTTPAndRESTExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    HTTP AND REST BASICS - COMPLETE GUIDE               ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	HTTPServerBasicsDemo()
	HTTPClientBasicsDemo()
	RequestResponseHandlingDemo()
	HeadersAndCookiesDemo()
	URLParsingDemo()
	RESTBasicsDemo()
	PracticalHTTPServerDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         ALL HTTP AND REST EXAMPLES COMPLETE             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ HTTP is request-response protocol")
	fmt.Println("✓ Use appropriate HTTP methods (GET, POST, PUT, DELETE)")
	fmt.Println("✓ Always set correct status codes")
	fmt.Println("✓ Close response bodies with defer")
	fmt.Println("✓ Set timeouts on HTTP clients")
	fmt.Println("✓ Parse URLs and query parameters correctly")
	fmt.Println("✓ Headers and cookies for metadata")
	fmt.Println("✓ REST design: resource-oriented with standard methods\n")
}
