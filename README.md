# http7 - High Performance Static Web Server

A lightweight, high-performance static file server written in Go, designed as a more professional and secure alternative to `python -m http.server`.

## 🚀 Command Line Arguments

You can configure the server behavior using the following arguments:

| Argument | Default | Description |
| :--- | :--- | :--- |
| **`-p`** | `8080` | Specifies the port number to listen on. |
| **`-s`** | `nginx` | Custom `Server` header (Identity Masking) to hide or spoof server information. |
| **`-d`** | `false` | **Directory Listing**. Disabled by default for security. When enabled, it allows users to browse file lists if no `index.html` is present. |

**Example:**
```bash
# Start on port 9090, spoof as cloudflare, and enable directory listing
./http7 -p 9090 -s cloudflare -d
```

## 📦 Installation & Running

### Method 1: Using go install (Recommended)
This installs `http7` to your `$GOPATH/bin` directory, allowing you to run it from any folder.

```bash
go install github.com/tvxy/http7@latest
```
*(Note: Ensure your GOPATH is correctly configured in your environment variables)*

### Method 2: Direct Run via Go
If you have the Go environment installed locally, you can run the source code directly:

```bash
go run main.go -p 8080
```

### Method 3: Download or Build Binary
You can either download pre-compiled binaries for quick use or build them yourself.

1. **Download from Releases**:
   Go to the [Releases](https://github.com/tvxy/http7/releases) page, download the archive for your OS, and extract the executable.

2. **Self-Compile** (Requires Go):
   ```bash
   # Compiles into a standalone binary named 'http7'
   go build -o http7 main.go
   ```

3. **Run**:
   ```bash
   ./http7
   ```

## 💡 Why http7?

**http7** is designed to provide a more robust and secure static serving experience than the common `python -m http.server`.

**Key Advantages:**

1.  **Detailed Audit Logs**:
    *   **Python**: Only provides basic request paths and status codes.
    *   **http7**: Features production-grade logging including **Precise Timestamps**, **Client IP**, **HTTP Method**, **Status Codes**, **Response Latency**, and **User-Agent** for better traffic analysis.

2.  **Enhanced Security**:
    *   **Python**: Exposes all files by default, which can lead to data leaks.
    *   **http7**: **Disables directory listing by default** and automatically serves `index.html`. The custom Server header (`-s`) also prevents fingerprinting by scanners.

3.  **High Performance & Zero Dependencies**:
    *   Compiled as a single standalone binary with **zero external dependencies** (no Python required). Leveraging Go's concurrency model, it handles concurrent requests significantly better than Python's single-threaded server.
