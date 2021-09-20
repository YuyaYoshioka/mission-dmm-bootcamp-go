package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := request.IDOf(r)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	
	daoStatus := h.app.Dao.Status()
	err = daoStatus.DeleteByID(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	
	w.Header().Set("Content-Type", "applicatin/json")
	if err := json.NewEncoder(w).Encode(id); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
