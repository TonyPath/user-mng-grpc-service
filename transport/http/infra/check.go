package infra

import (
	"database/sql"
	"net/http"

	// internal
	sqlrepo "github.com/TonyPath/user-mng-grpc-service/internal/repo/sql"
	thttp "github.com/TonyPath/user-mng-grpc-service/transport/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

type checkHandler struct {
	DB *sql.DB
}

func (h *checkHandler) Readiness(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := sqlrepo.StatusCheck(ctx, h.DB)
	if err != nil {
		thttp.Respond(w, statusResponse{
			Status: "db not ready",
		}, http.StatusInternalServerError)
		return
	}
}

func (h *checkHandler) Liveness(w http.ResponseWriter, r *http.Request) {
}
