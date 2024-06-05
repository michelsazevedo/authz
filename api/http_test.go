package api

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/michelsazevedo/authz/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockService  = domain.NewMockService()
	authzHandler = NewHandler(mockService)
)

func TestSignInHandler(t *testing.T) {
	e := echo.New()
	var jsonPayload string

	t.Run("Returns Status Code 400", func(t *testing.T) {
		jsonPayload = `{
			"email": "peter.parker@marvel.com",
			"password": 111
		}`

		req := httptest.NewRequest(http.MethodPost, "/users/signin", bytes.NewBufferString(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, authzHandler.SignIn(c)) {
			if status := rec.Code; status != http.StatusBadRequest {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusBadRequest, status)
			}
		}
	})

	t.Run("Returns Status Code 401", func(t *testing.T) {
		jsonPayload = `{
			"email": "peter.parker@marvel.com",
			"password": "password123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/users/signin", bytes.NewBufferString(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, authzHandler.SignIn(c)) {
			if status := rec.Code; status != http.StatusUnauthorized {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusUnauthorized, status)
			}
		}
	})

	t.Run("Returns Status Code 201", func(t *testing.T) {
		jsonPayload = `{
			"email": "johndoe@gmail.com",
			"password": "password123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/users/signin", bytes.NewBufferString(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, authzHandler.SignIn(c)) {
			if status := rec.Code; status != http.StatusCreated {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusCreated, status)
			}
		}
	})
}

func TestSignUpHandler(t *testing.T) {
	e := echo.New()
	var jsonPayload string

	t.Run("Returns Status Code 400", func(t *testing.T) {
		jsonPayload = `{
		    "password": 111
		}`

		req := httptest.NewRequest(http.MethodPost, "/users/signup", bytes.NewBufferString(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, authzHandler.SignUp(c)) {
			if status := rec.Code; status != http.StatusBadRequest {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusBadRequest, status)
			}
		}
	})

	t.Run("Returns Status Code 201", func(t *testing.T) {
		jsonPayload = `{
		    "first_name": "Harry",
		    "last_name": "Osborn",
		    "email": "harry.osborn@marvel.com",
		    "locale": "Americas/Sao_Paolo",
		    "password": "StrongPassword"
		}`

		req := httptest.NewRequest(http.MethodPost, "/users/signup", bytes.NewBufferString(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, authzHandler.SignUp(c)) {
			if status := rec.Code; status != http.StatusCreated {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusCreated, status)
			}
		}
	})
}

func TestAuthHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Returns Status Code 201", func(t *testing.T) {
		if assert.NoError(t, authzHandler.Auth(c)) {
			if status := rec.Code; status != http.StatusCreated {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusCreated, status)
			}
		}
	})
}

func TestRefreshHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/refresh", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Returns Status Code 201", func(t *testing.T) {
		if assert.NoError(t, authzHandler.Refresh(c)) {
			if status := rec.Code; status != http.StatusCreated {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusCreated, status)
			}
		}
	})
}
