package config

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register(struct{}{})
}

type Cache struct {
	ttl time.Duration
}

type cacheItem struct {
	Value     interface{}
	Timestamp time.Time
}

func Set(key string, value interface{}, duration time.Duration) bool {
	item := cacheItem{
		Value:     value,
		Timestamp: time.Now().Add(duration),
	}

	filename := "runtime/cache/" + key + ".bin"

	file, err := os.Create(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(item); err != nil {
		return false
	}

	return true
}

func Get(key string) (interface{}, bool) {
	filename := "runtime/cache/" + key + ".bin"

	file, err := os.Open(filename)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var item cacheItem
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&item); err != nil {
		return nil, false
	}

	itemTime := item.Timestamp.Format("2006-01-02 15:04:05")
	now := time.Now().Format("2006-01-02 15:04:05")

	if now > itemTime {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) error {

	filename := "runtime/cache/" + key + ".bin"

	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("faylni o'chirishda xatolik: %v", err)
	}

	return nil
}

func (c *Cache) CleanUp() error {
	cacheDir := "runtime/cache/"

	files, err := filepath.Glob(cacheDir + "*.bin")
	if err != nil {
		return fmt.Errorf("katalogni o'qishda xatolik: %v", err)
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			continue
		}

		var item cacheItem
		decoder := gob.NewDecoder(f)
		if err := decoder.Decode(&item); err != nil {
			f.Close()
			continue
		}
		f.Close()

		if time.Since(item.Timestamp) > c.ttl {
			os.Remove(file)
		}
	}

	return nil
}
