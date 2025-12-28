/*
HTTP AND REST BASICS IN GO - Complete Learning Guide
============================================================================

HTTP (HyperText Transfer Protocol) is the foundation of web communication.
REST (Representational State Transfer) is an architectural style for APIs.

TOPICS COVERED:
- HTTP server fundamentals
- HTTP client operations
- Request and Response handling
- Headers and cookies
- URL parsing and query parameters
- REST API principles and design patterns
*/

package concepts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ============================================================================
// 1. HTTP SERVER FUNDAMENTALS
// ============================================================================

/*
HTTP SERVER BASICS:
- A handler is a function that receives ResponseWriter and Request
- ResponseWriter: allows sending response back to client
- Request: contains information about the client's request
- HTTP methods: GET (retrieve), POST (create), PUT (replace), DELETE (remove)

HTTP Methods:
- GET    - Retrieve data (safe, idempotent)
- POST   - Create new data (not idempotent)
- PUT    - Replace entire resource (idempotent)
- PATCH  - Partial update
- DELETE - Remove resource (idempotent)
- HEAD   - Like GET but no body
- OPTIONS- Describe communication options

HTTP Status Codes (2xx=success, 3xx=redirect, 4xx=client error, 5xx=server error):
- 200 OK, 201 Created, 204 No Content
- 301 Moved Permanently, 302 Found
- 400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found
- 500 Internal Server Error, 503 Service Unavailable
*/

func HTTPServerBasicsDemo() {
	fmt.Println("\n========== HTTP SERVER BASICS ==========")

	// Example handler that processes different HTTP methods
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		// w: ResponseWriter - used to send response
		// r: Request - contains client request info
		fmt.Fprintf(w, "Hello, World!")
	}

	methodHandler := func(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("\n1. SIMPLE HANDLER:")
	fmt.Println("   - Handler receives ResponseWriter and Request")
	fmt.Println("   - Write response with fmt.Fprintf(w, ...)")
	fmt.Printf("   - Example: http.HandleFunc(\"/hello\", helloHandler)\n")

	fmt.Println("\n2. METHOD-BASED ROUTING:")
	fmt.Println("   - Check r.Method to handle different HTTP methods")
	fmt.Println("   - Return appropriate status codes")
	fmt.Printf("   - Example: switch r.Method { case http.MethodGet: ... }\n")

	// Demonstrate status codes
	fmt.Println("\n3. HTTP STATUS CODES:")
	fmt.Println("   2xx: Success")
	fmt.Printf("      - %d OK (request succeeded)\n", http.StatusOK)
	fmt.Printf("      - %d Created (resource created)\n", http.StatusCreated)
	fmt.Printf("      - %d No Content (success, no body)\n", http.StatusNoContent)
	fmt.Println("   4xx: Client Error")
	fmt.Printf("      - %d Bad Request (malformed)\n", http.StatusBadRequest)
	fmt.Printf("      - %d Unauthorized (auth required)\n", http.StatusUnauthorized)
	fmt.Printf("      - %d Not Found (resource missing)\n", http.StatusNotFound)
	fmt.Println("   5xx: Server Error")
	fmt.Printf("      - %d Internal Server Error\n", http.StatusInternalServerError)

	// Show that handlers are functions
	_ = helloHandler
	_ = methodHandler

	fmt.Println("\n✓ HTTP Server basics demonstrated")
}

// ============================================================================
// 2. HTTP CLIENT OPERATIONS
// ============================================================================

/*
HTTP CLIENT:
- Make requests to HTTP servers
- http.Get() for simple GET requests
- http.NewRequest() for complex requests with custom headers
- Always close response bodies with defer (important!)
- Always set timeouts to prevent hanging
- Check status codes and handle errors

Key patterns:
1. Simple GET: http.Get(url) returns *Response, error
2. Custom requests: http.NewRequest() then client.Do()
3. Timeouts: create http.Client with Timeout field
4. Error handling: both network errors and status codes
*/

