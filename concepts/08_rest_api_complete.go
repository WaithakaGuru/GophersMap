package concepts

import (
	_ "encoding/json"
	"fmt"
	_ "io"
	_ "net/http"
	_ "strconv"
	_ "strings"
	"sync"
)

// ============================================================================
// COMPLETE REST API DESIGN AND IMPLEMENTATION
// ============================================================================
// This file covers:
// - Complete REST API structure
// - CRUD operations implementation
// - Request validation
// - Response formatting
// - Error handling in APIs
// - Real-world patterns
// ============================================================================

// Resource represents a sample resource (e.g., Blog Post)
type Resource struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// APIResponse wraps all API responses
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code"`
}

// ResourceStore holds resources in memory
type ResourceStore struct {
	mu        sync.RWMutex
	resources map[int]Resource
	nextID    int
}

// CompleteAPIStructureDemo shows the complete API structure
func CompleteAPIStructureDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║     COMPLETE REST API STRUCTURE DEMO                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. API ENDPOINTS DESIGN:")
	fmt.Println(`
// RESTful API Endpoints:

# Collections
GET    /api/v1/resources           - List all resources
POST   /api/v1/resources           - Create new resource

# Individual Resources
GET    /api/v1/resources/{id}      - Get specific resource
PUT    /api/v1/resources/{id}      - Replace entire resource
PATCH  /api/v1/resources/{id}      - Partial update
DELETE /api/v1/resources/{id}      - Delete resource

# Nested Resources
GET    /api/v1/resources/{id}/comments    - Comments on resource
POST   /api/v1/resources/{id}/comments    - Add comment

# Query Parameters
GET    /api/v1/resources?page=1&limit=10        - Pagination
GET    /api/v1/resources?author=john&sort=date - Filtering/Sorting
GET    /api/v1/resources?search=golang          - Search
`)

	fmt.Println("\n2. UNIFIED RESPONSE FORMAT:")
	fmt.Println(`
// ALL API responses follow same structure:

SUCCESS Response (200):
{
  "success": true,
  "message": "Resource retrieved successfully",
  "code": 200,
  "data": {
    "id": 1,
    "title": "Go Basics",
    "content": "...",
    "author": "John"
  }
}

ERROR Response (400):
{
  "success": false,
  "message": "Request failed",
  "code": 400,
  "error": "Invalid title: cannot be empty"
}

LIST Response (200):
{
  "success": true,
  "message": "Resources retrieved",
  "code": 200,
  "data": [
    { "id": 1, "title": "Post 1", ... },
    { "id": 2, "title": "Post 2", ... }
  ]
}
`)

	fmt.Println("✓ Complete API structure demonstrated")
}

