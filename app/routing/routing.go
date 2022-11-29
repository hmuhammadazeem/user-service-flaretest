package routing

import (
	"net/http"
	"encoding/json"

	"github.com/hmuhammadazeem/user-service/app/health"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := health.Health()
		if err {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		json.NewEncoder(w).Encode(&result)
	}
}