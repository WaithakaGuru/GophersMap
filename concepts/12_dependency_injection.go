package concepts

import (
	"fmt"
)

// ============================================================================
// DEPENDENCY INJECTION IN GO
// ============================================================================
// This file covers:
// - Dependency Injection concepts
// - Constructor injection pattern
// - Interface-based DI
// - Service locator pattern
// - Testing with DI
// - Real-world DI examples
// ============================================================================

// DIConceptsDemo shows DI fundamentals
func DIConceptsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       DEPENDENCY INJECTION CONCEPTS DEMO               ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. WHAT IS DEPENDENCY INJECTION:")
	fmt.Println(`
// Dependency = Something a component needs to work
// Injection = Providing that dependency from outside

// WITHOUT DI (TIGHTLY COUPLED):
type UserService struct {
    // Creates its own database - PROBLEM!
    // Hard to test, can't swap implementations
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Must use THIS database implementation
    db := NewRealDatabase()
    return db.GetUser(id)
}

// With DI (LOOSELY COUPLED):
type UserService struct {
    db UserRepository  // Injected dependency
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Can use any UserRepository implementation
    return s.db.GetUser(id)
}

// Benefits:
// ✓ Easy to test (use mock implementations)
// ✓ Flexible (swap implementations)
// ✓ Maintainable (changes in one place)
// ✓ Decoupled (components don't know about each other)
`)

	fmt.Println("\n2. DEPENDENCY TYPES:")
	fmt.Println(`
// 1. Constructor Injection (most common in Go)
type Service struct {
    logger Logger
    db     Database
}

func NewService(logger Logger, db Database) *Service {
    return &Service{
        logger: logger,
        db:     db,
    }
}

// 2. Setter Injection
func (s *Service) SetLogger(logger Logger) {
    s.logger = logger
}

// 3. Interface Injection
type LoggerInjector interface {
    SetLogger(Logger)
}

func (s *Service) SetLogger(logger Logger) {
    s.logger = logger
}

// 4. Service Locator (NOT recommended)
type ServiceLocator map[string]interface{}

func (sl ServiceLocator) Get(name string) interface{} {
    return sl[name]
}

// Constructor injection is PREFERRED in Go because:
// ✓ Clear dependencies (visible in constructor)
// ✓ Compile-time safety
// ✓ Easy to understand
// ✓ No reflection/magic
`)

	fmt.Println("✓ DI concepts demonstrated")
}

// ConstructorInjectionDemo shows constructor pattern
func ConstructorInjectionDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    CONSTRUCTOR INJECTION DEMO                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. BASIC CONSTRUCTOR INJECTION:")
	fmt.Println(`
// Define interfaces (depend on abstractions, not implementations)
type Logger interface {
    Info(string)
    Error(string)
}

type Database interface {
    GetUser(id int) (*User, error)
    SaveUser(*User) error
}

// Service depends on interfaces
type UserService struct {
    logger Logger
    db     Database
}

// Constructor: dependency injection point
func NewUserService(logger Logger, db Database) *UserService {
    return &UserService{
        logger: logger,
        db:     db,
    }
}

// Use dependencies
func (s *UserService) Register(name, email string) (*User, error) {
    s.logger.Info(fmt.Sprintf("Registering user: %s", email))
    
    user := &User{Name: name, Email: email}
    err := s.db.SaveUser(user)
    if err != nil {
        s.logger.Error(fmt.Sprintf("Failed to save user: %v", err))
        return nil, err
    }
    
    s.logger.Info(fmt.Sprintf("User registered: %s", email))
    return user, nil
}
`)

	fmt.Println("\n2. REAL AND MOCK IMPLEMENTATIONS:")
	fmt.Println(`
// Real implementation
type RealLogger struct{}

func (l *RealLogger) Info(msg string) {
    fmt.Printf("[INFO] %s\n", msg)
}

func (l *RealLogger) Error(msg string) {
    fmt.Printf("[ERROR] %s\n", msg)
}

// Mock implementation for testing
type MockLogger struct {
    messages []string
}

func (l *MockLogger) Info(msg string) {
    l.messages = append(l.messages, msg)
}

func (l *MockLogger) Error(msg string) {
    l.messages = append(l.messages, msg)
}

// Real database
type RealDatabase struct{}

func (db *RealDatabase) GetUser(id int) (*User, error) {
    // Actual database query
    return queryDatabase(id)
}

func (db *RealDatabase) SaveUser(u *User) error {
    // Actual database write
    return saveToDatabase(u)
}

// Mock database for testing
type MockDatabase struct {
    users map[int]*User
}

func (db *MockDatabase) GetUser(id int) (*User, error) {
    user, exists := db.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (db *MockDatabase) SaveUser(u *User) error {
    db.users[u.ID] = u
    return nil
}
`)

	fmt.Println("\n3. TESTING WITH DI:")
	fmt.Println(`
func TestUserServiceRegister(t *testing.T) {
    // Use mock implementations
    mockLogger := &MockLogger{}
    mockDB := &MockDatabase{users: make(map[int]*User)}
    
    // Inject mocks
    service := NewUserService(mockLogger, mockDB)
    
    // Test
    user, err := service.Register("John", "john@example.com")
    
    // Verify
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if user.Name != "John" {
        t.Errorf("Got %s, want John", user.Name)
    }
    
    // Verify logger was called
    if len(mockLogger.messages) == 0 {
        t.Error("Logger should have been called")
    }
}

// In production: use real implementations
func main() {
    logger := &RealLogger{}
    db := &RealDatabase{}
    
    service := NewUserService(logger, db)
    // ... use service
}
`)

	fmt.Println("✓ Constructor injection demonstrated")
}

