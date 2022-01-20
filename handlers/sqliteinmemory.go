package handlers

import (
	"encoding/json"
	"errors"
	"github.com/enspzr/getir-go-assigment/cache"
	"github.com/enspzr/getir-go-assigment/model"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	errKeyRequired = errors.New("key is required")
	errKeyNotFound = errors.New("key is not found")
)

// InMemorySqlite
// This function redirect request by request method.
func InMemorySqlite(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		successInMemoryResponse(w, nil)
	case "GET":
		inMemorySqliteGet(w, r)
	case "POST":
		inMemorySqlitePost(w, r)
	default:
		w.WriteHeader(405)
	}
}

// This function read value according to key in sqlite-cache.
// If key is empty, returns 400 status code.
// If exist any error when it was taken from cache, returns 500 status code with error message.
// If key is not found, returns 404 status code.
// If key is exist, returns 200 status code, key and value.
// Request url: /in-memory-sqlite?key=getir (GET)
func inMemorySqliteGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		errorResponseInMemory(w, errKeyRequired, 400)
		return
	}
	cs := cache.NewSqlCacheService()
	// Get value by key
	val, err := cs.Get(key)
	if err != nil {
		errorResponseInMemory(w, err, 500)
		return
	}
	if val == "" {
		errorResponseInMemory(w, errKeyNotFound, 404)
		return
	}
	vm := model.InMemory{
		Key:   key,
		Value: val,
	}
	successInMemoryResponse(w, vm)
}

// This function set key and value in sqlite-cache.
// If doesn't read body, returns 500 status code with error message.
// If doesn't bind body in model, returns 500 status code with error message.
// If validation condition not true, returns 400 status code with validation errors.
// If operation is successful, returns 200 status code with request body.
// Request url: /in-memory-sqlite (POST)
// Request model: { "key":"test", "value":"getir" }
func inMemorySqlitePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponseInMemory(w, errors.New("Body couldn't be read => "+err.Error()), 500)
		return
	}
	var vm model.InMemory
	err = json.Unmarshal(body, &vm)
	if err != nil {
		errorResponseInMemory(w, errors.New("Form failed to bind to model => "+err.Error()), 500)
		return
	}

	if errs := vm.Validate(); len(errs) > 0 {
		errorResponseInMemory(w, errors.New(strings.Join(errs, "\n")), 400)
		return
	}
	cs := cache.NewSqlCacheService()
	err = cs.Set(vm.Key, vm.Value)
	if err != nil {
		errorResponseInMemory(w, err, 500)
		return
	}
	successInMemoryResponse(w, vm)
}
