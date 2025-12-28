/*
AUTHENTICATION AND JWT (JSON WEB TOKENS) IN GO
============================================================================

Authentication is crucial for secure applications. This file covers:
- Password hashing with bcrypt (never store plain text!)
- JWT tokens for stateless authentication
- Token claims and expiration
- Refresh tokens for better security
- Integration with HTTP handlers
- Security best practices

Key principles:
- Never store plain text passwords
- Always hash passwords on server side
- Use HTTPS for all authentication
- Keep tokens short-lived
- Implement rate limiting for login attempts
*/

package concepts

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// ============================================================================
// 1. PASSWORD HASHING WITH BCRYPT
// ============================================================================

/*
BCRYPT:
- One-way hashing algorithm (cannot be reversed)
- Automatically includes salt (random, prevents rainbow tables)
- Slow by design (resistant to brute force attacks)
- Adaptive cost (can increase as computers get faster)
- Industry standard for password storage

Never store plain text passwords!
*/

// HashPassword creates a bcrypt hash of a password
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost = 10 (good balance of speed/security)
	// Higher cost = slower but more secure
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if password matches the hash
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	return err == nil
}

// ValidatePassword checks password strength requirements
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

// PasswordHashingDemo demonstrates bcrypt usage
func PasswordHashingDemo() {
	fmt.Println("\n========== PASSWORD HASHING WITH BCRYPT ==========")

	fmt.Println("\n1. HASHING A PASSWORD:")
	// Example of hashing a password
	password := "MySecurePassword123!"
	hash, err := HashPassword(password)
	if err != nil {
		fmt.Printf("Error hashing: %v\n", err)
		return
	}
	fmt.Printf("Original password: %s\n", password)
	fmt.Printf("Hashed (bcrypt):   %s\n", hash)

	fmt.Println("\n2. VERIFYING PASSWORD:")
	// Verify correct password
	isValid := VerifyPassword(hash, "MySecurePassword123!")
	fmt.Printf("Password matches: %v\n", isValid)

	// Verify wrong password
	isValid = VerifyPassword(hash, "WrongPassword")
	fmt.Printf("Wrong password matches: %v\n", isValid)

	fmt.Println("\n3. PASSWORD VALIDATION:")
	// Test password strength
	err = ValidatePassword("weak")
	fmt.Printf("Weak password check: %v\n", err)

	err = ValidatePassword("StrongPass123!")
	fmt.Printf("Strong password check: %v\n", err)

	fmt.Println("\n4. BCRYPT PROPERTIES:")
	fmt.Println("   - One-way function (cannot unhash)")
	fmt.Println("   - Includes random salt (prevents rainbow tables)")
	fmt.Println("   - Adaptive cost (increases with hardware)")
	fmt.Println("   - Slow by design (resists brute force)")

	fmt.Println("\n✓ Password hashing demonstrated")
}

// ============================================================================
// 2. JWT (JSON WEB TOKEN) CONCEPTS
// ============================================================================

/*
JWT STRUCTURE:
- Consists of three parts: header.payload.signature
- All parts are Base64 encoded

Header:
{
  "alg": "HS256",    // Signing algorithm
  "typ": "JWT"       // Token type
}

Payload (Claims):
{
  "sub": "user123",           // Subject (user ID)
  "name": "John Doe",         // Custom claim
  "iat": 1516239022,          // Issued at (Unix timestamp)
  "exp": 1516242622           // Expiration (Unix timestamp)
}

Signature:
HMACSHA256(base64(header).base64(payload), secret)

Benefits:
- Stateless (no server-side session storage)
- Self-contained (claims are in token)
- Scalable (works with microservices)
- Mobile-friendly (can be sent in headers or URLs)
- Cross-domain/CORS compatible
*/

// JWTClaims represents standard JWT claims
type JWTClaims struct {
	Subject   string // "sub" - User identifier
	IssuedAt  int64  // "iat" - When token was created
	ExpiresAt int64  // "exp" - When token expires
	NotBefore int64  // "nbf" - Before this time, invalid
	Issuer    string // "iss" - Who created the token
	Audience  string // "aud" - Who token is for

	// Custom claims
	Username string
	Email    string
	Role     string
}

