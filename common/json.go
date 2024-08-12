package common

import (
	"encoding/json"
	"github.com/DavinPr/toserba-go/errors"
	"io"
	"net/http"
)

func JsonDecoder[T any](reader io.Reader, destination *T) error {
	if err := json.NewDecoder(reader).Decode(destination); err != nil {
		return err
	}
	return nil
}

func JsonDecoderFromHttpReq[T any](t *T, r *http.Request) error {
	err := JsonDecoder(r.Body, t)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
