/*
FILE OPERATIONS AND FILE SYSTEM IN GO
============================================================================

Working with files is fundamental. Go provides excellent APIs for:
- Reading and writing files
- Directory operations
- File permissions
- Walking directory trees
- JSON/CSV operations

TOPICS COVERED:
- Basic file reading and writing
- Directory operations
- File permissions
- JSON marshaling/unmarshaling
- CSV operations
- Walking directory trees
*/

package concepts

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ============================================================================
// 1. BASIC FILE READING AND WRITING
// ============================================================================

func FileReadingWritingDemo() {
	fmt.Println("\n========== FILE READING & WRITING ==========")

	// 1. Reading entire file
	fmt.Println("--- Reading Entire File ---")

	// Write a test file first
	content := "Hello, World!\nLine 2\nLine 3"
	filename := "test_read.txt"
	os.WriteFile(filename, []byte(content), 0644)
	defer os.Remove(filename)

	// Read entire file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("File content:\n%s\n", string(data))

	// 2. Writing to file
	fmt.Println("\n--- Writing to File ---")

	writeFile := "test_write.txt"
	content = "This is new content\nwritten to file"

	err = os.WriteFile(writeFile, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Printf("Successfully wrote to %s\n", writeFile)
	defer os.Remove(writeFile)

	// 3. Appending to file
	fmt.Println("\n--- Appending to File ---")

	appendFile := "test_append.txt"
	os.WriteFile(appendFile, []byte("Line 1\n"), 0644)

	f, err := os.OpenFile(appendFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("Line 2\n")
	if err != nil {
		fmt.Println("Error appending:", err)
		return
	}
	fmt.Println("Successfully appended to file")
	defer os.Remove(appendFile)

	// 4. Reading line by line
	fmt.Println("\n--- Reading Line by Line ---")

	multilineFile := "test_multiline.txt"
	os.WriteFile(multilineFile, []byte("Line 1\nLine 2\nLine 3"), 0644)

	file, err := os.Open(multilineFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	defer os.Remove(multilineFile)

	fmt.Println("\n✓ File reading/writing demonstrated")
}

// ============================================================================
// 2. DIRECTORY OPERATIONS
// ============================================================================

func DirectoryOperationsDemo() {
	fmt.Println("\n========== DIRECTORY OPERATIONS ==========")

	// 1. Create directory
	fmt.Println("--- Creating Directories ---")

	dirName := "test_dir"
	err := os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Printf("Created directory: %s\n", dirName)
	defer os.RemoveAll(dirName)

	// Create nested directories
	nestedDir := "test_dir/subdir/nested"
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Created nested directory: %s\n", nestedDir)

	// 2. List directory contents
	fmt.Println("\n--- Listing Directory Contents ---")

	testDir := "test_dir"
	entries, err := os.ReadDir(testDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Printf("Contents of %s:\n", testDir)
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  [DIR] %s\n", entry.Name())
		} else {
			fmt.Printf("  [FILE] %s\n", entry.Name())
		}
	}

	// 3. Check if file/directory exists
	fmt.Println("\n--- Checking File Existence ---")

	exists := true
	_, err = os.Stat(testDir)
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		}
	}
	fmt.Printf("%s exists: %v\n", testDir, exists)

	// 4. Get file info
	fmt.Println("\n--- Getting File Information ---")

	// Create a test file
	testFile := filepath.Join(testDir, "info.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	info, err := os.Stat(testFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Modified: %v\n", info.ModTime())
	fmt.Printf("Is Directory: %v\n", info.IsDir())

	fmt.Println("\n✓ Directory operations demonstrated")
}

// ============================================================================
// 3. WALKING DIRECTORY TREES
// ============================================================================

func WalkDirectoryDemo() {
	fmt.Println("\n========== WALKING DIRECTORY TREES ==========")

	// Create test directory structure
	testDir := "test_tree"
	os.MkdirAll(filepath.Join(testDir, "subdir1", "nested"), 0755)
	os.MkdirAll(filepath.Join(testDir, "subdir2"), 0755)
	os.WriteFile(filepath.Join(testDir, "file1.txt"), []byte("1"), 0644)
	os.WriteFile(filepath.Join(testDir, "subdir1", "file2.txt"), []byte("2"), 0644)
	os.WriteFile(filepath.Join(testDir, "subdir1", "nested", "file3.txt"), []byte("3"), 0644)
	defer os.RemoveAll(testDir)

	fmt.Printf("Walking directory tree: %s\n\n", testDir)

	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate indentation based on depth
		relPath, _ := filepath.Rel(testDir, path)
		depth := strings.Count(relPath, string(os.PathSeparator))
		indent := strings.Repeat("  ", depth)

		if info.IsDir() {
			fmt.Printf("%s[DIR] %s\n", indent, info.Name())
		} else {
			fmt.Printf("%s[FILE] %s (%d bytes)\n", indent, info.Name(), info.Size())
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
	}

	fmt.Println("\n✓ Directory tree walking demonstrated")
}

