package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

type MemoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]string)}
}

func (m *MemoryStorage) Get(key string) (string, error) {
	v, ok := m.data[key]
	if !ok {
		return "", fmt.Errorf("key %q not found", key)
	}
	return v, nil
}

func (m *MemoryStorage) Set(key, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

type CachingStorage struct {
	next  Storage
	cache map[string]string
}

func WithCaching(next Storage) *CachingStorage {
	return &CachingStorage{next: next, cache: make(map[string]string)}
}

func (c *CachingStorage) Get(key string) (string, error) {
	if v, ok := c.cache[key]; ok {
		fmt.Printf("  [cache] hit  %q\n", key)
		return v, nil
	}
	fmt.Printf("  [cache] miss %q — fetching from storage\n", key)
	v, err := c.next.Get(key)
	if err == nil {
		c.cache[key] = v
	}
	return v, err
}

func (c *CachingStorage) Set(key, value string) error {
	delete(c.cache, key)
	return c.next.Set(key, value)
}

func (c *CachingStorage) Delete(key string) error {
	delete(c.cache, key)
	return c.next.Delete(key)
}

type LoggingStorage struct {
	next Storage
}

func WithLogging(next Storage) *LoggingStorage {
	return &LoggingStorage{next: next}
}

func (l *LoggingStorage) Get(key string) (string, error) {
	v, err := l.next.Get(key)
	if err != nil {
		log.Printf("[storage] Get(%q) error: %v", key, err)
	} else {
		log.Printf("[storage] Get(%q) = %q", key, v)
	}
	return v, err
}

func (l *LoggingStorage) Set(key, value string) error {
	err := l.next.Set(key, value)
	if err != nil {
		log.Printf("[storage] Set(%q) error: %v", key, err)
	} else {
		log.Printf("[storage] Set(%q, %q)", key, value)
	}
	return err
}

func (l *LoggingStorage) Delete(key string) error {
	err := l.next.Delete(key)
	log.Printf("[storage] Delete(%q)", key)
	return err
}

type ValidatingStorage struct {
	next Storage
}

func WithValidation(next Storage) *ValidatingStorage {
	return &ValidatingStorage{next: next}
}

func (v *ValidatingStorage) Get(key string) (string, error) {
	if err := validateKey(key); err != nil {
		return "", err
	}
	return v.next.Get(key)
}

func (v *ValidatingStorage) Set(key, value string) error {
	if err := validateKey(key); err != nil {
		return err
	}
	if strings.TrimSpace(value) == "" {
		return errors.New("value must not be empty")
	}
	return v.next.Set(key, value)
}

func (v *ValidatingStorage) Delete(key string) error {
	if err := validateKey(key); err != nil {
		return err
	}
	return v.next.Delete(key)
}

func validateKey(key string) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("key must not be empty")
	}
	return nil
}

func main() {
	store := WithValidation(
		WithLogging(
			WithCaching(
				NewMemoryStorage(),
			),
		),
	)

	fmt.Println("=== Set ===")
	store.Set("user:1", "alice")
	store.Set("user:2", "bob")

	fmt.Println("\n=== Get (cold cache) ===")
	store.Get("user:1")

	fmt.Println("\n=== Get (warm cache) ===")
	store.Get("user:1")

	fmt.Println("\n=== Set invalidates cache ===")
	store.Set("user:1", "alice updated")
	store.Get("user:1") // cache miss again

	fmt.Println("\n=== Validation errors ===")
	if err := store.Set("", "value"); err != nil {
		log.Printf("[validation] %v", err)
	}
	if err := store.Set("key", "  "); err != nil {
		log.Printf("[validation] %v", err)
	}
}