func HTTPClientBasicsDemo() {
	fmt.Println("\n========== HTTP CLIENT BASICS ==========")

	fmt.Println("\n1. SIMPLE GET REQUEST:")
	// Example GET request
	makeSimpleGet := func(url string) (string, error) {
		resp, err := http.Get(url)
		if err != nil {
			// Network error, timeout, etc.
			return "", err
		}
		defer resp.Body.Close() // CRITICAL: Always close the body

		// Check status code
		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("server returned %d", resp.StatusCode)
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		return string(body), nil
	}

	fmt.Println("   - http.Get(url) returns response and error")
	fmt.Println("   - defer resp.Body.Close() (important!)")
	fmt.Println("   - Check resp.StatusCode")
	fmt.Println("   - Read body with io.ReadAll(resp.Body)")
	_ = makeSimpleGet

	fmt.Println("\n2. CUSTOM REQUESTS WITH HEADERS:")
	// Example POST request with custom headers
	makeCustomRequest := func(url, jsonPayload string) error {
		req, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			return err
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer token123")
		req.Header.Set("X-Custom-Header", "custom-value")

		// Execute request with timeout
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return nil
	}

	fmt.Println("   - http.NewRequest() to create request")
	fmt.Println("   - Set headers with req.Header.Set()")
	fmt.Println("   - client.Do(req) to execute")
	fmt.Println("   - Don't forget defer resp.Body.Close()")
	_ = makeCustomRequest

	fmt.Println("\n3. TIMEOUT HANDLING:")
	// Example with timeout
	makeRequestWithTimeout := func(url string) error {
		client := &http.Client{
			Timeout: 5 * 1000000000, // 5 seconds in nanoseconds (see time.Second)
		}

		resp, err := client.Get(url)
		if err != nil {
			// Could be timeout error
			return err
		}
		defer resp.Body.Close()

		return nil
	}

	fmt.Println("   - Create http.Client with Timeout")
	fmt.Println("   - Prevents requests from hanging indefinitely")
	fmt.Println("   - Example: client := &http.Client{Timeout: 5 * time.Second}")
	_ = makeRequestWithTimeout

	fmt.Println("\n✓ HTTP Client basics demonstrated")
}

// ============================================================================
// 3. REQUEST AND RESPONSE HANDLING
// ============================================================================

/*
REQUEST OBJECT (*http.Request):
- Method: GET, POST, PUT, DELETE, etc.
- URL: parsed URL with path, query params, fragments
- Header: HTTP headers
- Body: request body (for POST, PUT, etc.)
- RemoteAddr: client's IP address
- Proto: HTTP protocol version

RESPONSE WRITER (http.ResponseWriter):
- Header(): get headers map to set response headers
- Write([]byte): write response body
- WriteHeader(int): set status code (only once!)
- Order matters: set headers, then WriteHeader, then Write
*/

func RequestResponseHandlingDemo() {
	fmt.Println("\n========== REQUEST AND RESPONSE HANDLING ==========")

	// Example handler that accesses request properties
	detailedHandler := func(w http.ResponseWriter, r *http.Request) {
		// Request Method
		fmt.Printf("Method: %s\n", r.Method) // "GET", "POST", etc.

		// Request URL
		fmt.Printf("Path: %s\n", r.URL.Path)
		fmt.Printf("Query: %s\n", r.URL.RawQuery)

		// Request Headers
		fmt.Printf("User-Agent: %s\n", r.Header.Get("User-Agent"))
		fmt.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))

		// Request Body (for POST/PUT)
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err == nil {
			fmt.Printf("Body: %s\n", string(body))
		}

		// Remote address
		fmt.Printf("Remote Address: %s\n", r.RemoteAddr)

		// Protocol
		fmt.Printf("Protocol: %s\n", r.Proto) // "HTTP/1.1"
	}

	fmt.Println("\n1. REQUEST PROPERTIES:")
	fmt.Println("   - r.Method: HTTP method (GET, POST, etc.)")
	fmt.Println("   - r.URL.Path: request path")
	fmt.Println("   - r.URL.RawQuery: query string")
	fmt.Println("   - r.Header.Get(): get header value")
	fmt.Println("   - r.Body: request body (as io.ReadCloser)")
	fmt.Println("   - r.RemoteAddr: client IP address")
	fmt.Println("   - r.Proto: protocol version")
	_ = detailedHandler

	// Example response writing
	responseExample := func(w http.ResponseWriter, r *http.Request) {
		// Set headers BEFORE writing status or body
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "value")

		// Write status code (default is 200)
		// WARNING: Can only be called ONCE!
		w.WriteHeader(http.StatusOK)

		// Write response body
		fmt.Fprintf(w, `{"status": "success"}`)

		// After WriteHeader, you cannot change headers!
	}

	fmt.Println("\n2. RESPONSE WRITING:")
	fmt.Println("   - w.Header().Set(key, value): set response headers")
	fmt.Println("   - w.WriteHeader(statusCode): send status (only once!)")
	fmt.Println("   - fmt.Fprintf(w, ...): write response body")
	fmt.Println("   - Order: set headers → WriteHeader → Write")
	_ = responseExample

	fmt.Println("\n3. RESPONSEWRITER INTERFACE:")
	fmt.Println("   - Header() Header: get headers map")
	fmt.Println("   - Write([]byte): write body")
	fmt.Println("   - WriteHeader(int): send status code")
	fmt.Println("   - Status code can only be written once!")

	fmt.Println("\n✓ Request/Response handling demonstrated")
}

