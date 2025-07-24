# ASCII Art Generator

This project is a Terminal User Interface (TUI) application for converting images (JPEG, PNG) into ASCII art.

## Description

The application allows users to select an image from their filesystem, configure conversion parameters, and save the resulting ASCII art as a new JPEG image.

## Features

-   Interactive terminal-based interface.
-   Preview the selected image as ASCII art.
-   Customize parameters for ASCII art generation.
-   Save the result as a JPEG file.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd ascii_art_go
    ```

2.  **Install dependencies:**
    Ensure you have Go installed (version 1.24 or newer). The dependencies will be downloaded automatically when you build or run the project.

## Usage

There are several ways to run the application.

### 1. Using `go run`

This is the simplest way to run the application.

```bash
go run main.go
```

### 2. Build and Run the Binary

You can build the application and then run the compiled binary.

```bash
# Build
go build -o ascii-art-generator .

# Run
./ascii-art-generator
```

### 3. For Development (with `air`)

The project is configured to use `air` for live reloading during development.

1.  **Install `air`:**
    ```bash
    go install github.com/cosmtrek/air@latest
    ```

2.  **Run `air`:**
    ```bash
    air
    ```
    `air` will automatically rebuild and restart the application whenever the source code changes.

## Key Dependencies

-   [github.com/rivo/tview](https://github.com/rivo/tview) - For building the terminal user interface.
-   [github.com/nfnt/resize](https://github.com/nfnt/resize) - For resizing images.
-   [golang.org/x/image](https://pkg.go.dev/golang.org/x/image) - For image manipulation.

## Project Structure

```
.
├── internal/         # Internal application logic
│   ├── convert/      # Image conversion functions
│   ├── image_utils/  # Image manipulation utilities
│   ├── shared/       # Shared code (constants, events)
│   └── ui/           # User interface components
├─��� main.go           # Application entry point
├── go.mod            # Go dependencies file
└── .air.toml         # Configuration for live-reload
```

## License

This project is unlicensed. Feel free to add a license if needed.
