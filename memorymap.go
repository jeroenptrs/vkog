package main

import (
	"sync"

	"github.com/shamaton/msgpack/v2"
)

// MemoryMapData is used for serialization
type MemoryMapData struct {
	KV     map[string]string
	KVSize int
}

type MemoryMap struct {
	lock    sync.RWMutex
	kv      map[string]string
	kvSize  int
	maxSize int
}

func newMemoryMap(maxSize int) *MemoryMap {
	return &MemoryMap{
		kv:      make(map[string]string),
		kvSize:  0,
		maxSize: maxSize,
	}
}

func (vkog *MemoryMap) Set(key, value string) {
	vkog.lock.Lock()
	defer vkog.lock.Unlock()
	vkog.kv[key] = value
}

func (vkog *MemoryMap) Get(key string) string {
	vkog.lock.RLock()
	defer vkog.lock.RUnlock()
	return vkog.kv[key]
}

func (vkog *MemoryMap) Delete(key string) {
	vkog.lock.Lock()
	defer vkog.lock.Unlock()
	delete(vkog.kv, key)
}

func (vkog *MemoryMap) Compress() []byte {
	vkog.lock.RLock()
	defer vkog.lock.RUnlock()

	// Create a data structure with both the map and size
	data := MemoryMapData{
		KV:     vkog.kv,
		KVSize: vkog.kvSize,
	}

	d, err := msgpack.Marshal(data)
	if err != nil {
		panic(err)
	}
	return d
}

func (vkog *MemoryMap) Decompress(data []byte) {
	vkog.lock.Lock()
	defer vkog.lock.Unlock()

	// Create a structure to unmarshal into
	var mapData MemoryMapData

	err := msgpack.Unmarshal(data, &mapData)
	if err != nil {
		panic(err)
	}

	// Update the memory map with the unmarshaled data
	vkog.kv = mapData.KV
	vkog.kvSize = mapData.KVSize
}