// InterfaceBasedDIDemo shows interface patterns
func InterfaceBasedDIDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      INTERFACE-BASED DI DEMO                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. DEFINE INTERFACES, NOT IMPLEMENTATIONS:")
	fmt.Println(`
// ✓ Good - Depend on interface
type Service struct {
    storage Storage  // Interface
}

interface Storage {
    Save(key string, value interface{}) error
    Get(key string) (interface{}, error)
}

// ✗ Bad - Depend on implementation
type Service struct {
    storage PostgresDatabase  // Concrete type
}

// This way, you can:
// 1. Use PostgreSQL in production
// 2. Use MySQL in other environment
// 3. Use in-memory store for testing
// All WITHOUT changing ServiceService code!
`)

	fmt.Println("\n2. COMPOSITION WITH INTERFACES:")
	fmt.Println(`
// Small, focused interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

// Compose into larger interface
type ReadWriter interface {
    Reader
    Writer
}

// Service depends on composed interface
type FileService struct {
    file ReadWriter
}

func (fs *FileService) Copy() error {
    data := make([]byte, 1024)
    n, _ := fs.file.Read(data)
    _, err := fs.file.Write(data[:n])
    return err
}

// Works with ANY ReadWriter implementation:
// - os.File
// - bufio.ReadWriter
// - bytes.Buffer
// - Network connection
// - Test mock
`)

	fmt.Println("\n3. INTERFACE SEGREGATION PRINCIPLE:")
	fmt.Println(`
// ✗ Large interface (bad)
type Everything interface {
    Read() error
    Write() error
    Delete() error
    List() error
    Execute() error
}

// ✓ Small, focused interfaces (good)
type Reader interface {
    Read() error
}

type Writer interface {
    Write() error
}

type Deleter interface {
    Delete() error
}

// Services only depend on what they need
type ReadService struct {
    r Reader  // Only needs Read
}

type WriteService struct {
    w Writer  // Only needs Write
}

// Benefits:
// - Easier to implement (only required methods)
// - Easier to test (mock less methods)
// - Easier to understand (clear responsibility)
// - More flexible (can use different implementations)
`)

	fmt.Println("✓ Interface-based DI demonstrated")
}

// ServiceLocatorDemo shows (avoided) service locator pattern
func ServiceLocatorDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      SERVICE LOCATOR PATTERN DEMO                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. SERVICE LOCATOR (NOT RECOMMENDED):")
	fmt.Println(`
// Service Locator: Central registry of services
type ServiceLocator struct {
    services map[string]interface{}
}

func (sl *ServiceLocator) Register(name string, service interface{}) {
    sl.services[name] = service
}

func (sl *ServiceLocator) Get(name string) interface{} {
    return sl.services[name]
}

// Usage:
locator := &ServiceLocator{services: make(map[string]interface{})}
locator.Register("logger", &RealLogger{})
locator.Register("db", &RealDatabase{})

type Service struct {
    locator *ServiceLocator
}

func (s *Service) GetLogger() Logger {
    return s.locator.Get("logger").(Logger)
}

// Problems with Service Locator:
// ✗ Hidden dependencies (not visible in constructor)
// ✗ Runtime errors (missing service discovered at runtime)
// ✗ Hard to test (need to set up entire locator)
// ✗ Hard to understand (where does dependency come from?)
// ✗ Magic strings (typos cause runtime errors)

// Why Constructor Injection is better:
// ✓ Dependencies are explicit and visible
// ✓ Compile-time safety
// ✓ Easy to test
// ✓ Type-safe (no casting needed)
// ✓ No magic or reflection
`)

	fmt.Println("✓ Service Locator pattern demonstrated")
}

