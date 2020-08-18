package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/viewAllItem", viewAllItem).Methods("GET")
	router.HandleFunc("/addItem", addItem).Methods("POST")
	router.HandleFunc("/", test).Methods("GET")

	return router

}

// both test must come as fail
func TestViewAllItem(t *testing.T) {

	request, _ := http.NewRequest("GET", "/viewAllItem", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok response expected")
}
func TestTest(t *testing.T) {

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok response expected")
}
// func TestAdd(t *testing.T) {
// 	test2:=
// 	request, _ := http.NewRequest("POST", "/addItem", test2)
// 	response := httptest.NewRecorder()
// 	request.Header.Set("Content-Type", "application/json")
// 	Router().ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "Ok response expected")
// }
