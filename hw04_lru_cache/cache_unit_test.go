package hw04lrucache

import (
	"testing"
)

func TestCache_BasicSetAndGet(t *testing.T) {
	cache := NewCache(2)

	cache.Set("a", 1)
	cache.Set("b", 2)

	if val, ok := cache.Get("a"); !ok || val != 1 {
		t.Errorf("Expected (1, true), got (%v, %v)", val, ok)
	}
	if val, ok := cache.Get("b"); !ok || val != 2 {
		t.Errorf("Expected (2, true), got (%v, %v)", val, ok)
	}
	if _, ok := cache.Get("c"); ok {
		t.Errorf("Expected (nil, false) for non-existent key")
	}
}

func TestCache_UpdateExistingKey(t *testing.T) {
	cache := NewCache(2)

	cache.Set("a", 1)
	wasPresent := cache.Set("a", 10)

	if !wasPresent {
		t.Error("Set should return true for existing key")
	}

	if val, ok := cache.Get("a"); !ok || val != 10 {
		t.Errorf("Expected updated value 10, got %v", val)
	}
}

func TestCache_EvictionOnOverflow(t *testing.T) {
	cache := NewCache(2)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	if _, ok := cache.Get("a"); ok {
		t.Error("Key 'a' should have been evicted")
	}
	if val, ok := cache.Get("b"); !ok || val != 2 {
		t.Errorf("Key 'b' should remain, got (%v, %v)", val, ok)
	}
	if val, ok := cache.Get("c"); !ok || val != 3 {
		t.Errorf("Key 'c' should be present, got (%v, %v)", val, ok)
	}
}

func TestCache_LRUOrdering(t *testing.T) {
	cache := NewCache(2)

	cache.Set("a", 1)
	cache.Set("b", 2)

	cache.Get("a")

	cache.Set("c", 3)

	if _, ok := cache.Get("b"); ok {
		t.Error("Key 'b' should have been evicted (least recently used)")
	}
	if val, ok := cache.Get("a"); !ok || val != 1 {
		t.Errorf("Key 'a' should remain, got (%v, %v)", val, ok)
	}
	if val, ok := cache.Get("c"); !ok || val != 3 {
		t.Errorf("Key 'c' should be present, got (%v, %v)", val, ok)
	}
}

func TestCache_Clear(t *testing.T) {
	cache := NewCache(2)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Clear()

	if val, ok := cache.Get("a"); ok {
		t.Errorf("After Clear, key 'a' should not exist, got %v", val)
	}
	if val, ok := cache.Get("b"); ok {
		t.Errorf("After Clear, key 'b' should not exist, got %v", val)
	}

	cache.Set("x", 42)
	if val, ok := cache.Get("x"); !ok || val != 42 {
		t.Errorf("After Clear, should be able to insert new data")
	}
}

func TestCache_ZeroCapacity(t *testing.T) {
	cache := NewCache(0)

	cache.Set("a", 1)
	if _, ok := cache.Get("a"); ok {
		t.Error("With capacity 0, nothing should be stored")
	}
}
