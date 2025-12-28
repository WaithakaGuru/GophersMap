package concepts

import (
	"fmt"
)

// ============================================================================
// 1. DEPENDENCY INJECTION CONCEPTS
// ============================================================================

/*
WHAT IS DEPENDENCY INJECTION?
- Dependency: Something a component needs to work
- Injection: Providing that dependency from outside
- Pattern: Pass dependencies in, don't create them inside

WITHOUT DI (BAD - TIGHTLY COUPLED):
- Service creates its own dependencies
- Hard to test (can't use mocks)
- Can't swap implementations
- Hidden dependencies

WITH DI (GOOD - LOOSELY COUPLED):
- Dependencies are passed in
- Easy to test (inject mocks)
- Can swap implementations
- Clear dependencies

BENEFITS:
✓ Testability: Use mocks instead of real dependencies
✓ Flexibility: Swap implementations without changing code
✓ Maintainability: Changes in one place
✓ Decoupling: Components don't depend on each other
✓ Clarity: Dependencies are explicit
*/

// DIConceptsDemo explains DI fundamentals
func DIConceptsDemo() {
	fmt.Println("\n========== DEPENDENCY INJECTION CONCEPTS ==========")

	fmt.Println("\n1. PROBLEM WITHOUT DI (TIGHTLY COUPLED):")
	fmt.Println("   - Service creates its own database")
	fmt.Println("   - Hard to test (can't use mock database)")
	fmt.Println("   - Can't use different databases")
	fmt.Println("   - Database choice is hardcoded")

	fmt.Println("\n2. SOLUTION WITH DI (LOOSELY COUPLED):")
	fmt.Println("   - Service receives database via constructor")
	fmt.Println("   - Easy to test (pass mock database)")
	fmt.Println("   - Can use any database implementation")
	fmt.Println("   - Dependencies are explicit")

	fmt.Println("\n3. EXAMPLE:")
	fmt.Println("   Without DI: service := NewService()  // Creates database internally")
	fmt.Println("   With DI:    service := NewService(db, logger)  // Receives dependencies")

	fmt.Println("\n4. GO'S PREFERRED: CONSTRUCTOR INJECTION")
	fmt.Println("   - Dependencies in constructor parameters")
	fmt.Println("   - Clear what dependencies are needed")
	fmt.Println("   - Type-safe (no casting)")
	fmt.Println("   - No reflection or magic")

	fmt.Println("\n✓ DI concepts demonstrated")
}

// ============================================================================
// 2. CONSTRUCTOR INJECTION PATTERN
// ============================================================================

/*
CONSTRUCTOR INJECTION:
- Dependencies passed to constructor
- Most common in Go
- Type-safe and explicit
- Easy to understand

PATTERN:
1. Define interfaces for dependencies
2. Add fields to struct
3. Create constructor function
4. Accept dependencies as parameters
5. Return initialized struct

This is how Go does dependency injection!
*/

// Example interfaces
type Logger interface {
	Info(string)
	Error(string)
}

type Database interface {
	Query(string) error
	Save(interface{}) error
}

// Service with injected dependencies
type UserService struct {
	logger Logger
	db     Database
}

// Constructor: This is the dependency injection point
func NewUserService(logger Logger, db Database) *UserService {
	return &UserService{
		logger: logger,
		db:     db,
	}
}

// Method uses injected dependencies
func (s *UserService) CreateUser(name string) error {
	s.logger.Info("Creating user: " + name)
	err := s.db.Save(map[string]string{"name": name})
	if err != nil {
		s.logger.Error("Failed to create user")
		return err
	}
	s.logger.Info("User created successfully")
	return nil
}

// ConstructorInjectionDemo shows the pattern
func ConstructorInjectionDemo() {
	fmt.Println("\n========== CONSTRUCTOR INJECTION ==========")

	fmt.Println("\n1. DEFINE INTERFACES:")
	fmt.Println("   type Logger interface { Info(string); Error(string) }")
	fmt.Println("   type Database interface { Query(string) error; Save(interface{}) error }")

	fmt.Println("\n2. ADD FIELDS TO STRUCT:")
	fmt.Println("   type UserService struct {")
	fmt.Println("       logger Logger")
	fmt.Println("       db     Database")
	fmt.Println("   }")

	fmt.Println("\n3. CREATE CONSTRUCTOR:")
	fmt.Println("   func NewUserService(logger Logger, db Database) *UserService {")
	fmt.Println("       return &UserService{logger: logger, db: db}")
	fmt.Println("   }")

	fmt.Println("\n4. USE DEPENDENCIES:")
	fmt.Println("   func (s *UserService) CreateUser(name string) error {")
	fmt.Println("       s.logger.Info(\"Creating user\")")
	fmt.Println("       return s.db.Save(user)")
	fmt.Println("   }")

	fmt.Println("\n5. ADVANTAGES:")
	fmt.Println("   ✓ Dependencies are explicit (visible in constructor)")
	fmt.Println("   ✓ Type-safe (compiler checks types)")
	fmt.Println("   ✓ Easy to test (pass mock implementations)")
	fmt.Println("   ✓ No magic or reflection")
	fmt.Println("   ✓ Clear what dependencies are needed")

	fmt.Println("\n✓ Constructor injection demonstrated")
}