// ============================================================================
// 4. HEADERS AND COOKIES
// ============================================================================

/*
HEADERS:
- Key-value pairs sent with request/response
- Case-insensitive
- Request headers: User-Agent, Content-Type, Authorization
- Response headers: Content-Type, Cache-Control, Set-Cookie

COOKIES:
- Small data stored on client
- Sent with each request to domain
- Attributes: Name, Value, Path, Domain, MaxAge, HttpOnly, Secure, SameSite
- HttpOnly: prevents JavaScript access (security)
- Secure: only sent over HTTPS (security)
- SameSite: prevents CSRF attacks
*/

func HeadersAndCookiesDemo() {
	fmt.Println("\n========== HEADERS AND COOKIES ==========")

	// Example header handling
	headerExample := func(w http.ResponseWriter, r *http.Request) {
		// REQUEST HEADERS
		// Get single header value
		userAgent := r.Header.Get("User-Agent")
		fmt.Printf("User-Agent: %s\n", userAgent)

		// Get all values for a header (some can have multiple)
		acceptLanguages := r.Header.Values("Accept-Language")
		fmt.Printf("Accept-Language: %v\n", acceptLanguages)

		// Check if header exists
		if _, exists := r.Header["Authorization"]; exists {
			// Authorization header is present
		}

		// RESPONSE HEADERS
		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "max-age=3600")
		w.Header().Set("X-Request-ID", "req123")

		// Add multiple header values
		w.Header().Add("Set-Cookie", "session=abc123")
		w.Header().Add("Set-Cookie", "theme=dark")

		// Delete a header
		w.Header().Del("X-Powered-By")
	}

	fmt.Println("\n1. WORKING WITH HEADERS:")
	fmt.Println("   - r.Header.Get(key): get header value")
	fmt.Println("   - r.Header.Values(key): get all values")
	fmt.Println("   - w.Header().Set(key, value): set response header")
	fmt.Println("   - w.Header().Add(key, value): add multiple values")
	fmt.Println("   - w.Header().Del(key): delete header")
	_ = headerExample

	// Example cookie handling
	cookieExample := func(w http.ResponseWriter, r *http.Request) {
		// GET A COOKIE
		cookie, err := r.Cookie("sessionID")
		if err == http.ErrNoCookie {
			fmt.Println("No session cookie")
		} else if err == nil {
			fmt.Printf("Session ID: %s\n", cookie.Value)
		}

		// SET A COOKIE
		newCookie := &http.Cookie{
			Name:     "sessionID",
			Value:    "abc123def456",
			Path:     "/",
			MaxAge:   86400 * 7,            // 7 days in seconds
			HttpOnly: true,                 // Not accessible from JavaScript
			Secure:   false,                // Only HTTPS (set to true in production)
			SameSite: http.SameSiteLaxMode, // CSRF protection
		}
		http.SetCookie(w, newCookie)

		// GET ALL COOKIES
		allCookies := r.Cookies()
		for _, cookie := range allCookies {
			fmt.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
		}
	}

	fmt.Println("\n2. COOKIES:")
	fmt.Println("   - r.Cookie(name): get specific cookie")
	fmt.Println("   - r.Cookies(): get all cookies")
	fmt.Println("   - http.SetCookie(w, cookie): set cookie")
	fmt.Println("   - Cookie attributes: Name, Value, Path, MaxAge, HttpOnly, Secure")
	_ = cookieExample

	fmt.Println("\n✓ Headers and cookies demonstrated")
}

// ============================================================================
// 5. URL PARSING AND QUERY PARAMETERS
// ============================================================================

/*
URL STRUCTURE:
scheme://host:port/path?query#fragment
https://api.example.com:8080/users?role=admin&sort=date#section

URL Components:
- Scheme: protocol (http, https, ftp)
- Host: hostname:port
- Path: resource path
- RawQuery: query string (key=value&key=value)
- Fragment: anchor/section reference

Query Parameters:
- Key-value pairs in URL
- Multiple values can have same key
- Must be parsed with r.ParseForm()
- Access with r.FormValue(key) or r.Form[key]
*/

