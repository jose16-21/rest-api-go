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
	// Configura el router para pruebas
	initDB()
	r := gin.Default()
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	return r
}

func TestGetUsers(t *testing.T) {
	r := setupRouter()

	// Crear una solicitud GET para /users
	req, _ := http.NewRequest("GET", "/users?page=1&limit=2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificar el contenido de la respuesta
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "data")
	assert.Contains(t, response, "total")
	assert.Contains(t, response, "page")
	assert.Contains(t, response, "limit")
}

func TestCreateUser(t *testing.T) {
	r := setupRouter()

	// Crear un cuerpo de solicitud válido
	user := map[string]string{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}
	body, _ := json.Marshal(user)

	// Crear una solicitud POST para /users
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verificar el contenido de la respuesta
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", response["name"])
	assert.Equal(t, "john.doe@example.com", response["email"])
}
