package concepts

import (
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
	"fmt"
	"time"

    "golang.org/x/crypto/bcrypt"
)

// ============================================================================
// AUTHENTICATION AND JWT (JSON WEB TOKENS)
// ============================================================================
// This file covers:
// - Password hashing (bcrypt)
// - JWT creation and validation
// - Token claims
// - Refresh tokens
// - Integration with REST APIs
// - Security best practices
// ============================================================================

// PasswordHashingDemo demonstrates secure password handling
func PasswordHashingDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       PASSWORD HASHING DEMO                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. BCRYPT FOR PASSWORD HASHING:")
}
// NEVER store plain text passwords!
// Use bcrypt for hashing:

// Hash a password
func HashPassword(password string) (string, error) {
    // bcrypt.DefaultCost = 10 (good balance)
    // Higher cost = slower (better security but slower)
    hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// Verify password
func VerifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword(
        []byte(hashedPassword),
        []byte(password),
    )
    return err == nil
}

// Example usage:
func exampleBcryptUse () {
hash, _ := HashPassword("myPassword123")
// hash = "$2a$10$..." (hashed password)

// Later, verify on login:
    if VerifyPassword(hash, "myPassword123") {
        // Password matches!
    }
    fmt.Println(` Properties of bcrypt:`)
    fmt.Println(`
- One-way function (cannot unhash)
- Slow (resistant to brute force)
- Includes salt (random, prevents rainbow tables)
- Adaptive (cost can be increased as computers get faster) 
`)
}

func PasswordSecurityInfo() {
	fmt.Println("\n2. PASSWORD SECURITY RULES:")
	fmt.Println(`
// ✓ DO:
// - Use bcrypt/scrypt/argon2 for hashing
// - Never log passwords
// - Use HTTPS only (no HTTP)
// - Implement rate limiting for login attempts
// - Force password reset on breach
// - Use strong password requirements
// - Hash on server, never client-side only

// ✗ DON'T:
// - Store plain text passwords
// - Use simple hash (MD5, SHA1)
// - Use reversible encryption
// - Send passwords via email
// - Log authentication attempts
// - Create weak password 
`)
}
// Password requirements example:
func ValidatePassword(password string) error {
    if len(password) < 12 {
        return fmt.Errorf("password too short (min 12 chars)")
    }
    hasUpper := false
    hasLower := false
    hasDigit := false
    hasSpecial := false
    
    for _, r := range password {
        switch {
        case r >= 'A' && r <= 'Z':
            hasUpper = true
        case r >= 'a' && r <= 'z':
            hasLower = true
        case r >= '0' && r <= '9':
            hasDigit = true
        case r >= '!' && r <= '~':
            hasSpecial = true
        }
    }
    
    if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
        return fmt.Errorf("password must contain uppercase, lowercase, digit, special char")
    }
    return nil
}
`)

	fmt.Println("✓ Password hashing demonstrated")
}

// JWTBasicsDemo demonstrates JWT concepts
func JWTBasicsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         JWT BASICS DEMO                                ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. WHAT IS JWT:")
	fmt.Println(`
// JWT = JSON Web Token
// Format: header.payload.signature

// Structure:
{
  "alg": "HS256",    // Algorithm (header)
  "typ": "JWT"
}
.
{
  "sub": "user123",           // Subject (payload)
  "name": "John Doe",
  "email": "john@example.com",
  "iat": 1516239022,          // Issued at
  "exp": 1516242622           // Expiration
}
.
HMACSHA256(base64(header).base64(payload), secret)

// Example JWT:
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
// eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
// SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
`)

	fmt.Println("\n2. JWT CLAIMS:")
	fmt.Println(`
type Claims struct {
    // Standard claims
    Subject   string ` + "`json:\"sub\"`" + `  // User ID
    IssuedAt  int64  ` + "`json:\"iat\"`" + `  // When token was created
    ExpiresAt int64  ` + "`json:\"exp\"`" + `  // When token expires
    NotBefore int64  ` + "`json:\"nbf\"`" + `  // Before this time, invalid
    Issuer    string ` + "`json:\"iss\"`" + `  // Who created token
    Audience  string ` + "`json:\"aud\"`" + `  // Who token is for
    ID        string ` + "`json:\"jti\"`" + `  // Unique token ID
    
    // Custom claims
    Username  string ` + "`json:\"username\"`" + `
    Email     string ` + "`json:\"email\"`" + `
    Role      string ` + "`json:\"role\"`" + `
}

// Standard claim times are Unix timestamps
claims := Claims{
    Subject:   "user123",
    Username:  "john",
    Email:     "john@example.com",
    Role:      "admin",
    IssuedAt:  time.Now().Unix(),
    ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
}
`)

	fmt.Println("\n3. JWT ADVANTAGES AND DISADVANTAGES:")
	fmt.Println(`
// Advantages:
// ✓ Stateless (no server-side session storage)
// ✓ Scalable (works great with microservices)
// ✓ Mobile-friendly (can send in headers or URLs)
// ✓ Self-contained (claims are in token)
// ✓ Cross-domain/CORS (works with any domain)
// ✓ Can be encrypted (JWE)

// Disadvantages:
// ✗ Token cannot be revoked until expiration
// ✗ If compromised, attacker has access until expiration
// ✗ Larger than session cookies
// ✗ Cannot store sensitive data (it's encoded, not encrypted)
// ✗ Key rotation is complex
`)

	fmt.Println("✓ JWT basics demonstrated")
}

