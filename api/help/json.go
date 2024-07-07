package help

import (
	"encoding/json"
	"net/http"
)

func DecodeStruct[F any](r *http.Request, dest *F) (error, int) {
	err := json.NewDecoder(r.Body).Decode(dest)

	if err != nil {
		return err, 500
	}

	return nil, 200
}

func SendJson(w http.ResponseWriter, target any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}
