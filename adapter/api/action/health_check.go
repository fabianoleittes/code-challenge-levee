package action

import "net/http"

// HealthCheck return the APIs status
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
