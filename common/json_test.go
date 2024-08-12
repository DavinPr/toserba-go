package common_test

import (
	"github.com/DavinPr/toserba-go/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonDecoder(t *testing.T) {
	body := `{"name":"test"}`

	dest := make(map[string]string)

	err := common.JsonDecoder(strings.NewReader(body), &dest)

	assert.NoError(t, err)
	assert.Equal(t, "test", dest["name"])
}

func TestJsonDecoder_WrongJson(t *testing.T) {
	body := `{"name":"test"`

	dest := make(map[string]string)

	err := common.JsonDecoder(strings.NewReader(body), &dest)

	assert.Error(t, err)
}

func TestJsonDecoderFromHttpReq(t *testing.T) {
	body := `{"name":"test"}`

	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	dest := make(map[string]string)

	err := common.JsonDecoderFromHttpReq(&dest, r)

	assert.NoError(t, err)
	assert.Equal(t, "test", dest["name"])
}

func TestJsonDecoderFromHttpReq_WrongJsons(t *testing.T) {
	body := `{"name":"test"`

	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	dest := make(map[string]string)

	err := common.JsonDecoderFromHttpReq(&dest, r)

	assert.Error(t, err)
}
