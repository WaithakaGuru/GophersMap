/*
COMPLETE REST API DESIGN AND IMPLEMENTATION IN GO
============================================================================

Building production-ready REST APIs requires careful design:
- Consistent request/response structures
- Proper HTTP status codes
- Input validation
- Error handling
- Thread-safe data access
- Resource management

TOPICS COVERED:
- Complete REST API architecture
- CRUD operations implementation
- Request and response validation
- Error handling patterns
- Real-world API design
*/

package concepts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Resource represents a blog post or article
type Resource struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// APIResponse wraps all API responses with consistent format
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code"`
}

// ValidationErrorResponse for detailed validation errors
type ValidationErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Errors  map[string]string `json:"errors"`
}

// APIError for structured error responses
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// ResourceStore holds resources in memory with thread-safe access
type ResourceStore struct {
	mu        sync.RWMutex
	resources map[int]Resource
	nextID    int
}


// ============================================================================
// 1. COMPLETE REST API STRUCTURE
// ============================================================================

/*
API DESIGN PRINCIPLES:
- Resource-oriented URLs
- Standard HTTP methods for operations
- Consistent response format
- Proper status codes
- Clear error messages
- Pagination support
- API versioning

URL Patterns:
- GET    /api/v1/resources           - List all
- POST   /api/v1/resources           - Create
- GET    /api/v1/resources/{id}      - Get one
- PUT    /api/v1/resources/{id}      - Replace
- PATCH  /api/v1/resources/{id}      - Partial update
- DELETE /api/v1/resources/{id}      - Delete
*/

func CompleteAPIStructureDemo() {
	fmt.Println("\n========== COMPLETE REST API STRUCTURE ==========")

	fmt.Println("\n1. API ENDPOINT PATTERNS:")
	fmt.Println("   Collections:")
	fmt.Println("      GET    /api/v1/resources           - List all resources")
	fmt.Println("      POST   /api/v1/resources           - Create new resource")
	fmt.Println("\n   Individual Resources:")
	fmt.Println("      GET    /api/v1/resources/{id}      - Get specific resource")
	fmt.Println("      PUT    /api/v1/resources/{id}      - Replace entire resource")
	fmt.Println("      PATCH  /api/v1/resources/{id}      - Partial update")
	fmt.Println("      DELETE /api/v1/resources/{id}      - Delete resource")

	fmt.Println("\n2. UNIFIED RESPONSE FORMAT:")
	fmt.Println("   All responses follow same structure:")
	fmt.Println("   {")
	fmt.Println("     \"success\": true/false,")
	fmt.Println("     \"message\": \"Human readable message\",")
	fmt.Println("     \"code\": 200,                     (HTTP status code)")
	fmt.Println("     \"data\": {...},                   (if successful)")
	fmt.Println("     \"error\": \"Error details\"       (if failed)")
	fmt.Println("   }")

	fmt.Println("\n3. QUERY PARAMETERS:")
	fmt.Println("   - Pagination: ?page=1&limit=10")
	fmt.Println("   - Filtering: ?author=john&status=active")
	fmt.Println("   - Sorting: ?sort=date&order=desc")
	fmt.Println("   - Search: ?search=golang")

	fmt.Println("\n✓ Complete API structure demonstrated")
}

// ============================================================================
// 2. CRUD OPERATIONS
// ============================================================================

/*
CRUD = Create, Read, Update, Delete

CREATE (POST):
- Receive JSON data
- Validate input
- Store resource
- Return 201 Created with Location header

READ (GET):
- Retrieve one or many resources
- Support filtering and pagination
- Return 200 OK with data

UPDATE (PUT/PATCH):
- PUT: Replace entire resource
- PATCH: Partial update
- Validate changes
- Return 200 OK with updated data

DELETE:
- Remove resource
- Return 204 No Content
*/

