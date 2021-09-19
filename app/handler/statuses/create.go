package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := auth.AccountOf(r)
	accountID := int64(account.ID)

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	daoAccount := h.app.Dao.Status() // domain/repository の取得
	status, err := daoAccount.CreateStatus(ctx, req.Status, accountID)

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
