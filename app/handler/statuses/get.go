package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

// Handle request for `GET /v1/statuses/:id`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := request.IDOf(r)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	daoStatus := h.app.Dao.Status() // domain/repository の取得
	status, err := daoStatus.FindByID(ctx, id)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	daoAccount := h.app.Dao.Account()
	account, err := daoAccount.FindByID(ctx, status.AccountID)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	res := map[string]interface{}{
		"id": status.ID,
		"account": account,
		"content": status.Content,
		"create_at": status.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