func CRUDOperationsDemo() {
	fmt.Println("\n========== CRUD OPERATIONS ==========")

	// Example CREATE operation
	create := func(store *ResourceStore, w http.ResponseWriter, r *http.Request) {
		var resource Resource

		// Decode JSON from request body
		if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Validate required fields
		if resource.Title == "" || resource.Author == "" {
			http.Error(w, "Title and Author are required", http.StatusBadRequest)
			return
		}

		// Generate ID and timestamps
		store.mu.Lock()
		resource.ID = store.nextID
		store.nextID++
		resource.CreatedAt = time.Now().Format(time.RFC3339)
		resource.UpdatedAt = resource.CreatedAt
		store.resources[resource.ID] = resource
		store.mu.Unlock()

		// Send 201 Created response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Location", fmt.Sprintf("/api/v1/resources/%d", resource.ID))
		json.NewEncoder(w).Encode(resource)
	}

	// Example READ single resource
	readOne := func(store *ResourceStore, w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/api/v1/resources/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid resource ID", http.StatusBadRequest)
			return
		}

		store.mu.RLock()
		resource, exists := store.resources[id]
		store.mu.RUnlock()

		if !exists {
			http.Error(w, "Resource not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resource)
	}

	// Example UPDATE (PUT - replace entire resource)
	updateFull := func(store *ResourceStore, w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/api/v1/resources/"):]
		id, _ := strconv.Atoi(idStr)

		var updated Resource
		json.NewDecoder(r.Body).Decode(&updated)

		// Validate required fields
		if updated.Title == "" {
			http.Error(w, "Title required", http.StatusBadRequest)
			return
		}

		store.mu.Lock()
		if resource, exists := store.resources[id]; exists {
			resource.Title = updated.Title
			resource.Content = updated.Content
			resource.Author = updated.Author
			resource.UpdatedAt = time.Now().Format(time.RFC3339)
			store.resources[id] = resource
		}
		store.mu.Unlock()
	}

	// Example UPDATE (PATCH - partial update)
	updatePartial := func(store *ResourceStore, w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/api/v1/resources/"):]
		id, _ := strconv.Atoi(idStr)

		updates := map[string]interface{}{}
		json.NewDecoder(r.Body).Decode(&updates)

		store.mu.Lock()
		if resource, exists := store.resources[id]; exists {
			// Only update provided fields
			if title, ok := updates["title"].(string); ok {
				resource.Title = title
			}
			if content, ok := updates["content"].(string); ok {
				resource.Content = content
			}
			resource.UpdatedAt = time.Now().Format(time.RFC3339)
			store.resources[id] = resource
		}
		store.mu.Unlock()
	}

	// Example DELETE
	delete_ := func(store *ResourceStore, w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/api/v1/resources/"):]
		id, _ := strconv.Atoi(idStr)

		store.mu.Lock()
		_, exists := store.resources[id]
		if exists {
			delete(store.resources, id)
		}
		store.mu.Unlock()

		if !exists {
			http.Error(w, "Resource not found", http.StatusNotFound)
			return
		}

		// 204 No Content - successful deletion, no response body
		w.WriteHeader(http.StatusNoContent)
	}

	fmt.Println("\n1. CREATE (POST):")
	fmt.Println("   - Parse and validate JSON")
	fmt.Println("   - Generate ID and timestamps")
	fmt.Println("   - Return 201 Created with Location header")
	_ = create

	fmt.Println("\n2. READ (GET):")
	fmt.Println("   - Retrieve by ID")
	fmt.Println("   - Return 200 OK with resource")
	_ = readOne

	fmt.Println("\n3. UPDATE - PUT (full replacement):")
	fmt.Println("   - Validate all required fields")
	fmt.Println("   - Update timestamp")
	fmt.Println("   - Return 200 OK with updated resource")
	_ = updateFull

	fmt.Println("\n4. UPDATE - PATCH (partial):")
	fmt.Println("   - Update only provided fields")
	fmt.Println("   - Keep existing values for others")
	fmt.Println("   - Return 200 OK")
	_ = updatePartial

	fmt.Println("\n5. DELETE:")
	fmt.Println("   - Remove resource")
	fmt.Println("   - Return 204 No Content")
	_ = delete_

	fmt.Println("\n✓ CRUD operations demonstrated")
}

// ============================================================================
// 3. REQUEST VALIDATION
// ============================================================================

/*
INPUT VALIDATION:
- Validate before processing
- Check required fields
- Validate data types
- Check value ranges
- Return detailed error messages

Validation function example:
- Check field presence
- Check field length
- Check value ranges
- Custom business logic validation
*/

func RequestValidationDemo() {
	fmt.Println("\n========== REQUEST VALIDATION ==========")

	// Validation function
	validateRequest := func(req Resource) error {
		if req.Title == "" {
			return fmt.Errorf("title is required")
		}
		if len(req.Title) < 3 || len(req.Title) > 100 {
			return fmt.Errorf("title must be 3-100 characters")
		}

		if req.Content == "" {
			return fmt.Errorf("content is required")
		}
		if len(req.Content) < 10 {
			return fmt.Errorf("content must be at least 10 characters")
		}

		if len(req.Tags) == 0 || len(req.Tags) > 5 {
			return fmt.Errorf("1-5 tags required")
		}

		return nil
	}

	// Pagination validation
	validatePagination := func(r *http.Request) (page, limit int, err error) {
		page = 1
		limit = 10

		if p := r.URL.Query().Get("page"); p != "" {
			pageNum, err := strconv.Atoi(p)
			if err != nil || pageNum < 1 {
				return 0, 0, fmt.Errorf("page must be positive integer")
			}
			page = pageNum
		}

		if l := r.URL.Query().Get("limit"); l != "" {
			limitNum, err := strconv.Atoi(l)
			if err != nil || limitNum < 1 || limitNum > 100 {
				return 0, 0, fmt.Errorf("limit must be 1-100")
			}
			limit = limitNum
		}

		return page, limit, nil
	}

	fmt.Println("\n1. JSON VALIDATION:")
	fmt.Println("   - Parse JSON from request body")
	fmt.Println("   - Check required fields")
	fmt.Println("   - Validate field lengths")
	fmt.Println("   - Return 400 Bad Request if invalid")
	_ = validateRequest

	fmt.Println("\n2. QUERY PARAMETER VALIDATION:")
	fmt.Println("   - Parse page and limit parameters")
	fmt.Println("   - Set defaults (page=1, limit=10)")
	fmt.Println("   - Enforce limits (0 < limit <= 100)")
	_ = validatePagination

	fmt.Println("\n3. VALIDATION ERROR RESPONSE:")
	fmt.Println("   {")
	fmt.Println("     \"success\": false,")
	fmt.Println("     \"message\": \"Validation failed\",")
	fmt.Println("     \"code\": 400,")
	fmt.Println("     \"errors\": {")
	fmt.Println("       \"title\": \"Title must be 3-100 characters\",")
	fmt.Println("       \"tags\": \"1-5 tags required\"")
	fmt.Println("     }")
	fmt.Println("   }")

	fmt.Println("\n✓ Request validation demonstrated")
}

// ============================================================================
// 4. RESPONSE FORMATTING AND HTTP STATUS CODES
// ============================================================================

/*
HTTP STATUS CODES:
- 2xx: Success (200 OK, 201 Created, 204 No Content)
- 3xx: Redirect (301 Moved Permanently, 302 Found)
- 4xx: Client Error (400 Bad Request, 404 Not Found, 409 Conflict)
- 5xx: Server Error (500 Internal Server Error, 503 Service Unavailable)

Status Code Usage in REST API:
- 200: GET successful, PUT/PATCH successful
- 201: POST successful (resource created)
- 204: DELETE successful (no content to return)
- 400: Invalid request data
- 401: Authentication required
- 403: Access forbidden
- 404: Resource not found
- 409: Conflict (duplicate, constraint violation)
- 500: Server error
*/

func ResponseFormattingDemo() {
	fmt.Println("\n========== RESPONSE FORMATTING ==========")

	// Helper functions for consistent responses
	sendSuccess := func(w http.ResponseWriter, msg string, data interface{}, code int) {
		response := APIResponse{
			Success: true,
			Message: msg,
			Data:    data,
			Code:    code,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(response)
	}

	sendError := func(w http.ResponseWriter, statusCode int, errMsg string) {
		response := APIResponse{
			Success: false,
			Message: "Request failed",
			Error:   errMsg,
			Code:    statusCode,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
	}

	fmt.Println("\n1. SUCCESS RESPONSES:")
	fmt.Println("   200 OK (GET, PUT, PATCH):")
	fmt.Println("      { \"success\": true, \"code\": 200, \"data\": {...} }")
	fmt.Println("\n   201 Created (POST):")
	fmt.Println("      { \"success\": true, \"code\": 201, \"data\": {...} }")
	fmt.Println("      + Location header: /api/v1/resources/123")
	fmt.Println("\n   204 No Content (DELETE):")
	fmt.Println("      HTTP 204 (empty body)")
	_ = sendSuccess

	fmt.Println("\n2. ERROR RESPONSES:")
	fmt.Println("   400 Bad Request: Invalid input data")
	fmt.Println("   401 Unauthorized: Authentication required")
	fmt.Println("   403 Forbidden: Access denied")
	fmt.Println("   404 Not Found: Resource doesn't exist")
	fmt.Println("   409 Conflict: Duplicate or constraint violation")
	fmt.Println("   500 Internal Server Error: Server error")
	_ = sendError

	fmt.Println("\n3. CONSISTENT STRUCTURE:")
	fmt.Println("   All errors return JSON with:")
	fmt.Println("      - success: false")
	fmt.Println("      - message: brief description")
	fmt.Println("      - error: detailed error message")
	fmt.Println("      - code: HTTP status code")

	fmt.Println("\n✓ Response formatting demonstrated")
}

// ============================================================================
// 5. ERROR HANDLING IN APIs
// ============================================================================

/*
ERROR HANDLING STRATEGIES:
- Validate input before processing
- Check resource existence
- Handle edge cases
- Return meaningful error messages
- Include error context
- Log errors for debugging

Error types:
- Validation errors (400)
- Not found errors (404)
- Conflict/duplicate errors (409)
- Authorization errors (401, 403)
- Server errors (500)
*/

func ErrorHandlingInAPIDemo() {
	fmt.Println("\n========== ERROR HANDLING IN APIs ==========")

	// Handle with validation errors
	handleCreateWithValidation := func(w http.ResponseWriter, r *http.Request) {
		var resource Resource
		if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Collect validation errors
		errors := make(map[string]string)
		if resource.Title == "" {
			errors["title"] = "Title is required"
		}
		if resource.Author == "" {
			errors["author"] = "Author is required"
		}
		if len(resource.Tags) == 0 {
			errors["tags"] = "At least one tag required"
		}

		if len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ValidationErrorResponse{
				Success: false,
				Message: "Validation failed",
				Code:    http.StatusBadRequest,
				Errors:  errors,
			})
			return
		}

		// Continue with processing if valid...
	}

	// Handle specific error types
	handleAPIError := func(w http.ResponseWriter, err error) {
		if err == nil {
			return
		}

		// Map error types to status codes
		switch err.Error() {
		case "ErrNotFound":
			apiErr := APIError{
				Code:    http.StatusNotFound,
				Message: "Resource not found",
				Details: err.Error(),
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apiErr)

		case "ErrValidation":
			apiErr := APIError{
				Code:    http.StatusBadRequest,
				Message: "Validation failed",
				Details: err.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apiErr)

		default:
			apiErr := APIError{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
				Details: err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apiErr)
		}
	}

	fmt.Println("\n1. VALIDATION ERROR RESPONSES:")
	fmt.Println("   Return detailed validation errors:")
	fmt.Println("   { \"errors\": { \"field\": \"error message\" } }")
	_ = handleCreateWithValidation

	fmt.Println("\n2. SPECIFIC ERROR HANDLING:")
	fmt.Println("   - Map custom errors to HTTP status codes")
	fmt.Println("   - Include error details in response")
	fmt.Println("   - Log errors for debugging")
	_ = handleAPIError

	fmt.Println("\n3. ERROR PROPAGATION:")
	fmt.Println("   - Catch errors early")
	fmt.Println("   - Don't expose internal details")
	fmt.Println("   - Return user-friendly messages")

	fmt.Println("\n✓ Error handling in APIs demonstrated")
}

// ============================================================================
// 6. REAL-WORLD API IMPLEMENTATION
// ============================================================================

/*
COMPLETE API STRUCTURE:
- Resource types (structs with JSON tags)
- Thread-safe data store (sync.RWMutex)
- Validation functions
- Handler functions (one per operation)
- Router/multiplexer
- Server startup
- Consistent error handling
*/

func RealWorldAPIExampleDemo() {
	fmt.Println("\n========== REAL-WORLD API EXAMPLE ==========")

	// API struct holds dependencies
	api := struct{}{} // Placeholder

	fmt.Println("\n1. API STRUCTURE:")
	fmt.Println("   type API struct {")
	fmt.Println("       store *ResourceStore")
	fmt.Println("   }")

	fmt.Println("\n2. MAIN ROUTER HANDLER:")
	fmt.Println("   - Route based on HTTP method")
	fmt.Println("   - Route based on path pattern")
	fmt.Println("   - Set common headers (Content-Type, CORS, etc.)")
	fmt.Println("   - Delegate to specific handlers")

	fmt.Println("\n3. HANDLER RESPONSIBILITIES:")
	fmt.Println("   - Parse and validate input")
	fmt.Println("   - Check resource existence")
	fmt.Println("   - Perform operation")
	fmt.Println("   - Return appropriate status code")
	fmt.Println("   - Send consistent response format")

	fmt.Println("\n4. THREAD SAFETY:")
	fmt.Println("   - Use sync.RWMutex for concurrent access")
	fmt.Println("   - Lock before reading/writing")
	fmt.Println("   - Defer unlock() to prevent deadlocks")
	fmt.Println("   - Use RLock() for read-only operations")

	fmt.Println("\n5. STARTING SERVER:")
	fmt.Println("   func (a *API) Start(port string) error {")
	fmt.Println("       http.HandleFunc(\"/api/v1/resources\", a.handleResources)")
	fmt.Println("       return http.ListenAndServe(\":\" + port, nil)")
	fmt.Println("   }")

	_ = api

	fmt.Println("\n✓ Real-world API example demonstrated")
}



// ============================================================================
// MAIN EXECUTION
// ============================================================================

// RunRESTAPIExamples executes all REST API demonstrations
func RunRESTAPIExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    COMPLETE REST API DESIGN - COMPREHENSIVE GUIDE      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	CompleteAPIStructureDemo()
	CRUDOperationsDemo()
	RequestValidationDemo()
	ResponseFormattingDemo()
	ErrorHandlingInAPIDemo()
	RealWorldAPIExampleDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         ALL REST API EXAMPLES COMPLETE                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Println("\nKEY TAKEAWAYS:")
	fmt.Println("✓ Resource-oriented URL design")
	fmt.Println("✓ Proper HTTP methods for CRUD operations")
	fmt.Println("✓ Consistent request/response format")
	fmt.Println("✓ Input validation before processing")
	fmt.Println("✓ Appropriate HTTP status codes")
	fmt.Println("✓ Detailed error messages")
	fmt.Println("✓ Thread-safe data access with mutexes")
	fmt.Println("✓ Pagination for large datasets")
	fmt.Println("✓ API versioning in URLs")
	fmt.Println("✓ Handler organization and routing\n")
}