// CRUDOperationsDemo demonstrates CRUD operations
func CRUDOperationsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       CRUD OPERATIONS IMPLEMENTATION DEMO              ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. CREATE (POST) OPERATION:")
	fmt.Println(`
func (s *ResourceStore) handleCreate(w http.ResponseWriter, r *http.Request) {
    var resource Resource
    
    // Decode JSON from request body
    if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
        s.sendError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }
    
    // Validate required fields
    if resource.Title == "" {
        s.sendError(w, http.StatusBadRequest, "Title is required")
        return
    }
    
    if resource.Author == "" {
        s.sendError(w, http.StatusBadRequest, "Author is required")
        return
    }
    
    // Generate ID and save
    s.mu.Lock()
    resource.ID = s.nextID
    s.nextID++
    resource.CreatedAt = time.Now().Format(time.RFC3339)
    resource.UpdatedAt = resource.CreatedAt
    s.resources[resource.ID] = resource
    s.mu.Unlock()
    
    // Send response with 201 Created status
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Location", fmt.Sprintf("/api/v1/resources/%d", resource.ID))
    s.sendSuccess(w, "Resource created", resource)
}
`)

	fmt.Println("\n2. READ (GET) OPERATIONS:")
	fmt.Println(`
// Get single resource
func (s *ResourceStore) handleGetOne(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/api/v1/resources/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        s.sendError(w, http.StatusBadRequest, "Invalid resource ID")
        return
    }
    
    s.mu.RLock()
    resource, exists := s.resources[id]
    s.mu.RUnlock()
    
    if !exists {
        s.sendError(w, http.StatusNotFound, "Resource not found")
        return
    }
    
    s.sendSuccess(w, "Resource retrieved", resource)
}

// Get all resources with pagination
func (s *ResourceStore) handleGetAll(w http.ResponseWriter, r *http.Request) {
    page := r.URL.Query().Get("page")
    limit := r.URL.Query().Get("limit")
    
    // Default values
    pageNum := 1
    pageLimit := 10
    
    if p, err := strconv.Atoi(page); err == nil && p > 0 {
        pageNum = p
    }
    if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 100 {
        pageLimit = l
    }
    
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    var resources []Resource
    for _, r := range s.resources {
        resources = append(resources, r)
    }
    
    // Implement pagination
    start := (pageNum - 1) * pageLimit
    end := start + pageLimit
    if end > len(resources) {
        end = len(resources)
    }
    
    if start >= len(resources) {
        resources = []Resource{}
    } else {
        resources = resources[start:end]
    }
    
    s.sendSuccess(w, "Resources retrieved", resources)
}
`)

	fmt.Println("\n3. UPDATE (PUT/PATCH) OPERATIONS:")
	fmt.Println(`
// PUT - Replace entire resource
func (s *ResourceStore) handlePUT(w http.ResponseWriter, r *http.Request) {
    idStr := extractID(r.URL.Path)
    id, _ := strconv.Atoi(idStr)
    
    var updated Resource
    json.NewDecoder(r.Body).Decode(&updated)
    
    // Validate
    if updated.Title == "" {
        s.sendError(w, http.StatusBadRequest, "Title required")
        return
    }
    
    s.mu.Lock()
    original, exists := s.resources[id]
    if exists {
        original.Title = updated.Title
        original.Content = updated.Content
        original.Author = updated.Author
        original.UpdatedAt = time.Now().Format(time.RFC3339)
        s.resources[id] = original
    }
    s.mu.Unlock()
    
    if !exists {
        s.sendError(w, http.StatusNotFound, "Resource not found")
        return
    }
    
    s.sendSuccess(w, "Resource updated", s.resources[id])
}

// PATCH - Partial update
func (s *ResourceStore) handlePATCH(w http.ResponseWriter, r *http.Request) {
    idStr := extractID(r.URL.Path)
    id, _ := strconv.Atoi(idStr)
    
    // Decode partial updates
    updates := map[string]interface{}{}
    json.NewDecoder(r.Body).Decode(&updates)
    
    s.mu.Lock()
    original, exists := s.resources[id]
    if exists {
        // Only update provided fields
        if title, ok := updates["title"].(string); ok {
            original.Title = title
        }
        if content, ok := updates["content"].(string); ok {
            original.Content = content
        }
        original.UpdatedAt = time.Now().Format(time.RFC3339)
        s.resources[id] = original
    }
    s.mu.Unlock()
    
    if !exists {
        s.sendError(w, http.StatusNotFound, "Resource not found")
        return
    }
    
    s.sendSuccess(w, "Resource updated", s.resources[id])
}
`)

	fmt.Println("\n4. DELETE OPERATION:")
	fmt.Println(`
func (s *ResourceStore) handleDELETE(w http.ResponseWriter, r *http.Request) {
    idStr := extractID(r.URL.Path)
    id, err := strconv.Atoi(idStr)
    if err != nil {
        s.sendError(w, http.StatusBadRequest, "Invalid ID")
        return
    }
    
    s.mu.Lock()
    _, exists := s.resources[id]
    if exists {
        delete(s.resources, id)
    }
    s.mu.Unlock()
    
    if !exists {
        s.sendError(w, http.StatusNotFound, "Resource not found")
        return
    }
    
    // 204 No Content - successful deletion
    w.WriteHeader(http.StatusNoContent)
}
`)

	fmt.Println("✓ CRUD operations demonstrated")
}

