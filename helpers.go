package main

import (
	"fmt"
	"os"
)

func storeVkog(filename string) {
	compressed := store.Compress()
	os.WriteFile(filename, compressed, 0644)
	fmt.Printf("Stored %d bytes, with %d bytes of kv data to %s\n", len(compressed), store.kvSize, filename)
}

func loadVkog(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		return
	}

	store.Decompress(data)
	fmt.Printf("Loaded a total of %d bytes, with %d bytes of kv data from %s\n", len(data), store.kvSize, filename)
}

func handleExit(file string) {
	fmt.Printf("Flushing vkog: writing %d byte(s) of kv data to %s\n", store.kvSize, file)
	storeVkog(file)
	fmt.Println("Flushed vkog successfully")
}
