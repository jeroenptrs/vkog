package main

import (
	"fmt"
	"net/http"
)

func byteCalc(key string, value string) int {
	return len(key) + len(value)
}

type BytesPlugin struct{}

func (p *BytesPlugin) getName() string {
	return "bytes.plugin"
}

func (p *BytesPlugin) guardPost(key string, value string) (bool, string, int) {
	var size = byteCalc(key, value)
	if store.kvSize+size > store.maxSize {
		return false, "Max size exceeded", http.StatusRequestEntityTooLarge
	}
	return true, "", 0
}

func (p *BytesPlugin) afterPost(key, value string) {
	var size = byteCalc(key, value)
	store.kvSize += size
	fmt.Printf("Memory size increased by %d bytes\n", size)
}

func (p *BytesPlugin) guardPut(key string, oldValue string, newValue string) (bool, string, int) {
	// Calculate the size difference
	var oldSize = byteCalc(key, oldValue)
	var newSize = byteCalc(key, newValue)
	var sizeDiff = newSize - oldSize

	// Check if the new size exceeds the maximum
	if store.kvSize+sizeDiff > store.maxSize {
		return false, "Max size exceeded", http.StatusRequestEntityTooLarge
	}
	return true, "", 0
}

func (p *BytesPlugin) afterPut(key string, oldValue string, newValue string) {
	var oldSize = byteCalc(key, oldValue)
	var newSize = byteCalc(key, newValue)
	store.kvSize += newSize - oldSize
	fmt.Printf("Memory size increased by %d bytes\n", newSize-oldSize)
}

func (p *BytesPlugin) afterDelete(key string, value string) {
	var size = byteCalc(key, value)
	store.kvSize -= size
	fmt.Printf("Memory size decreased by %d bytes\n", size)
}
