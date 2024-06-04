package api

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	authzHandler = NewHandler()
)

func TestSignInHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/signup", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Returns Status Code 200", func(t *testing.T) {
		if assert.NoError(t, authzHandler.SignIn(c)) {
			if status := rec.Code; status != http.StatusOK {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusOK, status)
			}
		}
	})
}

func TestSignUpHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/signup", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Returns Status Code 201", func(t *testing.T) {
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

	t.Run("Returns Status Code 200", func(t *testing.T) {
		if assert.NoError(t, authzHandler.Auth(c)) {
			if status := rec.Code; status != http.StatusOK {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
					http.StatusOK, status)
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
