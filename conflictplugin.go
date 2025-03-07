package main

import (
	"net/http"
)

type ConflictPlugin struct{}

func (p *ConflictPlugin) getName() string {
	return "conflict.plugin"
}

func (p *ConflictPlugin) guardPost(key string, value string) (bool, string, int) {
	storeValue := store.Get(key)
	if storeValue != "" {
		return false, "Key already exists", http.StatusConflict
	}
	return true, "", 0
}

func (p *ConflictPlugin) guardPut(key string, oldValue string, newValue string) (bool, string, int) {
	if oldValue == "" {
		return false, "Key does not exist", http.StatusNotFound
	}
	return true, "", 0
}
