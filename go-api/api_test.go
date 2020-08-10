package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Test).Methods("GET")
	return router

}

// both test must come as fail
func ViewAll(t *testing.T) {
	fmt.Println("asdasd")
	request, _ := http.NewRequest("GET", "/viewAllItem", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	
	assert.Equal(t, 400, response.Code, "Expect to see items from array")
}

func add(t *testing.T) {
	fmt.Println("asdasd")
	request, _ := http.NewRequest("POST", "/addItem", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	
	assert.Equal(t, 400, response.Code, "Expect to see items from array")
}