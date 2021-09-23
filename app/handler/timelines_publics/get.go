package timelines_publics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	daoTimelinesPublic := h.app.Dao.TimelinesPublic()
	timelinesPublic, err := daoTimelinesPublic.FetchAll(ctx)
	if err != nil {
		fmt.Println(1)
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timelinesPublic); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
