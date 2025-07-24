# 🖼️ ASCII Art Converter

Convert your images to ASCII art with an optional TUI interface built in Go!

## ✨ Features

- 🖥️ TUI (Terminal UI) to interactively convert images to ASCII
- 🔁 Converts image to ASCII representation using custom ramps
- 📏 Resize images with specified width
- 🧱 Save the output as a JPEG image
- 🧑‍💻 Command-line support for quick conversions

## 🛠 Installation

### Option 1: Install via `go install`

```bash
go install github.com/kais-blkc/ascii_art/cmd/ascii_art
```
This will download, build, and install the ascii_art binary into your $GOBIN directory (usually ~/go/bin). Make sure it's in your PATH.

### Option 2: Manual clone & build
```bash
git clone https://github.com/kais-blkc/ascii_art.git
cd ascii_art/cmd/ascii_art
go build -o ascii_art
```

## 🚀 Usage
### 1. Run with TUI
Simply launch the app:
```bash
ascii_art
```
This will start an interactive terminal UI for converting your image.

### 2. Run in CLI mode (without UI)
```bash
ascii_art <input-file> <width>
```

Example:
```bash
ascii_art ./example.jpg 100
```

This will:
- Load example.jpg
- Resize it to width 100
- Convert it to ASCII
- Save the result as output.jpg in the current directory
- If width is invalid, it defaults to 80.




