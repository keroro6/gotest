package miniredis

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestDoSomethingWithRedis(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.Set("q1mi", "liwenzhou.com")
	s.SAdd(KeyValidWebsite, "q1mi")

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})

	ok := DoSomethingWithRedis(rdb, "q1mi")
	if !ok {
		t.Fatal()
	}
	if got, err := s.Get("blog"); err != nil || got != "https://liwenzhou.com" {
		t.Fatalf("'blog' has the wrong value")
	}

	s.CheckGet(t, "blog", "https://liwenzhou.com")

	s.FastForward(5 * time.Second)
	if s.Exists("blog") {
		t.Fatal("'blog' should not have existed anymore")
	}
}