func URLParsingDemo() {
	fmt.Println("\n========== URL PARSING AND QUERY PARAMS ==========")

	fmt.Println("\n1. PARSING URLS:")
	urlString := "https://api.example.com:8080/users/123?role=admin&active=true#section"

	parsedURL, err := url.Parse(urlString)
	if err == nil {
		fmt.Println("   URL Components:")
		fmt.Printf("   - Scheme:   %s (protocol)\n", parsedURL.Scheme)          // https
		fmt.Printf("   - Host:     %s (hostname:port)\n", parsedURL.Host)       // api.example.com:8080
		fmt.Printf("   - Hostname: %s (just hostname)\n", parsedURL.Hostname()) // api.example.com
		fmt.Printf("   - Port:     %s (just port)\n", parsedURL.Port())         // 8080
		fmt.Printf("   - Path:     %s (resource path)\n", parsedURL.Path)       // /users/123
		fmt.Printf("   - Query:    %s (query string)\n", parsedURL.RawQuery)    // role=admin&active=true
		fmt.Printf("   - Fragment: %s (anchor)\n", parsedURL.Fragment)          // section
	}

	fmt.Println("\n2. QUERY PARAMETERS IN HANDLER:")
	// Example handler for query parameters
	queryParamHandler := func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters (must call first!)
		r.ParseForm()

		// Get single query parameter
		page := r.FormValue("page") // Returns empty string if not present
		fmt.Printf("Page: %s\n", page)

		// Get multiple values for same parameter
		// URL: /search?tag=golang&tag=web&tag=api
		tags := r.Form["tag"] // Returns []string
		fmt.Printf("Tags: %v\n", tags)

		// Iterate over all parameters
		for key, values := range r.Form {
			for _, value := range values {
				fmt.Printf("%s = %s\n", key, value)
			}
		}
	}

	fmt.Println("   - Call r.ParseForm() first")
	fmt.Println("   - r.FormValue(key): get single value")
	fmt.Println("   - r.Form[key]: get all values for key ([]string)")
	fmt.Println("   - Iterate with for range r.Form")
	_ = queryParamHandler

	fmt.Println("\n3. BUILDING URLS WITH QUERY PARAMETERS:")
	// Build URL with query parameters
	params := url.Values{}
	params.Add("search", "golang")
	params.Add("sort", "date")
	params.Add("limit", "10")

	baseURL, _ := url.Parse("https://api.example.com/search")
	baseURL.RawQuery = params.Encode()

	fullURL := baseURL.String()
	fmt.Printf("   Built URL: %s\n", fullURL)
	// Result: https://api.example.com/search?search=golang&sort=date&limit=10

	fmt.Println("\n✓ URL parsing and query parameters demonstrated")
}

// ============================================================================
// 6. REST API PRINCIPLES AND DESIGN
// ============================================================================

/*
REST (Representational State Transfer):
- Architectural style for APIs
- Uses HTTP methods for operations
- Resource-oriented (not action-oriented)
- Stateless (each request contains all needed info)
- Standard status codes for results

Key Principles:
1. Resource-Oriented: Think in resources (users, posts, etc.)
2. HTTP Methods: GET, POST, PUT, PATCH, DELETE
3. Standard URLs: /api/v1/users, /api/v1/users/123
4. Stateless: No session data on server
5. Status Codes: Indicate result (success, error, redirect)

RESTful URL Patterns:
GET     /api/v1/users           - List all users
GET     /api/v1/users/123       - Get specific user
POST    /api/v1/users           - Create new user
PUT     /api/v1/users/123       - Replace user 123
PATCH   /api/v1/users/123       - Partially update user 123
DELETE  /api/v1/users/123       - Delete user 123

Nested Resources:
GET     /api/v1/users/123/posts        - Posts by user 123
GET     /api/v1/users/123/posts/456    - Specific post
POST    /api/v1/users/123/posts        - Create post for user

Query Parameters for filtering/pagination:
GET     /api/v1/users?role=admin&limit=10&offset=20
*/

