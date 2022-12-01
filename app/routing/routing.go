package routing

import (
	"net/http"
	"encoding/json"

	"github.com/hmuhammadazeem/user-service/app/health"
	"github.com/hmuhammadazeem/user-service/app/user"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := health.Health()
		if err {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(&result)
	}
}


func CheckUsernameHandler(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}