// JWTBasicsDemo explains JWT concepts
func JWTBasicsDemo() {
	fmt.Println("\n========== JWT (JSON WEB TOKEN) BASICS ==========")

	fmt.Println("\n1. JWT STRUCTURE:")
	fmt.Println("   Format: header.payload.signature")
	fmt.Println("   Example:")
	fmt.Println("   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.")
	fmt.Println("   eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.")
	fmt.Println("   SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

	fmt.Println("\n2. JWT CLAIMS:")
	fmt.Println("   - sub:  Subject (user ID)")
	fmt.Println("   - iat:  Issued at (Unix timestamp)")
	fmt.Println("   - exp:  Expiration (Unix timestamp)")
	fmt.Println("   - nbf:  Not before (Unix timestamp)")
	fmt.Println("   - iss:  Issuer (who created token)")
	fmt.Println("   - aud:  Audience (who token is for)")
	fmt.Println("   - Custom claims: username, email, role, etc.")

	fmt.Println("\n3. JWT WORKFLOW:")
	fmt.Println("   1. User logs in with credentials")
	fmt.Println("   2. Server creates JWT token")
	fmt.Println("   3. Client stores token (cookie, localStorage, etc.)")
	fmt.Println("   4. Client sends token in Authorization header")
	fmt.Println("   5. Server verifies token signature and expiration")
	fmt.Println("   6. If valid, process request; if expired, reject")

	fmt.Println("\n4. ADVANTAGES:")
	fmt.Println("   ✓ Stateless (no session storage needed)")
	fmt.Println("   ✓ Scalable (works with multiple servers)")
	fmt.Println("   ✓ Self-contained (claims in token)")
	fmt.Println("   ✓ Mobile-friendly")
	fmt.Println("   ✓ Cross-domain/CORS compatible")

	fmt.Println("\n5. DISADVANTAGES:")
	fmt.Println("   ✗ Cannot be revoked until expiration")
	fmt.Println("   ✗ If compromised, attacker has access until expiration")
	fmt.Println("   ✗ Larger than session cookies")
	fmt.Println("   ✗ Cannot store sensitive data (only encoded)")
	fmt.Println("   ✗ Key rotation is complex")

	fmt.Println("\n✓ JWT basics demonstrated")
}

// ============================================================================
// 3. TOKEN PATTERNS
// ============================================================================

/*
ACCESS TOKEN vs REFRESH TOKEN:

Access Token:
- Short-lived (15 minutes)
- Used for API requests
- Has restricted permissions
- If compromised, limited damage

Refresh Token:
- Long-lived (7 days)
- Used only to get new access token
- Stored securely (HTTP-only cookie)
- Can be revoked

Flow:
1. User logs in → Get access token + refresh token
2. Use access token for API calls
3. When access token expires → Use refresh token to get new one
4. If refresh token expired → User must login again
*/

// TokenPair represents access and refresh tokens
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64 // seconds
}

// RefreshTokensDemo explains token refresh pattern
func RefreshTokensDemo() {
	fmt.Println("\n========== REFRESH TOKEN PATTERN ==========")

	fmt.Println("\n1. WHY REFRESH TOKENS?")
	fmt.Println("   - Security: Access tokens should be short-lived")
	fmt.Println("   - UX: Don't want users to login constantly")
	fmt.Println("   - Balance between security and convenience")

	fmt.Println("\n2. TOKEN TYPES AND LIFETIMES:")
	fmt.Println("   Access Token:  15 minutes")
	fmt.Println("   Refresh Token: 7 days")
	fmt.Println("   API Key:       Never expires (use carefully)")

	fmt.Println("\n3. LOGIN FLOW:")
	fmt.Println("   1. User sends username + password to /login")
	fmt.Println("   2. Server verifies credentials")
	fmt.Println("   3. Server creates access token (15 min expiry)")
	fmt.Println("   4. Server creates refresh token (7 day expiry)")
	fmt.Println("   5. Client stores both tokens")

	fmt.Println("\n4. ACCESSING API:")
	fmt.Println("   1. Client sends API request with access token in header")
	fmt.Println("   2. Server verifies token is valid and not expired")
	fmt.Println("   3. If valid: process request")
	fmt.Println("   4. If expired: return 401 Unauthorized")

	fmt.Println("\n5. TOKEN REFRESH:")
	fmt.Println("   1. Client detects access token expired (401)")
	fmt.Println("   2. Client sends refresh token to /refresh endpoint")
	fmt.Println("   3. Server verifies refresh token")
	fmt.Println("   4. Server creates new access token")
	fmt.Println("   5. Client retries original request with new token")

	fmt.Println("\n6. SECURE STORAGE:")
	fmt.Println("   Access Token:")
	fmt.Println("   - Can be in: localStorage (XSS vulnerable)")
	fmt.Println("   - Or in: memory (lost on refresh)")
	fmt.Println("")
	fmt.Println("   Refresh Token:")
	fmt.Println("   - MUST be in: HTTP-only cookie (XSS protected)")
	fmt.Println("   - MUST have: Secure flag (HTTPS only)")
	fmt.Println("   - MUST have: SameSite flag (CSRF protected)")

	fmt.Println("\n✓ Refresh token pattern demonstrated")
}

// ============================================================================
// 4. SECURITY BEST PRACTICES
// ============================================================================

