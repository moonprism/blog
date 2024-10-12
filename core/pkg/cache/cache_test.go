package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := &Cache{
		items: make(map[string]*CacheItem),
		// 每秒cleanup用于测试
		cleanupInterval: time.Second,
	}
	cache.StartCleanup()

	// 表驱动测试
	tests := []struct {
		name string

		key     string
		value   any
		ttl     time.Duration
		wantOK  bool // set 状态
		sleep   time.Duration
		wantVal any           // get 值
		wantTTL time.Duration // 0:大于等于0, -1: -1, -2: -2
		wantLen int           // map长度
	}{
		{
			name:    "Basic set and get",
			key:     "foo",
			value:   "bar",
			ttl:     3 * time.Second,
			wantOK:  true,
			sleep:   0,
			wantVal: "bar",
			wantTTL: 0,
			wantLen: 1,
		},
		{
			name:    "Nil Set",
			key:     "ns",
			value:   nil,
			ttl:     3 * time.Second,
			wantOK:  false,
			sleep:   0,
			wantVal: nil,
			wantTTL: -2,
			wantLen: 1,
		},
		{
			name:    "Expired key",
			key:     "expk",
			value:   "no",
			ttl:     1 * time.Second,
			wantOK:  true,
			sleep:   4 * time.Second,
			wantVal: nil,
			wantTTL: -2,
			wantLen: 0,
		},
		{
			name:    "No TTL set",
			key:     "no-ttl",
			value:   "always",
			ttl:     0,
			wantOK:  true,
			sleep:   1 * time.Second,
			wantVal: "always",
			wantTTL: -1,
			wantLen: 1, // 永不过期
		},
		{
			name: "Struct value set",
			key:  "svs",
			value: struct {
				i int
			}{i: 1},
			ttl:    3 * time.Second,
			wantOK: true,
			sleep:  0,
			wantVal: struct {
				i int
			}{i: 1},
			wantTTL: 0,
			wantLen: 2,
		},
		{
			name:    "Number value set",
			key:     "nvs",
			value:   12.3,
			ttl:     10 * time.Second,
			wantOK:  true,
			sleep:   4 * time.Second,
			wantVal: 12.3,
			wantTTL: 0,
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Set method
			setOk := cache.Set(tt.key, tt.value, tt.ttl)
			if setOk != tt.wantOK {
				t.Errorf("Set() = %v; want %v", setOk, tt.wantOK)
			}
			if tt.sleep > 0 {
				time.Sleep(tt.sleep)
			}
			// Test Get method
			gotVal := cache.Get(tt.key)
			if gotVal != tt.wantVal {
				t.Errorf("Get() = %v; want %v", gotVal, tt.wantVal)
			}

			// Test TTL method
			gotTTL := cache.TTL(tt.key)
			if gotTTL >= 0 {
				gotTTL = 0
			}
			if gotTTL != tt.wantTTL {
				t.Errorf("TTL() = %v; want %v", gotTTL, tt.wantTTL)
			}

			// Test auto cleanup
			if len(cache.items) != tt.wantLen {
				t.Errorf("cleanup() len %v; want %v", len(cache.items), tt.wantLen)
			}
		})
	}
}
