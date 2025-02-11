package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jetnoli/go-router/utils"
)

func CheckAuthorization(w *http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId")

	if userId == "" || userId == nil {
		http.Error(*w, "no access token provided", http.StatusUnauthorized)
		utils.CancelRequest(r)
	}
}

func DecodeToken(w *http.ResponseWriter, r *http.Request) {
	token := ""

	for _, cookie := range r.Cookies() {
		if cookie.Name == "Authorization" {
			token = cookie.Value
		}
	}

	if token == "" {
		fmt.Println("No Cookie")
		return
	}

	userId, err := utils.DecodeJwt(token)

	if err != nil {
		fmt.Println("error decoding token: " + err.Error())
		return
	}

	*r = *r.Clone(context.WithValue(r.Context(), "userId", userId))
}
