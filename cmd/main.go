package main

import (
	"fmt"
	"github.com/sysatom/cache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "123",
	"Jack": "456",
	"Sam":  "789",
}

func main() {
	cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := ":5000"
	peers := cache.NewHTTPPool(addr)
	log.Println("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
