# Cracking the Coding Interview Solutions (Go)

This repository contains my Go implementations and detailed explanations for problems from  
*Cracking the Coding Interview* by Gayle Laakmann McDowell.

The goal of this repository is to strengthen my understanding of algorithms, data structures, and system design concepts by solving each problem in Go and documenting the reasoning behind every solution.

---

## Repository Overview

Each chapter of the book is implemented as a separate directory.  
Within each directory, you will find:
- A Go source file containing the solution.
- A test file written using Go’s `testing` package.

Each solution includes:
- A concise summary of the problem.
- Explanation of the algorithm and reasoning.
- Time and space complexity analysis.
- Edge cases considered.
- Unit tests for correctness and reliability.

---

## Setup and Usage

### Prerequisites
- Go 1.22 or later installed on your system.
- Git for version control.

### Clone the Repository
```bash
git clone https://github.com/BerkAkipek/cracking-the-coding-interview-solutions-go.git
cd cracking-the-coding-interview-solutions-go
```

### Run a Specific Solution
```bash
go run ./01_Arrays_and_Strings/is_unique.go
```

### Run All Tests
```bash
go test ./...
```

### Run Tests Verbosely
```bash
go test -v ./...
```

## Topics Covered

- Arrays and Strings
- Linked Lists
- Stacks and Queues
- Trees and Graphs
- Bit Manipulation
- Recursion and Dynamic Programming
- Object-Oriented Design
- System Design
- Sorting and Searching
- Mathematical and Logical Puzzles

## Goals

- Build a complete Go-based reference for Cracking the Coding Interview
- Practice writing efficient and idiomatic Go code
- Develop clear and maintainable problem-solving patterns
- Prepare effectively for technical interviews

## Contributing

This repository is primarily for personal learning and documentation.  
However, feedback, suggestions, and discussions through issues or pull requests are welcome.

If you’d like to contribute:
1. Fork the repository.
2. Create a new branch for your feature or fix.
    ```bash
    git checkout -b feature/new-solution
    ```
3. Make your changes and ensure all tests pass.
```bash
go test ./... -v
```
4. Commit your changes with a clear message.
```bash
git commit -m "Add solution for <problem-name>"
```
5. Push your branch and open a pull request.

## License

This project is licensed under the **MIT License**.  
You are free to use, modify, and distribute this code for personal or commercial purposes, provided that proper attribution is given.

---

## Author

**Berk Akipek**  
[GitHub](https://github.com/BerkAkipek) | [LinkedIn](https://www.linkedin.com/in/berk-akipek)
