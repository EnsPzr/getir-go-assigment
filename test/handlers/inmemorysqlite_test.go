package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/enspzr/getir-go-assigment/cache"
	"github.com/enspzr/getir-go-assigment/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

// This test method testing sqlite cache(in-memory) handlers.
func TestInMemorySqlite(t *testing.T) {
	// Initialize sqlite cache service.
	err := cache.InitSqliteCache()
	if err != nil {
		t.Fatal("Cache init error => ", err.Error())
	}
	// Setup router.
	http.HandleFunc("/in-memory-sqlite", handlers.InMemorySqlite)
	// Start http server.
	go func() {
		err = http.ListenAndServe(":8082", nil)
		if err != nil {
			t.Fatal("Http server starting error => " + err.Error())
		}
	}()

	// Test inputs and outputs.
	var newTest = []struct {
		input struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		output string
	}{
		{struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{Key: "key-1", Value: "value-1"}, "value-1"},
		{struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{Key: "key-2", Value: "value-2"}, "value-2"},
	}

	type response struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	// Test set endpoint.
	for _, tk := range newTest {
		t.Run("Handlers-set-sqlite cache:"+tk.input.Key, func(t *testing.T) {
			bodyJson, err := json.Marshal(tk.input)
			if err != nil {
				t.Errorf("Body json marshall error => %s", err.Error())
			}
			resp, err := http.Post("http://localhost:8082/in-memory-sqlite",
				"application/json", bytes.NewBuffer(bodyJson))
			if err != nil {
				t.Errorf("Post Request Error => %s", err.Error())
			}

			if resp.StatusCode != 200 {
				t.Errorf("Status code different 200")
			}
			defer resp.Body.Close()
		})
	}

	// Test get endpoint.
	for _, tk := range newTest {
		t.Run("Handlers-get-sqlite cache:"+tk.input.Key, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8082/in-memory-sqlite?key=%s", tk.input.Key))
			if err != nil {
				t.Errorf("Post Request Error => %s", err.Error())
			}

			if resp.StatusCode != 200 {
				t.Errorf("Status code different 200")
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Reading body error => " + err.Error())
			}

			var result response
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("Json unmarshall error => " + err.Error())
			}

			if !assert.Equal(t, result.Value, tk.output) {
				t.Errorf("got %v want %v", result.Value, tk.output)
			}
		})
	}
}