// DIContainerDemo shows optional DI containers
func DIContainerDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       DI CONTAINER DEMO                                ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. SIMPLE DI CONTAINER (OPTIONAL):")
	fmt.Println(`
// For larger applications, optional DI container
type Container struct {
    logger Logger
    db     Database
    config Config
}

func NewContainer() *Container {
    return &Container{
        logger: &RealLogger{},
        db:     &RealDatabase{},
        config: &Config{},
    }
}

func (c *Container) NewUserService() *UserService {
    return NewUserService(c.logger, c.db)
}

// Usage:
container := NewContainer()
userService := container.NewUserService()

// Better than service locator because:
// ✓ Explicit dependencies (via factory methods)
// ✓ Type-safe (returns *UserService, not interface{})
// ✓ Compile-time safety
// ✓ Still has central place for construction

// Alternative: Wire generation (see github.com/google/wire)
// Generates dependency injection code at compile time
`)

	fmt.Println("\n2. TESTING WITH CONTAINER:")
	fmt.Println(`
type TestContainer struct {
    logger Logger
    db     Database
    config Config
}

func NewTestContainer() *TestContainer {
    return &TestContainer{
        logger: &MockLogger{},
        db:     &MockDatabase{users: make(map[int]*User)},
        config: &TestConfig{},
    }
}

func (c *TestContainer) NewUserService() *UserService {
    return NewUserService(c.logger, c.db)
}

func TestWithContainer(t *testing.T) {
    container := NewTestContainer()
    service := container.NewUserService()
    
    // Test with mocks
    user, err := service.Register("John", "john@example.com")
    if err != nil {
        t.Errorf("Error: %v", err)
    }
}
`)

	fmt.Println("✓ DI container demonstrated")
}

// RealWorldDIDemo shows practical application
func RealWorldDIDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    REAL WORLD DI EXAMPLE DEMO                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("COMPLETE DI APPLICATION:")
	fmt.Println(`
// File: services.go

type Logger interface {
    Info(string)
    Error(string)
}

type UserRepository interface {
    GetUser(id int) (*User, error)
    SaveUser(*User) error
    DeleteUser(id int) error
}

type EmailService interface {
    SendWelcome(email string) error
    SendPasswordReset(email, token string) error
}

// Service with injected dependencies
type UserService struct {
    logger Logger
    repo   UserRepository
    email  EmailService
}

func NewUserService(logger Logger, repo UserRepository, email EmailService) *UserService {
    return &UserService{
        logger: logger,
        repo:   repo,
        email:  email,
    }
}

func (s *UserService) Register(name, email, password string) (*User, error) {
    s.logger.Info("Registering user: " + email)
    
    // Validate
    if len(password) < 8 {
        return nil, fmt.Errorf("password too short")
    }
    
    // Create user
    user := &User{
        Name:     name,
        Email:    email,
        Password: hashPassword(password),
    }
    
    // Save to database (uses injected repository)
    err := s.repo.SaveUser(user)
    if err != nil {
        s.logger.Error("Failed to save user: " + err.Error())
        return nil, err
    }
    
    // Send email (uses injected email service)
    err = s.email.SendWelcome(email)
    if err != nil {
        s.logger.Error("Failed to send email: " + err.Error())
    }
    
    s.logger.Info("User registered: " + email)
    return user, nil
}

// Testing
func TestRegister(t *testing.T) {
    mockLogger := &MockLogger{}
    mockRepo := &MockUserRepository{}
    mockEmail := &MockEmailService{}
    
    service := NewUserService(mockLogger, mockRepo, mockEmail)
    user, err := service.Register("John", "john@example.com", "password123")
    
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    if user.Email != "john@example.com" {
        t.Errorf("Email mismatch")
    }
}
`)

	fmt.Println("✓ Real world DI example demonstrated")
}

// RunDependencyInjectionExamples executes all DI demos
func RunDependencyInjectionExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   DEPENDENCY INJECTION - COMPREHENSIVE GUIDE            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	DIConceptsDemo()
	ConstructorInjectionDemo()
	InterfaceBasedDIDemo()
	ServiceLocatorDemo()
	DIContainerDemo()
	RealWorldDIDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL DEPENDENCY INJECTION EXAMPLES COMPLETE           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ Inject dependencies via constructor (Go style)")
	fmt.Println("✓ Depend on interfaces, not implementations")
	fmt.Println("✓ Makes code testable and flexible")
	fmt.Println("✓ Avoid Service Locator (hidden dependencies)")
	fmt.Println("✓ Use small, focused interfaces")
	fmt.Println("✓ Interface Segregation Principle")
	fmt.Println("✓ Easy to swap implementations")
	fmt.Println("✓ Explicit dependencies (visible in code)")
	fmt.Println("✓ Type-safe (no casting needed)")
	fmt.Println("✓ Optional: Use DI container for complex apps\n")
}