// RequestValidationDemo shows validation patterns
func RequestValidationDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      REQUEST VALIDATION PATTERNS DEMO                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. JSON SCHEMA VALIDATION:")
	fmt.Println(`
type CreateResourceRequest struct {
    Title   string   ` + "`" + `json:"title" validate:"required,min=3,max=100"` + "`" + `
    Content string   ` + "`" + `json:"content" validate:"required,min=10"` + "`" + `
    Author  string   ` + "`" + `json:"author" validate:"required"` + "`" + `
    Tags    []string ` + "`" + `json:"tags" validate:"required,min=1,max=5"` + "`" + `
}

func validateRequest(req CreateResourceRequest) error {
    if req.Title == "" {
        return fmt.Errorf("title is required")
    }
    if len(req.Title) < 3 {
        return fmt.Errorf("title must be at least 3 characters")
    }
    if len(req.Title) > 100 {
        return fmt.Errorf("title cannot exceed 100 characters")
    }
    
    if req.Content == "" {
        return fmt.Errorf("content is required")
    }
    if len(req.Content) < 10 {
        return fmt.Errorf("content must be at least 10 characters")
    }
    
    if len(req.Tags) == 0 {
        return fmt.Errorf("at least one tag is required")
    }
    if len(req.Tags) > 5 {
        return fmt.Errorf("maximum 5 tags allowed")
    }
    
    return nil
}
`)

	fmt.Println("\n2. QUERY PARAMETER VALIDATION:")
	fmt.Println(`
func validatePagination(r *http.Request) (page, limit int, err error) {
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
`)

	fmt.Println("✓ Request validation demonstrated")
}

// ResponseFormattingDemo shows response formatting
func ResponseFormattingDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      RESPONSE FORMATTING PATTERNS DEMO                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. CONSISTENT RESPONSE FORMAT:")
	fmt.Println(`
// Helper: Send success response
func (s *ResourceStore) sendSuccess(w http.ResponseWriter, msg string, data interface{}) {
    response := APIResponse{
        Success: true,
        Message: msg,
        Data:    data,
        Code:    http.StatusOK,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Helper: Send error response
func (s *ResourceStore) sendError(w http.ResponseWriter, statusCode int, errMsg string) {
    response := APIResponse{
        Success: false,
        Message: "Request failed",
        Error:   errMsg,
        Code:    statusCode,
    }
    
    w.WriteHeader(statusCode)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
`)

	fmt.Println("\n2. DIFFERENT STATUS CODE RESPONSES:")
	fmt.Println(`
// 200 OK - GET, POST update successful
w.WriteHeader(http.StatusOK)
sendSuccess(w, "Operation successful", data)

// 201 Created - Resource created
w.WriteHeader(http.StatusCreated)
w.Header().Set("Location", fmt.Sprintf("/api/v1/resources/%d", id))
sendSuccess(w, "Resource created", data)

// 204 No Content - DELETE successful (often no body)
w.WriteHeader(http.StatusNoContent)

// 400 Bad Request - Invalid input
sendError(w, http.StatusBadRequest, "Invalid request data")

// 401 Unauthorized - Authentication required
sendError(w, http.StatusUnauthorized, "Authentication token required")

// 403 Forbidden - Access denied
sendError(w, http.StatusForbidden, "Access denied to this resource")

// 404 Not Found - Resource doesn't exist
sendError(w, http.StatusNotFound, "Resource not found")

// 409 Conflict - Data conflict (e.g., duplicate)
sendError(w, http.StatusConflict, "Resource already exists")

// 500 Internal Server Error
sendError(w, http.StatusInternalServerError, "Server error")
`)

	fmt.Println("✓ Response formatting demonstrated")
}