// JWTImplementationDemo shows JWT usage patterns
func JWTImplementationDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         JWT IMPLEMENTATION DEMO                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. CREATE JWT TOKEN:")
	fmt.Println(`
import "github.com/golang-jwt/jwt/v4"

// Create token
func CreateToken(userID string, username string) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "username": username,
        "email": username + "@example.com",
        "role": "user",
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Sign with secret key
    secretKey := "your-secret-key-keep-it-safe"
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}
`)

	fmt.Println("\n2. VERIFY JWT TOKEN:")
	fmt.Println(`
func VerifyToken(tokenString string) (*jwt.Claims, error) {
    secretKey := "your-secret-key-keep-it-safe"
    
    claims := &jwt.StandardClaims{}
    token, err := jwt.ParseWithClaims(
        tokenString,
        claims,
        func(token *jwt.Token) (interface{}, error) {
            // Verify signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(secretKey), nil
        },
    )
    
    if err != nil {
        return nil, fmt.Errorf("token parsing failed: %v", err)
    }
    
    if !token.Valid {
        return nil, fmt.Errorf("token is not valid")
    }
    
    return claims, nil
}
`)

	fmt.Println("\n3. USE IN HTTP MIDDLEWARE:")
	fmt.Println(`
func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Get token from Authorization header
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Missing authorization header", http.StatusUnauthorized)
                return
            }
            
            // Expected format: "Bearer token123..."
            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 || parts[0] != "Bearer" {
                http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
                return
            }
            
            tokenString := parts[1]
            
            // Verify token
            claims, err := VerifyToken(tokenString)
            if err != nil {
                http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
                return
            }
            
            // Check expiration
            if claims.ExpiresAt < time.Now().Unix() {
                http.Error(w, "Token expired", http.StatusUnauthorized)
                return
            }
            
            // Store claims in request context
            ctx := context.WithValue(r.Context(), "claims", claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// Usage in handler:
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    claims := r.Context().Value("claims").(*jwt.StandardClaims)
    fmt.Fprintf(w, "Hello %s", claims.Subject)
}
`)

	fmt.Println("✓ JWT implementation demonstrated")
}

// RefreshTokensDemo shows token refresh pattern
func RefreshTokensDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      REFRESH TOKENS DEMO                               ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. REFRESH TOKEN PATTERN:")
	fmt.Println(`
    // Problem: Access tokens should be short-lived (security)
    // But constantly asking user to login is annoying (UX)
    // Solution: Refresh tokens!

    // Token types:
    // - Access Token: Short-lived (15 min), for API requests
    // - Refresh Token: Long-lived (7 days), for getting new access token

    type TokenPair struct {
        AccessToken  string  ` + "`json:\"access_token\"`" + ` 
        RefreshToken string ` + "`json:\"refresh_token\"`" + `
        ExpiresIn    int64  ` + "`json:\"expires_in\"`" + `
        ff int ` + "`json:\"jgjg\"`" + `
    }
}

func CreateTokenPair(userID string) (*TokenPair, error) {
    // Create short-lived access token (15 minutes)
    accessToken, err := CreateToken(userID, "access", 15*time.Minute)
    if err != nil {
        return nil, err
    }
    
    // Create long-lived refresh token (7 days)
    refreshToken, err := CreateToken(userID, "refresh", 7*24*time.Hour)
    if err != nil {
        return nil, err
    }
    
    return &TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ExpiresIn:    15 * 60, // 15 minutes in seconds
    }, nil
}

// Login endpoint returns both tokens:
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Verify credentials
    userID := "user123"
    
    tokens, _ := CreateTokenPair(userID)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tokens)
}
`)

	fmt.Println("\n2. REFRESH TOKEN ENDPOINT:")
	fmt.Println(`
// Client sends refresh token to get new access token
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RefreshToken string ` + "`json:\"refresh_token\"`" + `
    }
    json.NewDecoder(r.Body).Decode(&req)
    
    // Verify refresh token
    claims, err := VerifyToken(req.RefreshToken)
    if err != nil {
        http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
        return
    }
    
    // Check token type
    if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
        http.Error(w, "Not a refresh token", http.StatusUnauthorized)
        return
    }
    
    // Create new access token
    userID := claims["sub"].(string)
    newAccessToken, _ := CreateToken(userID, "access", 15*time.Minute)
    
    response := map[string]string{
        "access_token": newAccessToken,
        "token_type":   "Bearer",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Refresh token should be stored securely on client:
// - HTTP-only cookie (best, protected from XSS)
// - Local storage (vulnerable to XSS, but simpler)
// - Memory (lost on page refresh)
`)

	fmt.Println("✓ Refresh tokens demonstrated")
}

