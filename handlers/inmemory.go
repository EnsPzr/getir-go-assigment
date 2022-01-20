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

// InMemory
// This function redirects request by request method.
func InMemory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		successInMemoryResponse(w, nil)
	case "GET":
		inMemoryGet(w, r)
	case "POST":
		inMemoryPost(w, r)
	default:
		w.WriteHeader(405)
	}
}

// This function read value by key in go-cache.
// If key is empty, returns 400 status code.
// If key is not found, returns 404 status code.
// If key is exist, returns 200 status code, key and value.
// Request url: /in-memory?key=getir (GET)
func inMemoryGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		errorResponseInMemory(w, errKeyRequired, 400)
		return
	}
	val, ok := cache.Cache.Get(key)
	if !ok {
		errorResponseInMemory(w, errKeyNotFound, 404)
		return
	}
	vm := model.InMemory{
		Key:   key,
		Value: val.(string),
	}
	successInMemoryResponse(w, vm)
}

// This function set key and value in go-cache.
// If doesn't read body, returns 500 status code with error message.
// If doesn't bind body in model, returns 500 status code with error message.
// If validation condition not true, returns 400 status code with validation errors.
// If is succesfuly, returns 200 status code with request body.
// Request url: /in-memory (POST)
// Request model: { "key":"test", "value":"getir" }
func inMemoryPost(w http.ResponseWriter, r *http.Request) {
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

	cache.Cache.Set(vm.Key, vm.Value)
	successInMemoryResponse(w, vm)
}
