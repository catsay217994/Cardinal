package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_GetAllBulletins(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/manager/bulletins", nil)
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestService_NewBulletin(t *testing.T) {
	// JSON bind fail
	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(map[string]interface{}{
		"Title": "this is a bulletin",
	})
	req, _ := http.NewRequest("POST", "/manager/bulletin", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// success
	w = httptest.NewRecorder()
	jsonData, _ = json.Marshal(map[string]interface{}{
		"Title":   "this is a bulletin",
		"Content": "test test test",
	})
	req, _ = http.NewRequest("POST", "/manager/bulletin", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestService_EditBulletin(t *testing.T) {
	// JSON bind fail
	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(map[string]interface{}{
		"Title":   "this is a bulletin",
		"Content": "new content",
	})
	req, _ := http.NewRequest("PUT", "/manager/bulletin", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// not found
	w = httptest.NewRecorder()
	jsonData, _ = json.Marshal(map[string]interface{}{
		"ID":      2,
		"Title":   "this is a bulletin",
		"Content": "new content",
	})
	req, _ = http.NewRequest("PUT", "/manager/bulletin", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)

	// success
	w = httptest.NewRecorder()
	jsonData, _ = json.Marshal(map[string]interface{}{
		"ID":      1,
		"Title":   "this is a bulletin",
		"Content": "new content",
	})
	req, _ = http.NewRequest("PUT", "/manager/bulletin", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestService_DeleteBulletin(t *testing.T) {
	// error id
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/manager/bulletin?id=asdfg", nil)
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// id not exist
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/manager/bulletin?id=233", nil)
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)

	// success
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/manager/bulletin?id=1", nil)
	req.Header.Set("Authorization", managerToken)
	service.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