// SecurityBestPracticesDemo shows security patterns
func SecurityBestPracticesDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      AUTHENTICATION SECURITY DEMO                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. SECURE TOKEN STORAGE:")
	fmt.Println(`
// Browser storage options:

// 1. HTTP-only Cookies (RECOMMENDED)
// Pros:
// - Automatic, sent with every request
// - Protected from XSS (JavaScript cannot access)
// - Protected from CSRF (with SameSite flag)
// Cons:
// - CSRF possible if not configured correctly
// - Cannot send to different domain

// Implementation:
cookie := &http.Cookie{
    Name:     "access_token",
    Value:    tokenString,
    HttpOnly: true,  // ← CRITICAL: Prevents JavaScript access
    Secure:   true,  // ← HTTPS only
    SameSite: http.SameSiteLaxMode,  // ← Prevents CSRF
    Path:     "/",
    MaxAge:   15 * 60,  // 15 minutes
}
http.SetCookie(w, cookie)

// 2. Local Storage
// Pros:
// - Simple to implement
// - Accessible from JavaScript
// Cons:
// - Vulnerable to XSS attacks
// - Must manually send in requests

// 3. Session Storage
// Pros:
// - Cleared on browser close
// - Vulnerable to XSS
// Cons:
// - Lost on page refresh
`)

	fmt.Println("\n2. COMMON ATTACKS AND DEFENSES:")
	fmt.Println(`
// XSS (Cross-Site Scripting)
// Attack: Inject malicious script into page
// Defense:
// - Use HTTP-only cookies
// - Sanitize all user input
// - Use Content Security Policy headers
// - Escape output in templates

// CSRF (Cross-Site Request Forgery)
// Attack: Make user perform unwanted action
// Defense:
// - Use SameSite cookie flag
// - Use CSRF tokens
// - Verify Origin/Referer headers
// - Use POST for state-changing operations

// Token Hijacking
// Attack: Steal token from storage
// Defense:
// - Use HTTPS (encrypt in transit)
// - Use short-lived tokens
// - Use refresh tokens
// - Implement token rotation

// Brute Force
// Attack: Try many passwords
// Defense:
// - Rate limit login attempts
// - Implement account lockout
// - Use slow hashing (bcrypt)
// - Require strong passwords
`)

	fmt.Println("✓ Security best practices demonstrated")
}

// RealWorldAuthExampleDemo shows complete auth system
func RealWorldAuthExampleDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    REAL WORLD AUTH EXAMPLE DEMO                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("COMPLETE AUTHENTICATION SYSTEM:")
	fmt.Println(`
// File: auth.go
package auth

type User struct {
    ID       string
    Username string
    Email    string
    Password string // This is the HASH, not plain text
    Role     string
}

type AuthService struct {
    secretKey string
    users     map[string]*User
}

// Register user
func (s *AuthService) Register(username, email, password string) error {
    // Validate input
    if len(password) < 12 {
        return fmt.Errorf("password too short")
    }
    
    // Hash password
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }
    
    // Create user
    user := &User{
        ID:       generateID(),
        Username: username,
        Email:    email,
        Password: hashedPassword,
        Role:     "user",
    }
    
    s.users[user.ID] = user
    return nil
}

// Login and get tokens
func (s *AuthService) Login(username, password string) (*TokenPair, error) {
    // Find user
    var user *User
    for _, u := range s.users {
        if u.Username == username {
            user = u
            break
        }
    }
    
    if user == nil {
        return nil, fmt.Errorf("user not found")
    }
    
    // Verify password
    if !VerifyPassword(user.Password, password) {
        return nil, fmt.Errorf("invalid password")
    }
    
    // Create token pair
    return CreateTokenPair(user.ID)
}

// Middleware handler
func (s *AuthService) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }
        
        // Remove "Bearer " prefix
        token = strings.TrimPrefix(token, "Bearer ")
        
        claims, err := VerifyToken(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        // Add claims to context
        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
`)

	fmt.Println("✓ Real world auth example demonstrated")
}

// RunAuthenticationExamples executes all authentication demos
func RunAuthenticationExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   AUTHENTICATION AND JWT - COMPREHENSIVE GUIDE          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	PasswordHashingDemo()
	JWTBasicsDemo()
	JWTImplementationDemo()
	RefreshTokensDemo()
	SecurityBestPracticesDemo()
	RealWorldAuthExampleDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL AUTHENTICATION EXAMPLES COMPLETE                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ Never store plain text passwords (use bcrypt)")
	fmt.Println("✓ JWT tokens are stateless and scalable")
	fmt.Println("✓ Access tokens should be short-lived (15 min)")
	fmt.Println("✓ Use refresh tokens for long-term access")
	fmt.Println("✓ Use HTTP-only, Secure, SameSite cookies")
	fmt.Println("✓ Implement rate limiting for login attempts")
	fmt.Println("✓ Always use HTTPS for authentication")
	fmt.Println("✓ Validate tokens on every protected request")
	fmt.Println("✓ Handle token expiration gracefully")
	fmt.Println("✓ Protect against XSS, CSRF, and token hijacking\n")
}