// ============================================================================
// 3. INTERFACE-BASED DESIGN
// ============================================================================

/*
KEY PRINCIPLE: DEPEND ON ABSTRACTIONS (INTERFACES), NOT IMPLEMENTATIONS

WHY INTERFACES?
- Code against contracts, not implementations
- Easy to swap implementations
- Easy to create mocks for testing
- Loosely coupled

INTERFACE SEGREGATION:
- Small, focused interfaces
- Services only depend on what they need
- More flexible and easier to test
*/

// ✓ Good - depend on interface
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

// FileProcessor depends on interfaces, not implementations
type FileProcessor struct {
	reader Reader
	writer Writer
}

func NewFileProcessor(r Reader, w Writer) *FileProcessor {
	return &FileProcessor{reader: r, writer: w}
}

func (fp *FileProcessor) Copy() error {
	data := make([]byte, 1024)
	n, err := fp.reader.Read(data)
	if err != nil {
		return err
	}
	_, err = fp.writer.Write(data[:n])
	return err
}

// InterfaceBasedDIDemo explains interface design
func InterfaceBasedDIDemo() {
	fmt.Println("\n========== INTERFACE-BASED DI ==========")

	fmt.Println("\n1. DEFINE FOCUSED INTERFACES:")
	fmt.Println("   type Reader interface { Read([]byte) (int, error) }")
	fmt.Println("   type Writer interface { Write([]byte) (int, error) }")

	fmt.Println("\n2. DEPEND ON INTERFACES:")
	fmt.Println("   type Processor struct {")
	fmt.Println("       reader Reader  // Interface, not implementation")
	fmt.Println("       writer Writer")
	fmt.Println("   }")

	fmt.Println("\n3. WORKS WITH ANY IMPLEMENTATION:")
	fmt.Println("   - os.File")
	fmt.Println("   - bufio.Reader/Writer")
	fmt.Println("   - bytes.Buffer")
	fmt.Println("   - Network connection")
	fmt.Println("   - Test mock")

	fmt.Println("\n4. INTERFACE SEGREGATION PRINCIPLE:")
	fmt.Println("   - Small interfaces (Reader, Writer)")
	fmt.Println("   - Not large interfaces (CanReadWriteExecute...)")
	fmt.Println("   - Services only depend on what they need")
	fmt.Println("   - Easier to implement")
	fmt.Println("   - Easier to test (mock less)")

	fmt.Println("\n✓ Interface-based DI demonstrated")
}

// ============================================================================
// 4. TESTING WITH DI
// ============================================================================

/*
DI MAKES TESTING EASY:
- Create mock implementations
- Pass mocks to constructors
- Test logic in isolation
- No need for complex test setup
*/

// MockLogger for testing
type MockLogger struct {
	Messages []string
}

func (m *MockLogger) Info(msg string) {
	m.Messages = append(m.Messages, "INFO: "+msg)
}

func (m *MockLogger) Error(msg string) {
	m.Messages = append(m.Messages, "ERROR: "+msg)
}

// MockDatabase for testing
type MockDatabase struct {
	Data map[string]interface{}
}

func (m *MockDatabase) Query(q string) error {
	// Mock implementation
	return nil
}

func (m *MockDatabase) Save(data interface{}) error {
	m.Data["test"] = data
	return nil
}

// Example test using DI
/*
func TestUserServiceWithDI(t *testing.T) {
	mockLogger := &MockLogger{}
	mockDB := &MockDatabase{Data: make(map[string]interface{})}

	// Inject mocks
	service := NewUserService(mockLogger, mockDB)

	// Test
	err := service.CreateUser("John")

	// Verify
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Check that logger was called
	if len(mockLogger.Messages) == 0 {
		t.Error("Logger should have been called")
	}
}
*/