/*
AUTHENTICATION SECURITY PRINCIPLES:

1. NEVER store plain text passwords - use bcrypt
2. ALWAYS use HTTPS (encrypt in transit)
3. Use short-lived access tokens (15-60 minutes)
4. Use HTTP-only cookies for sensitive tokens
5. Implement rate limiting on login endpoints
6. Log security events
7. Use strong password requirements
8. Implement account lockout after failed attempts
9. Force password reset on breach
10. Use multi-factor authentication (MFA)

Common Attacks:

XSS (Cross-Site Scripting):
- Attack: Inject malicious script into web page
- Victim: User's browser executes attacker's script
- Steals: Tokens from localStorage
- Defense: Use HTTP-only cookies, sanitize input, CSP headers

CSRF (Cross-Site Request Forgery):
- Attack: Make user perform unwanted action
- Method: User logs into Bank.com, then visits Evil.com
- Evil.com makes request to Bank.com in user's browser
- Defense: Use SameSite cookies, CSRF tokens, verify origin

Token Hijacking:
- Attack: Steal token from network or storage
- Defense: Use HTTPS, short-lived tokens, refresh tokens

Brute Force:
- Attack: Try many passwords/tokens
- Defense: Rate limit, bcrypt (slow), account lockout
*/

// SecurityBestPracticesDemo covers security patterns
func SecurityBestPracticesDemo() {
	fmt.Println("\n========== SECURITY BEST PRACTICES ==========")

	fmt.Println("\n1. PASSWORD SECURITY:")
	fmt.Println("   ✓ DO:")
	fmt.Println("     - Use bcrypt/scrypt/argon2 for hashing")
	fmt.Println("     - Require minimum 12 characters")
	fmt.Println("     - Require mixed case, numbers, special chars")
	fmt.Println("     - Hash on server side ONLY")
	fmt.Println("     - Never log passwords")
	fmt.Println("     - Implement rate limiting on login")
	fmt.Println("")
	fmt.Println("   ✗ DON'T:")
	fmt.Println("     - Store plain text passwords")
	fmt.Println("     - Use weak hashing (MD5, SHA1, SHA256)")
	fmt.Println("     - Use reversible encryption")
	fmt.Println("     - Send passwords via email")
	fmt.Println("     - Log authentication attempts")

	fmt.Println("\n2. TOKEN SECURITY:")
	fmt.Println("   ✓ DO:")
	fmt.Println("     - Use HTTPS only (encrypted)")
	fmt.Println("     - Keep access tokens short-lived (15 min)")
	fmt.Println("     - Keep refresh tokens in HTTP-only cookies")
	fmt.Println("     - Set Secure flag (HTTPS only)")
	fmt.Println("     - Set SameSite flag (prevent CSRF)")
	fmt.Println("     - Implement token rotation")
	fmt.Println("     - Verify token signature on every request")
	fmt.Println("")
	fmt.Println("   ✗ DON'T:")
	fmt.Println("     - Send tokens via URL query parameters")
	fmt.Println("     - Store sensitive data in JWT claims")
	fmt.Println("     - Use long-lived access tokens")
	fmt.Println("     - Send HTTP (unencrypted)")
	fmt.Println("     - Disable SSL/TLS verification")

	fmt.Println("\n3. PREVENTING COMMON ATTACKS:")
	fmt.Println("   XSS Prevention:")
	fmt.Println("   - Use HTTP-only cookies")
	fmt.Println("   - Sanitize all user input")
	fmt.Println("   - Set Content-Security-Policy headers")
	fmt.Println("")
	fmt.Println("   CSRF Prevention:")
	fmt.Println("   - Use SameSite=Lax cookie flag")
	fmt.Println("   - Verify Origin/Referer headers")
	fmt.Println("   - Use CSRF tokens for state-changing operations")
	fmt.Println("")
	fmt.Println("   Brute Force Prevention:")
	fmt.Println("   - Rate limit login attempts (5 fails = 15 min lockout)")
	fmt.Println("   - Use slow hashing (bcrypt)")
	fmt.Println("   - Log failed attempts")
	fmt.Println("   - Implement CAPTCHA after N failures")

	fmt.Println("\n✓ Security best practices demonstrated")
}

// ============================================================================
// 5. MAIN EXECUTION
// ============================================================================

// RunAuthenticationExamples executes all authentication demonstrations
func RunAuthenticationExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   AUTHENTICATION AND JWT - COMPREHENSIVE GUIDE          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	PasswordHashingDemo()
	JWTBasicsDemo()
	RefreshTokensDemo()
	SecurityBestPracticesDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL AUTHENTICATION EXAMPLES COMPLETE                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Println("\nKEY TAKEAWAYS:")
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
