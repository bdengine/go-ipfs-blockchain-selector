package selector

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

var c *cache.Cache

func TestGoCache(t *testing.T) {
	c = cache.New(2*time.Hour, 1*time.Hour)
	c.Add("testKey", "testValue", 1*time.Hour)
	get, b := c.Get("testKey")
	if !b {
		t.Fatal("应该找到\"testKey\"")
	}
	fmt.Println(get)
	get2, b2 := c.Get("testKey2")
	if b2 || get2 != nil {
		t.Fatal("不应该找到\"testKey2\"")
	}

}
