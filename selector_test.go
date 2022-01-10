package selector

import (
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func TestGoCache(t *testing.T) {
	c := cache.New(2*time.Hour, 1*time.Hour)
	err := c.Add("testKey", "testValue", 1*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	get, b := c.Get("testKey")
	if !b {
		t.Fatal("应该找到\"testKey\"")
	}
	if get != "testValue" {
		t.Fatal("值获取错误")
	}
	get2, b2 := c.Get("testKey2")
	if b2 || get2 != nil {
		t.Fatal("不应该找到\"testKey2\"")
	}
}
