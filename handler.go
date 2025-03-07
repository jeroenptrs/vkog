package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handler(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	path := request.URL.Path

	// TEMP - Path should be in format /k/<keyname> - FUTURE: we will support other functions like syncing which won't be under /k/
	if !strings.HasPrefix(path, "/k/") {
		http.Error(response, "Invalid path. Use /k/<keyname>", http.StatusBadRequest)
		return
	}

	// Extract key from path
	key := strings.TrimPrefix(path, "/k/")
	if key == "" {
		http.Error(response, "Key cannot be empty", http.StatusBadRequest)
		return
	}

	switch method {
	case "GET":
		eventRegistry.beforeGet(key)

		value := store.Get(key)
		if value == "" {
			http.Error(response, "Not Found", http.StatusNotFound)
		} else {
			fmt.Fprintf(response, "%s", value)
			eventRegistry.afterGet(key, value)
		}

	case "POST":
		// Read value from request body
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(response, "Error reading request body", http.StatusInternalServerError)
			return
		}
		value := string(body)
		eventRegistry.beforePost(key, value)

		if value == "" {
			http.Error(response, "No content provided", http.StatusBadRequest)
		} else {
			ok, reason, reasonHttpCode := eventRegistry.guardPost(key, value)
			if !ok {
				http.Error(response, reason, reasonHttpCode)
				return
			}

			store.Set(key, value)
			fmt.Fprintf(response, "Ok")
			eventRegistry.afterPost(key, value)
		}

	case "PUT":
		// Read value from request body
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(response, "Error reading request body", http.StatusInternalServerError)
			return
		}
		value := string(body)
		eventRegistry.beforePut(key, value)

		if value == "" {
			http.Error(response, "No content provided", http.StatusBadRequest)
		} else {
			oldValue := store.Get(key)

			ok, reason, reasonHttpCode := eventRegistry.guardPut(key, oldValue, value)
			if !ok {
				http.Error(response, reason, reasonHttpCode)
				return
			}

			store.Set(key, value)
			fmt.Fprintf(response, "Ok")
			eventRegistry.afterPut(key, oldValue, value)
		}

	case "DELETE":
		value := store.Get(key)
		eventRegistry.beforeDelete(key, value)

		if value == "" {
			http.Error(response, "Not Found", http.StatusNotFound)
		} else {
			store.Delete(key)
			fmt.Fprintf(response, "Ok")
			eventRegistry.afterDelete(key, value)
		}

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
