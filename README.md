# EchoDiff - A Smarter Diff Tool

EchoDiff is a command-line tool written in Go that compares two files and provides a human-readable summary of their differences. It can be used to quickly understand code or text changes without manually inspecting raw diffs.

## Features
✅ Compare two files and generate a line-by-line diff  
✅ Provide a summarized explanation of changes in natural language  
✅ Simple and lightweight with minimal dependencies  

---

## Installation

### **Prerequisites**
- Go 1.18+ installed on your system ([Download Go](https://go.dev/dl/))

### **Clone the Repository**
```sh
git clone https://github.com/jmendick/echodiff.git
cd echodiff
```

### **Initialize the Go Module**
```sh
go mod tidy
```

### **Build the Binary**
```sh
go build -o echodiff
```
This will create an executable file named `echodiff`.

---

## Usage

### **Compare Two Files**
To see a simple diff between two files:
```sh
go run main.go diff file1.txt file2.txt
```
or, if you built the binary:
```sh
./echodiff diff file1.txt file2.txt
```
This will output a raw textual diff.

### **Get a Human-Readable Explanation**
```sh
go run main.go explain file1.txt file2.txt
```
or
```sh
./echodiff explain file1.txt file2.txt
```
This provides a summary of added and removed lines in a more readable format.

### **Help Command**
To see all available commands:
```sh
go run main.go --help
```
or  
```sh
./echodiff --help
```

---

## Example Output

### **Raw Diff (`diff` command)**
```
- Removed line: "Old function definition"
+ Added line: "New function definition"
```

### **Human-Readable Explanation (`explain` command)**
```
Changes detected:
- Removed: "Old function definition"
+ Added: "New function definition"
```

---

## Roadmap

🔹 **Git Integration** – Compare staged changes using `git diff`  
🔹 **Interactive Mode** – Ask questions about file changes  
🔹 **Color-coded Terminal Output** – Improved readability  
🔹 **Markdown & JSON Output** – Export results in multiple formats  
🔹 **AI-Powered Explanations** – Summarize changes in plain English using AI  

---

## Contributing
Contributions are welcome! Feel free to submit issues or pull requests.

1. Fork the repository  
2. Create a new feature branch  
3. Commit your changes  
4. Push to your fork and create a PR  

---

## License
MIT License. See [LICENSE](LICENSE) for details.

---

## Author
👤 **Joe Mendick**  
📧 Contact: [joe.mendick@gmail.com]  
GitHub: [github.com/jmendick](https://github.com/jmendick)

---
