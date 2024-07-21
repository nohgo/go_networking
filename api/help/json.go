// The help package contains helper functions
package help

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Decode struct takes a request with an http body and a struct it should decode to. It returns sanitized errors.
func DecodeStruct[F any](r *http.Request, dest *F) (error, int) {
	err := json.NewDecoder(r.Body).Decode(dest)

	if err == nil {
		return nil, 200
	}

	var syntaxError *json.SyntaxError
	var unmarshallTypeError *json.UnmarshalTypeError

	var msg string

	switch {
	case errors.As(err, &syntaxError), errors.Is(err, io.ErrUnexpectedEOF):
		msg = "Request body contains malformed json"
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		msg = "Unknown field in json"
	case errors.As(err, &unmarshallTypeError):
		msg = "Request body contains an invalid type"
	case errors.Is(err, io.EOF):
		msg = "Request body missing"
	default:
		msg = err.Error()
		return errors.New(msg), http.StatusInternalServerError
	}

	return errors.New(msg), http.StatusBadRequest

}

func SendJson(w http.ResponseWriter, target any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}