// ErrorHandlingInAPIDemo shows error handling patterns
func ErrorHandlingInAPIDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      ERROR HANDLING IN APIs DEMO                       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. VALIDATION ERROR RESPONSES:")
	fmt.Println(`
type ValidationErrorResponse struct {
    Success bool                   ` + "`json:\"success\"`" + `
    Message string                 ` + "`json:\"message\"`" + `
    Code    int                    ` + "`json:\"code\"` " + `
    Errors  map[string]string      ` + "`json:\"errors\"` " + `
}

// When validation fails, return details:
func handleCreateWithValidation(w http.ResponseWriter, r *http.Request) {
    var resource Resource
    json.NewDecoder(r.Body).Decode(&resource)
    
    errors := make(map[string]string)
    if resource.Title == "" {
        errors["title"] = "Title is required"
    }
    if resource.Author == "" {
        errors["author"] = "Author is required"
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
    
    // Continue with valid data...
}
`)

	fmt.Println("\n2. STRUCTURED ERROR HANDLING:")
	fmt.Println(`
// Define specific error types
type APIError struct {
    Code    int    ` + "`json:\"code\"` " + `
    Message string ` + "`json:\"message\"` " + `
    Details string ` + "`json:\"details\"` " + `
}

// Custom error handler
func (s *ResourceStore) handleAPIError(w http.ResponseWriter, err error) {
    if err == nil {
        return
    }
    
    // Map errors to HTTP status codes
    switch err {
    case ErrNotFound:
        apiErr := APIError{
            Code:    http.StatusNotFound,
            Message: "Resource not found",
            Details: err.Error(),
        }
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(apiErr)
        
    case ErrValidation:
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
`)

	fmt.Println("✓ Error handling in APIs demonstrated")
}

// RealWorldAPIExampleDemo shows complete real API
func RealWorldAPIExampleDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      REAL WORLD API EXAMPLE DEMO                       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("COMPLETE API HANDLER STRUCTURE:")
	fmt.Println(`
type API struct {
    store *ResourceStore
}

// Main router - directs to appropriate handler
func (a *API) handleResources(w http.ResponseWriter, r *http.Request) {
    // Set common headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-API-Version", "1.0")
    
    // Route based on method and path
    if strings.HasPrefix(r.URL.Path, "/api/v1/resources/") && r.URL.Path != "/api/v1/resources/" {
        // Specific resource: /resources/{id}
        switch r.Method {
        case http.MethodGet:
            a.store.handleGetOne(w, r)
        case http.MethodPut:
            a.store.handlePUT(w, r)
        case http.MethodPatch:
            a.store.handlePATCH(w, r)
        case http.MethodDelete:
            a.store.handleDELETE(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    } else {
        // Collection: /resources
        switch r.Method {
        case http.MethodGet:
            a.store.handleGetAll(w, r)
        case http.MethodPost:
            a.store.handleCreate(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

// Start server
func (a *API) Start(port string) error {
    http.HandleFunc("/api/v1/resources", a.handleResources)
    http.HandleFunc("/api/v1/resources/", a.handleResources)
    
    fmt.Printf("API Server starting on port %s\n", port)
    return http.ListenAndServe(":"+port, nil)
}
`)

	fmt.Println("✓ Real world API example demonstrated")
}

// RunRESTAPIExamples executes all REST API demos
func RunRESTAPIExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    COMPLETE REST API DESIGN - COMPREHENSIVE GUIDE       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	CompleteAPIStructureDemo()
	CRUDOperationsDemo()
	RequestValidationDemo()
	ResponseFormattingDemo()
	ErrorHandlingInAPIDemo()
	RealWorldAPIExampleDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         ALL REST API EXAMPLES COMPLETE                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ Consistent API structure across all endpoints")
	fmt.Println("✓ Proper HTTP methods for CRUD operations")
	fmt.Println("✓ Always validate request data")
	fmt.Println("✓ Use correct HTTP status codes")
	fmt.Println("✓ Consistent response format (success/error)")
	fmt.Println("✓ Thread-safe data store (use mutexes)")
	fmt.Println("✓ Handle errors gracefully")
	fmt.Println("✓ API versioning (/v1/, /v2/)")
	fmt.Println("✓ Pagination support for large datasets")
	fmt.Println("✓ Comprehensive error messages\n")
}
