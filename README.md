# 🦫 GophersMap: Learn & Practice Go

---

## 📚 Table of Contents
- [Getting Started with Golang](#getting-started-with-golang)
- [Project Overview](#project-overview)
- [Advantages of Golang](#advantages-of-golang)
- [Exercises Preview](#exercises-preview)
- [About allExercises.md](#about-allexercisesmd)
- [Contributing](#contributing)
- [License](#license)

---

## 🚀 Getting Started with Golang

### 1. Install Go
- Visit [golang.org/dl](https://golang.org/dl/) and download the installer for your OS.
- Follow the installation instructions for your platform.
- Verify installation:
  ```sh
  go version
  ```

### 2. Set Up Your First Project
- Create a new folder for your project:
  ```sh
  mkdir GophersMap
  cd GophersMap
  ```
- Initialize a Go module:
  ```sh
  go mod init github.com/yourusername/GophersMap
  ```

### 3. Run Your First Program
- Create a file called `main.go`:
  ```go
  package main
  import "fmt"
  func main() {
      fmt.Println("Hello, Gophers!")
  }
  ```
- Run it:
  ```sh
  go run main.go
  ```

---

## 🗺️ Project Overview
GophersMap is a beginner-friendly Go learning resource. It is split into two main sections:

- **concepts/**: Contains annotated Go files explaining key language features and concepts.
- **exercise/**: Contains practical coding exercises to reinforce your learning.

---

## ✨ Advantages of Golang
- **Simplicity**: Clean, readable syntax.
- **Performance**: Compiled language with fast execution.
- **Concurrency**: Built-in support for goroutines and channels.
- **Cross-Platform**: Easy to build for Windows, Linux, and macOS.
- **Strong Standard Library**: Rich set of packages for web, file I/O, and more.
- **Static Typing**: Catch errors at compile time.
- **Great Tooling**: Formatters, linters, and package management built-in.

---

## 🏋️ Exercises Preview
Here are some of the exercises you'll find in the `exercise/` folder:

| #  | Title                  | Description                                 |
|----|------------------------|---------------------------------------------|
| 1  | Hello, World!          | Print a greeting to the console             |
| 5  | Reverse a String       | Write a function to reverse a string        |
| 10 | Struct for Person      | Define and use a custom struct              |
| 15 | Simple HTTP Server     | Build a basic web server in Go              |
| 28 | RESTful API (CRUD)     | Create a simple API for managing tasks      |
| 35 | Concurrent Web Scraper | Scrape multiple websites concurrently       |

---

## 📄 About allExercises.md
The file [`allExercises.md`](exercise/allExercises.md) contains a curated list of 35 exercises, ranging from basic syntax to real-world applications. Use it as your roadmap to master Go, tackling each challenge in order of increasing complexity.

---

## 🤝 Contributing
Contributions are welcome! Feel free to submit pull requests for new exercises, improved explanations, or bug fixes.

---

## 📜 License
This project is licensed under the MIT License.

---

> Happy Coding, and welcome to the Go community! 🦫
