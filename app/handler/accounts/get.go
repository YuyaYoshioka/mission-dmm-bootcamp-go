package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/accounts/:username`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := strings.TrimPrefix(r.URL.Path, "/v1/accounts")
	username = strings.Replace(username, "/", "", 1)

	daoAccount := h.app.Dao.Account() // domain/repository の取得
	account, err := daoAccount.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if account == nil {
		err := errors.New("user not found")
    println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
