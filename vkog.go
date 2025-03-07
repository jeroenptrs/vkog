package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var store *MemoryMap
var eventRegistry *Registry

const VERSION = "0.0.1"
const DEFAULT_MAX_SIZE = 1024 * 1024 * 128 // 128MB
const DEFAULT_VKOG_FILE = "local.vkog"

func main() {
	var size int
	var file string
	flag.IntVar(&size, "s", DEFAULT_MAX_SIZE, "The maximum size of the vkog file")
	flag.StringVar(&file, "f", DEFAULT_VKOG_FILE, "The file to store the vkog data")
	flag.Parse()

	fmt.Printf("vkog: %s\n", VERSION)
	fmt.Printf("Starting with max size %d bytes of kv data from file %s\n", size, file)

	// Initialize vkog kv
	store = newMemoryMap(size)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist, starting fresh\n", file)
	} else {
		loadVkog(file)
	}

	// Initialize vkog plugins
	eventRegistry = newRegistry()
	eventRegistry.Register(&ConflictPlugin{})
	eventRegistry.Register(&BytesPlugin{})

	// Initialize http server
	fmt.Println("Starting vkog on port 3110")
	server := &http.Server{
		Addr:    ":3110",
		Handler: http.HandlerFunc(handler),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-done
	handleExit(file)

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(context); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	} else {
		fmt.Println("vkog exiting")
	}
}
