package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a router for testing
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	return router
}

func TestGetAlbums(t *testing.T) {
	// Setup
	router := setupRouter()

	// Create test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []album
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, len(albums), len(response))
	assert.Equal(t, albums[0].ID, response[0].ID)
	assert.Equal(t, albums[0].Title, response[0].Title)
}

func TestGetAlbumByID(t *testing.T) {
	// Setup
	router := setupRouter()

	// Test cases
	testCases := []struct {
		name         string
		id           string
		expectedCode int
		checkBody    bool
	}{
		{
			name:         "valid album",
			id:           "1",
			expectedCode: http.StatusOK,
			checkBody:    true,
		},
		{
			name:         "album not found",
			id:           "999",
			expectedCode: http.StatusNotFound,
			checkBody:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/albums/"+tc.id, nil)
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tc.expectedCode, w.Code)

			if tc.checkBody {
				// Parse response
				var response album
				err := json.Unmarshal(w.Body.Bytes(), &response)

				// Assert
				assert.Nil(t, err)
				assert.Equal(t, tc.id, response.ID)
			}
		})
	}
}

func TestPostAlbums(t *testing.T) {
	// Setup
	router := setupRouter()
	originalLength := len(albums)

	// Create test payload
	newAlbum := album{
		ID:     "4",
		Title:  "Giant Steps",
		Artist: "John Coltrane",
		Price:  63.99,
	}
	payload, _ := json.Marshal(newAlbum)

	// Create test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse response
	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, newAlbum.ID, response.ID)
	assert.Equal(t, newAlbum.Title, response.Title)
	assert.Equal(t, originalLength+1, len(albums))

	// Verify album was actually added to the collection
	found := false
	for _, a := range albums {
		if a.ID == newAlbum.ID {
			found = true
			break
		}
	}
	assert.True(t, found)
}

func TestPostAlbumsInvalidJSON(t *testing.T) {
	// Setup
	router := setupRouter()
	originalLength := len(albums)

	// Create invalid test payload
	invalidJSON := []byte(`{"id": "4", "title": "Giant Steps", "artist":}`)

	// Create test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Assert status code (should indicate bad request)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert albums length unchanged
	assert.Equal(t, originalLength, len(albums))
}
