package cache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "123",
	"Jack": "456",
	"Sam":  "789",
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	g := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key] += 1
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	for k, v := range db {
		if view, err := g.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value of %s", k)
		}
		if _, err := g.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := g.Get("unknown"); err == nil {
		t.Fatalf("the value of unknown should be empty, but %s got", view)
	}
}
