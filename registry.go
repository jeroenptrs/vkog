package main

import "fmt"

type Registry struct {
	plugins []Plugin
}

func newRegistry() *Registry {
	return &Registry{
		plugins: make([]Plugin, 0),
	}
}

func (registry *Registry) Register(plugin Plugin) {
	registry.plugins = append(registry.plugins, plugin)
	name := plugin.getName()
	fmt.Printf("Registered: %s\n", name)
}

func (registry *Registry) beforeGet(key string) {
	for _, plugin := range registry.plugins {
		if beforeGet, ok := plugin.(beforeGet); ok {
			beforeGet.beforeGet(key)
		}
	}
}

func (registry *Registry) afterGet(key string, value string) {
	for _, plugin := range registry.plugins {
		if afterGet, ok := plugin.(afterGet); ok {
			afterGet.afterGet(key, value)
		}
	}
}

func (registry *Registry) beforePost(key string, value string) {
	for _, plugin := range registry.plugins {
		if beforePost, ok := plugin.(beforePost); ok {
			beforePost.beforePost(key, value)
		}
	}
}

func (registry *Registry) guardPost(key string, value string) (bool, string, int) {
	for _, plugin := range registry.plugins {
		if guardPost, ok := plugin.(guardPost); ok {
			ok, reason, reasonHttpCode := guardPost.guardPost(key, value)
			if !ok {
				return false, reason, reasonHttpCode
			}
		}
	}

	return true, "", 0
}

func (registry *Registry) afterPost(key string, value string) {
	for _, plugin := range registry.plugins {
		if afterPost, ok := plugin.(afterPost); ok {
			afterPost.afterPost(key, value)
		}
	}
}

func (registry *Registry) beforePut(key string, value string) {
	for _, plugin := range registry.plugins {
		if beforePut, ok := plugin.(beforePut); ok {
			beforePut.beforePut(key, value)
		}
	}
}

func (registry *Registry) guardPut(key string, oldValue string, newValue string) (bool, string, int) {
	for _, plugin := range registry.plugins {
		if guardPut, ok := plugin.(guardPut); ok {
			ok, reason, reasonHttpCode := guardPut.guardPut(key, oldValue, newValue)
			if !ok {
				return false, reason, reasonHttpCode
			}
		}
	}

	return true, "", 0
}

func (registry *Registry) afterPut(key string, oldValue string, newValue string) {
	for _, plugin := range registry.plugins {
		if afterPut, ok := plugin.(afterPut); ok {
			afterPut.afterPut(key, oldValue, newValue)
		}
	}
}

func (registry *Registry) beforeDelete(key string, value string) {
	for _, plugin := range registry.plugins {
		if beforeDelete, ok := plugin.(beforeDelete); ok {
			beforeDelete.beforeDelete(key, value)
		}
	}
}

func (registry *Registry) afterDelete(key string, value string) {
	for _, plugin := range registry.plugins {
		if afterDelete, ok := plugin.(afterDelete); ok {
			afterDelete.afterDelete(key, value)
		}
	}
}
