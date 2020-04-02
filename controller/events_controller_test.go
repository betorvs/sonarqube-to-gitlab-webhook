package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostEvents(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sonarqube-to-gitlab-webhook/v1/events")

	// Assertions
	if assert.NoError(t, ReceiveEvents(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}
