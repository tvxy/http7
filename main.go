package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Port       string
	ServerName string
	EnableDir  bool
	RootDir    string
}

func init() {
	log.SetFlags(0)
}

func printBanner() {
	banner := `
 ██╗  ██╗████████╗████████╗██████╗ ███████╗
 ██║  ██║╚══██╔══╝╚══██╔══╝██╔══██╗╚════██║
 ███████║   ██║      ██║   ██████╔╝    ██╔╝
 ██╔══██║   ██║      ██║   ██╔═══╝    ██╔╝
 ██║  ██║   ██║      ██║   ██║        ██║
 ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝        ╚═╝

      >> Developed by Hex <<
   https://github.com/tvxy/http7
`
	fmt.Println(banner)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func main() {
	port := flag.String("p", "8080", "Port to listen on")
	serverName := flag.String("s", "nginx", "Custom Server header")
	enableDir := flag.Bool("d", false, "Enable directory listing")
	flag.Parse()

	config := Config{
		Port:       *port,
		ServerName: *serverName,
		EnableDir:  *enableDir,
		RootDir:    ".",
	}

	printBanner()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{w, http.StatusOK}

		handleRequest(lrw, r, config)

		duration := time.Since(start)
		timestamp := time.Now().Format("2006/01/02 15:04:05")
		fmt.Printf("%s | %s | %s | %s | %d | %v | %s\n",
			timestamp,
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			lrw.statusCode,
			duration,
			r.UserAgent(),
		)
	})

	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s Identity: %s | Directory Listing: %v\n", timestamp, config.ServerName, config.EnableDir)
	fmt.Printf("%s Listening on :%s\n", timestamp, config.Port)

	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatalf("Critical Error: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request, config Config) {
	w.Header().Set("Server", config.ServerName)

	upath := filepath.Clean(r.URL.Path)
	path := filepath.Join(config.RootDir, filepath.FromSlash(upath))

	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if f.IsDir() {
		if !config.EnableDir {
			indexPath := filepath.Join(path, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				http.ServeFile(w, r, indexPath)
				return
			}
			http.NotFound(w, r)
			return
		}
	}

	http.FileServer(http.Dir(config.RootDir)).ServeHTTP(w, r)
}