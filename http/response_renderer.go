package http

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/DavinPr/toserba-go/constants"
	"github.com/DavinPr/toserba-go/errors"
)

func RenderResponse(w http.ResponseWriter, successResponse interface{}, statusCode int) {
	writeResponse(w, successResponse, statusCode)
}

func RenderErrorResponse(r *http.Request, w http.ResponseWriter, err error) {
	statusCode, errResponse := getErrorConfig(err)
	logFunc := func(logEvent *zerolog.Event, msg string) {
		logEvent.Int("ResponseCode", statusCode).
			Interface("Request", map[string]string{
				"Method": r.Method,
				"Host":   r.Host,
				"Path":   r.URL.Path,
			}).
			Msg(msg)
	}
	if statusCode < 500 {
		logFunc(log.Warn(), err.Error())
	} else {
		logFunc(log.Error().Err(err), "Internal server error")
	}
	writeResponse(w, errResponse, statusCode)
}

func writeResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set(constants.HeaderContentType, "application/json")
	w.WriteHeader(statusCode)
	if response != nil {
		body, err := json.Marshal(response)
		if err == nil {
			w.Write(body)
		}
	}
}

func getErrorConfig(err error) (int, ServiceResponse) {
	switch e := err.(type) {
	case *errors.BadRequestError:
		return e.Code(), NewBadRequestError(err.Error())
	case *errors.NotFoundError:
		return e.Code(), NewNotFoundError()
	case *errors.UnprocessableEntityError:
		return e.Code(), NewUnprocessableEntityError(err.Error())
	case *errors.UnauthorizedError:
		return e.Code(), NewUnauthorizedError(err.Error())
	default:
		return http.StatusInternalServerError, NewInternalServerError()
	}
}