func RESTBasicsDemo() {
	fmt.Println("\n========== REST API BASICS ==========")

	fmt.Println("\n1. REST PRINCIPLES:")
	fmt.Println("   - Resource-oriented (not action-oriented)")
	fmt.Println("   - Use HTTP methods for operations")
	fmt.Println("   - Stateless design")
	fmt.Println("   - Standard status codes")
	fmt.Println("   - URL patterns: /api/v1/resource/id")

	fmt.Println("\n2. RESTful ENDPOINT PATTERNS:")
	fmt.Println("   GET    /api/v1/users           - List all users")
	fmt.Println("   GET    /api/v1/users/123       - Get user 123")
	fmt.Println("   POST   /api/v1/users           - Create user")
	fmt.Println("   PUT    /api/v1/users/123       - Replace user 123")
	fmt.Println("   PATCH  /api/v1/users/123       - Update user 123")
	fmt.Println("   DELETE /api/v1/users/123       - Delete user 123")

	fmt.Println("\n3. NESTED RESOURCES:")
	fmt.Println("   GET    /api/v1/users/123/posts        - Posts by user 123")
	fmt.Println("   POST   /api/v1/users/123/posts        - Create post for user")
	fmt.Println("   GET    /api/v1/users/123/posts/456    - Specific post")

	fmt.Println("\n4. QUERY PARAMETERS:")
	fmt.Println("   /api/v1/users?role=admin&limit=10&offset=20")
	fmt.Println("   - role=admin: filter by role")
	fmt.Println("   - limit=10: pagination limit")
	fmt.Println("   - offset=20: pagination offset")

	fmt.Println("\n5. API VERSIONING:")
	fmt.Println("   /api/v1/users   - Version 1 API")
	fmt.Println("   /api/v2/users   - Version 2 API (can be different)")

	fmt.Println("\n✓ REST API basics demonstrated")
}

// ============================================================================
// 7. PRACTICAL EXAMPLE - SIMPLE REST API
// ============================================================================

/*
This example shows a practical REST API implementation.
It's a simplified User API that demonstrates:
- Routing based on HTTP methods
- JSON encoding/decoding
- Status codes
- Error handling
*/

// User struct for API example
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PracticalHTTPServerDemo() {
	fmt.Println("\n========== PRACTICAL REST API EXAMPLE ==========")

	// In-memory storage for example
	users := map[int]User{
		1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
		2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
	}

	// Handler for /users (list and create)
	usersHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			// List all users
			fmt.Fprintf(w, "[")
			i := 0
			for _, user := range users {
				if i > 0 {
					fmt.Fprintf(w, ", ")
				}
				fmt.Fprintf(w, `{"id":%d,"name":"%s","email":"%s"}`,
					user.ID, user.Name, user.Email)
				i++
			}
			fmt.Fprintf(w, "]")

		case http.MethodPost:
			// Create new user (simplified)
			newID := len(users) + 1
			newUser := User{ID: newID, Name: "NewUser", Email: "new@example.com"}
			users[newID] = newUser

			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, `{"id":%d,"name":"%s","email":"%s"}`,
				newUser.ID, newUser.Name, newUser.Email)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	fmt.Println("\n1. HANDLER FOR /users (list and create):")
	fmt.Println("   GET /users   → List all users")
	fmt.Println("   POST /users  → Create new user")
	_ = usersHandler

	fmt.Println("\n2. TYPICAL REST API FLOW:")
	fmt.Println("   1. Client sends GET /api/v1/users/123")
	fmt.Println("   2. Server finds user by ID")
	fmt.Println("   3. Server returns JSON with 200 OK")
	fmt.Println("   4. Or returns 404 Not Found if missing")

	fmt.Println("\n3. ERROR RESPONSES:")
	fmt.Println("   - 400 Bad Request: invalid data in request")
	fmt.Println("   - 401 Unauthorized: authentication required")
	fmt.Println("   - 403 Forbidden: access denied")
	fmt.Println("   - 404 Not Found: resource doesn't exist")
	fmt.Println("   - 500 Internal Server Error: server error")

	fmt.Println("\n✓ Practical REST API example demonstrated")
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

// RunHTTPAndRESTExamples runs all HTTP and REST demonstrations
func RunHTTPAndRESTExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
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
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Println("\nKEY TAKEAWAYS:")
	fmt.Println("✓ HTTP is request-response protocol")
	fmt.Println("✓ Handlers receive ResponseWriter and Request")
	fmt.Println("✓ Use appropriate HTTP methods (GET, POST, PUT, DELETE)")
	fmt.Println("✓ Always set correct status codes")
	fmt.Println("✓ Close response bodies with defer")
	fmt.Println("✓ Set timeouts on HTTP clients")
	fmt.Println("✓ Parse URLs and query parameters correctly")
	fmt.Println("✓ Headers and cookies for metadata/state")
	fmt.Println("✓ REST design: resource-oriented with standard methods")
	fmt.Println("✓ API versioning and proper URL structure\n")
}
