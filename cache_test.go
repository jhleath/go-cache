package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(1 * time.Second)
	c.Store("test", 5)
	obj, stale := c.Get("test")
	if stale {
		t.Error("Cache is stale too early.")
	}
	if obj.(int) != 5 {
		t.Error("Cache didn't return correct value.")
	}
	time.Sleep(2 * time.Second)
	obj, stale = c.Get("test")
	if !stale {
		t.Error("Cache isn't alerting us to stale items.")
	}
	if obj.(int) != 5 {
		t.Error("Cache didn't return correct value.")
	}
}
