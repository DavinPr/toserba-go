package api_test

import (
	"github.com/DavinPr/toserba-go/api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestHTTPAPIServer_Start(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	server := api.NewHTTPServer(api.ServerConfig{ListenAddr: ":8080"}, handler)
	defer server.Shutdown()

	go server.Start()
	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8080")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHTTPAPIServer_ShutdownGracefully(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	server := api.NewHTTPServer(api.ServerConfig{ListenAddr: ":8081"}, handler)

	go server.Start()
	time.Sleep(100 * time.Millisecond) // Give the server time to start

	err := server.Shutdown()
	assert.NoError(t, err)

	resp, err := http.Get("http://localhost:8081")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestHTTPAPIServer_ListenAndServeError(t *testing.T) {
	// Start a server on the same port to cause an error
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("handler", "handler-1")
		w.WriteHeader(http.StatusOK)
	})
	conflictingServer := api.NewHTTPServer(api.ServerConfig{ListenAddr: ":8082"}, handler1)
	defer conflictingServer.Shutdown()
	go conflictingServer.Start()
	time.Sleep(100 * time.Millisecond)

	// Start a server to test conflict
	handler2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	server := api.NewHTTPServer(api.ServerConfig{ListenAddr: ":8082"}, handler2)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Info().Msg("Recovered from panic")
			}
		}()
		server.Start()
	}()

	resp, err := http.Get("http://localhost:8082")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "handler-1", resp.Header.Get("handler"))
}
