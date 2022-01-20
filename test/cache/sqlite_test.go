package cache

import (
	"github.com/enspzr/getir-go-assigment/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This test method testing sqlite cache(in-memory) service.
func TestSqliteCacheGetSet(t *testing.T) {
	// Initialize sqlite cache for test.
	err := cache.InitSqliteCache()
	if err != nil {
		t.Fatal("Cache init error => ", err.Error())
	}

	// Define test cases and outputs.
	var newTest = []struct {
		input struct {
			key   string
			value string
		}
		output string
	}{
		{struct {
			key   string
			value string
		}{key: "key-1", value: "value-1"}, "value-1"},
		{struct {
			key   string
			value string
		}{key: "key-2", value: "value-2"}, "value-2"},
	}
	cs := cache.NewSqlCacheService()

	// Set cache test
	for _, tk := range newTest {
		t.Run("set cache:"+tk.input.key, func(t *testing.T) {
			err = cs.Set(tk.input.key, tk.input.value)
			if err != nil {
				t.Errorf("%s key set cache error => %s", tk.input.key, err.Error())
			}
		})

	}
	// Get cache test
	for _, tk := range newTest {
		t.Run("get cache:"+tk.input.key, func(t *testing.T) {
			val, err := cs.Get(tk.input.key)
			if err != nil {
				t.Errorf("%s key get cache error => %s", tk.input.key, err.Error())
			}
			if !assert.Equal(t, val, tk.output) {
				t.Errorf("got %v want %v", val, tk.output)
			}
		})
	}
}
