package pokecache

import (
	"testing"
	"bytes"
)

func TestCache(t *testing.T) {
	var cache Cache
	cache.NewCache()

	if cache.Data == nil {
		t.Fatal("expected cache.Data to be initialized after NewCache()")
	}

	t.Run("missing key returns nil and false", func(t *testing.T) {
		val, ok := cache.Get("missing")
		if ok {
			t.Fatalf("expected ok=false for missing key, got true")
		}
		if val != nil {
			t.Fatalf("expected nil value for missing key, got %v", val)
		}
	})

	t.Run("add and retrieve multiple keys in the same cache", func(t *testing.T) {
		entries := map[string][]byte{
			"user:1": []byte("alice"),
			"user:2": []byte("bob"),
			"user:3": []byte("charlie"),
		}

		for k, v := range entries {
			cache.Add(k, v)
		}

		for k, expected := range entries {
			got, ok := cache.Get(k)
			if !ok {
				t.Fatalf("expected key %q to be present, got ok=false", k)
			}
			if !bytes.Equal(got, expected) {
				t.Errorf("for key %q: expected %q, got %q", k, expected, got)
			}
		}
	})

	t.Run("replace existing value without overwriting other keys", func(t *testing.T) {
		updated := []byte("robert")
		cache.Add("user:2", updated)

		got, ok := cache.Get("user:2")
		if !ok {
			t.Fatalf("expected key %q to be present after update", "user:2")
		}
		if !bytes.Equal(got, updated) {
			t.Errorf("expected updated value %q, got %q", updated, got)
		}

		if v1, ok1 := cache.Get("user:1"); !ok1 || !bytes.Equal(v1, []byte("alice")) {
			t.Errorf("expected user:1 to remain %q, got %q (ok=%v)", "alice", v1, ok1)
		}
		if v3, ok3 := cache.Get("user:3"); !ok3 || !bytes.Equal(v3, []byte("charlie")) {
			t.Errorf("expected user:3 to remain %q, got %q (ok=%v)", "charlie", v3, ok3)
		}
	})

	t.Run("missing key still returns nil and false after additions", func(t *testing.T) {
		val, ok := cache.Get("user:999")
		if ok || val != nil {
			t.Errorf("expected (nil, false) for missing key, got (%v, %v)", val, ok)
		}
	})
}
