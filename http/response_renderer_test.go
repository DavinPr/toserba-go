package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DavinPr/toserba-go/errors"
	localHttp "github.com/DavinPr/toserba-go/http"
	"github.com/stretchr/testify/assert"
)

func TestRenderResponse(t *testing.T) {
	w := httptest.NewRecorder()
	response := map[string]string{"message": "success"}
	localHttp.RenderResponse(w, response, http.StatusOK)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "success"}`, w.Body.String())
}

func TestRenderErrorResponse_WhenDefaultError(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.New("client error")

	localHttp.RenderErrorResponse(r, w, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"success":false,"errors": [{"code":"SERVER_ERROR","message_title":"Internal Server Error","message":"Some error occurred, please try again"}]}`, w.Body.String())
}

func TestRenderErrorResponse_WithBadRequestError(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.NewBadRequestError("client error")

	localHttp.RenderErrorResponse(r, w, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"success":false,"errors": [{"code":"BAD_REQUEST","message_title":"Bad Request","message":"client error"}]}`, w.Body.String())
}

func TestRenderErrorResponse_WithNotFoundError(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.NewNotFoundError("client error")

	localHttp.RenderErrorResponse(r, w, err)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, `{"success":false,"errors": [{"code":"NOT_FOUND","message_title":"Not Found","message":"data not found"}]}`, w.Body.String())
}

func TestRenderErrorResponse_WithUnprocessableError(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.NewUnprocessableEntityError("client error")

	localHttp.RenderErrorResponse(r, w, err)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.JSONEq(t, `{"success":false,"errors": [{"code":"INVALID_REQUEST","message_title":"Please Input Valid Request","message":"client error"}]}`, w.Body.String())
}

func TestRenderErrorResponse_WithUnauthorizedError(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.NewUnauthorizedError("client error")

	localHttp.RenderErrorResponse(r, w, err)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"success":false,"errors": [{"code":"INVALID_AUTHORIZATION","message_title":"You're not authorized","message":"client error"}]}`, w.Body.String())
}