// DITestingDemo shows testing with DI
func DITestingDemo() {
	fmt.Println("\n========== TESTING WITH DI ==========")

	fmt.Println("\n1. CREATE MOCK IMPLEMENTATIONS:")
	fmt.Println("   type MockLogger struct { Messages []string }")
	fmt.Println("   type MockDatabase struct { Data map[string]interface{} }")

	fmt.Println("\n2. INJECT MOCKS IN TEST:")
	fmt.Println("   mockLogger := &MockLogger{}")
	fmt.Println("   mockDB := &MockDatabase{Data: make(map[string]interface{})}")
	fmt.Println("   service := NewUserService(mockLogger, mockDB)")

	fmt.Println("\n3. TEST LOGIC:")
	fmt.Println("   err := service.CreateUser(\"John\")")
	fmt.Println("   if err != nil { t.Error(\"Failed\") }")

	fmt.Println("\n4. VERIFY BEHAVIOR:")
	fmt.Println("   if len(mockLogger.Messages) == 0 { t.Error(\"Logger not called\") }")
	fmt.Println("   if mockDB.Data[\"test\"] == nil { t.Error(\"Data not saved\") }")

	fmt.Println("\n5. BENEFITS:")
	fmt.Println("   ✓ No database connection needed")
	fmt.Println("   ✓ No file I/O")
	fmt.Println("   ✓ No network calls")
	fmt.Println("   ✓ Fast tests")
	fmt.Println("   ✓ Reliable tests")
	fmt.Println("   ✓ Easy to verify behavior")

	fmt.Println("\n✓ DI testing demonstrated")
}

// ============================================================================
// 5. AVOIDING SERVICE LOCATOR
// ============================================================================

/*
SERVICE LOCATOR ANTI-PATTERN:
- Central registry of services
- Services look up dependencies
- AVOID THIS IN GO!

PROBLEMS WITH SERVICE LOCATOR:
✗ Hidden dependencies (not visible in code)
✗ Runtime errors (missing service discovered at runtime)
✗ Hard to test (need entire locator)
✗ Magic strings (typos cause errors)
✗ Hard to understand where dependencies come from
✗ No compile-time safety

USE CONSTRUCTOR INJECTION INSTEAD!
*/

// ServiceLocatorDemo shows anti-pattern
func ServiceLocatorDemo() {
	fmt.Println("\n========== AVOID SERVICE LOCATOR ==========")

	fmt.Println("\n1. SERVICE LOCATOR (BAD):")
	fmt.Println("   type ServiceLocator struct {")
	fmt.Println("       services map[string]interface{}")
	fmt.Println("   }")
	fmt.Println("")
	fmt.Println("   func (sl *ServiceLocator) Get(name string) interface{} {")
	fmt.Println("       return sl.services[name]  // Magic string!")
	fmt.Println("   }")

	fmt.Println("\n2. PROBLEMS:")
	fmt.Println("   ✗ Dependencies hidden (not in constructor)")
	fmt.Println("   ✗ \"logger\" is a magic string (typo-prone)")
	fmt.Println("   ✗ Must set up entire locator for tests")
	fmt.Println("   ✗ Runtime errors (missing services)")
	fmt.Println("   ✗ Hard to understand")

	fmt.Println("\n3. BETTER: CONSTRUCTOR INJECTION")
	fmt.Println("   func NewService(logger Logger) *Service {")
	fmt.Println("       return &Service{logger: logger}")
	fmt.Println("   }")

	fmt.Println("\n4. ADVANTAGES:")
	fmt.Println("   ✓ Dependencies are explicit")
	fmt.Println("   ✓ Type-safe (no magic strings)")
	fmt.Println("   ✓ Compile-time safety")
	fmt.Println("   ✓ Easy to test")
	fmt.Println("   ✓ Easy to understand")

	fmt.Println("\n✓ Service Locator anti-pattern explained")
}

// ============================================================================
// 6. MAIN EXECUTION
// ============================================================================

// RunDependencyInjectionExamples executes all DI demonstrations
func RunDependencyInjectionExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   DEPENDENCY INJECTION - COMPREHENSIVE GUIDE            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	DIConceptsDemo()
	ConstructorInjectionDemo()
	InterfaceBasedDIDemo()
	DITestingDemo()
	ServiceLocatorDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL DEPENDENCY INJECTION EXAMPLES COMPLETE           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Println("\nKEY TAKEAWAYS:")
	fmt.Println("✓ Inject dependencies via constructor")
	fmt.Println("✓ Depend on interfaces, not implementations")
	fmt.Println("✓ Makes code testable with mocks")
	fmt.Println("✓ Makes code flexible (swap implementations)")
	fmt.Println("✓ Makes dependencies explicit and visible")
	fmt.Println("✓ Use small, focused interfaces")
	fmt.Println("✓ Interface Segregation Principle")
	fmt.Println("✓ AVOID Service Locator pattern")
	fmt.Println("✓ Type-safe, no magic or reflection")
	fmt.Println("✓ Loosely coupled, highly testable code\n")
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