// ============================================================================
// 4. JSON OPERATIONS - Very Common in Go
// ============================================================================

type Personfs struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	City  string `json:"city"`
}

type Companyfs struct {
	Name      string     `json:"name"`
	Founded   int        `json:"founded"`
	Employees []Personfs `json:"employees"`
}

func JSONOperationsDemo() {
	fmt.Println("\n========== JSON OPERATIONS ==========")

	// 1. Marshal (Go → JSON)
	fmt.Println("--- Marshaling (Go to JSON) ---")

	person := Person{
		Name:  "Alice",
		Age:   28,
		Email: "alice@example.com",
		City:  "Nairobi",
	}

	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Printf("JSON: %s\n", string(jsonBytes))

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Pretty JSON:\n%s\n", string(prettyJSON))

	// 2. Unmarshal (JSON → Go)
	fmt.Println("\n--- Unmarshaling (JSON to Go) ---")

	jsonString := `{"name":"Bob","age":35,"email":"bob@example.com","city":"Mombasa"}`

	var person2 Person
	err = json.Unmarshal([]byte(jsonString), &person2)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Printf("Unmarshaled: %+v\n", person2)

	// 3. Working with files
	fmt.Println("\n--- JSON File Operations ---")

	company := Companyfs{
		Name:    "TechCorp",
		Founded: 2010,
		Employees: []Personfs{
			{Name: "Alice", Age: 28, Email: "alice@techcorp.com", City: "Nairobi"},
			{Name: "Bob", Age: 35, Email: "bob@techcorp.com", City: "Mombasa"},
		},
	}

	// Write to file
	jsonFile := "company.json"
	data, err := json.MarshalIndent(company, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	os.WriteFile(jsonFile, data, 0644)
	fmt.Printf("Wrote company to %s\n", jsonFile)
	defer os.Remove(jsonFile)

	// Read from file
	fileData, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var loadedCompany Companyfs
	err = json.Unmarshal(fileData, &loadedCompany)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Loaded company: %s with %d employees\n",
		loadedCompany.Name, len(loadedCompany.Employees))

	fmt.Println("\n✓ JSON operations demonstrated")
}

// ============================================================================
// 5. CSV OPERATIONS
// ============================================================================

func CSVOperationsDemo() {
	fmt.Println("\n========== CSV OPERATIONS ==========")

	// 1. Writing CSV
	fmt.Println("--- Writing CSV ---")

	csvFile := "data.csv"
	file, err := os.Create(csvFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	defer os.Remove(csvFile)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Name", "Age", "City"})

	// Write data
	writer.Write([]string{"Alice", "28", "Nairobi"})
	writer.Write([]string{"Bob", "35", "Mombasa"})
	writer.Write([]string{"Charlie", "42", "Kampala"})

	fmt.Printf("Wrote CSV to %s\n", csvFile)

	// 2. Reading CSV
	fmt.Println("\n--- Reading CSV ---")

	file2, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	reader := csv.NewReader(file2)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	fmt.Println("CSV Contents:")
	for i, record := range records {
		if i == 0 {
			fmt.Printf("  Header: %v\n", record)
		} else {
			fmt.Printf("  Row %d: %v\n", i, record)
		}
	}

	fmt.Println("\n✓ CSV operations demonstrated")
}

// ============================================================================
// PRACTICAL EXAMPLE: Configuration Management
// ============================================================================

type AppConfig struct {
	AppName        string   `json:"app_name"`
	Port           int      `json:"port"`
	DatabaseURL    string   `json:"database_url"`
	Features       []string `json:"features"`
	MaxConnections int      `json:"max_connections"`
}

func LoadConfig(filename string) (*AppConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config AppConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}

func SaveConfig(filename string, config *AppConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func ConfigManagementDemo() {
	fmt.Println("\n========== PRACTICAL EXAMPLE: Config Management ==========")

	// Create a config
	config := &AppConfig{
		AppName:        "MyApp",
		Port:           8080,
		DatabaseURL:    "postgresql://localhost/mydb",
		Features:       []string{"auth", "api", "websocket"},
		MaxConnections: 100,
	}

	// Save config
	configFile := "app_config.json"
	err := SaveConfig(configFile, config)
	if err != nil {
		fmt.Println("Error saving config:", err)
		return
	}
	fmt.Printf("Saved config to %s\n", configFile)
	defer os.Remove(configFile)

	// Load config
	loadedConfig, err := LoadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("Loaded config: App=%s, Port=%d, Features=%v\n",
		loadedConfig.AppName, loadedConfig.Port, loadedConfig.Features)

	fmt.Println("\n✓ Configuration management demonstrated")
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunFileSystemExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    FILE OPERATIONS & FILE SYSTEM - COMPLETE GUIDE      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	FileReadingWritingDemo()
	DirectoryOperationsDemo()
	WalkDirectoryDemo()
	JSONOperationsDemo()
	CSVOperationsDemo()
	ConfigManagementDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      FILE SYSTEM MASTERY - MANAGE DATA LIKE A PRO!      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}
