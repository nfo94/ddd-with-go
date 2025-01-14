package ch2

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	IsUserPaymentActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	m := mux.NewRouter()
	m.HandleFunc("/user/{userID}/payment/active", func(w http.ResponseWriter, r *http.Request) {
		uID := mux.Vars(r)["userID"]
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		isActive := u.IsUserPaymentActive(r.Context(), uID)

		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(b)
	}).Methods(http.MethodGet)
}
