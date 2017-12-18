package middlewares

import (
	"net/http"
)

type Auth struct {
	authType   string
	authParams map[string]string
}

func CreateNewAuth(authType string, authParams map[string]string) *Auth {
	return &Auth{authType, authParams}
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	isValid := false

	switch a.authType {
	case "get-token":
		isValid = a.authParams["token"] == r.URL.Query().Get("token")
	case "header-token":
		isValid = a.authParams["token"] == r.Header.Get("Auth-Token")
	}

	if isValid {
		return true
	}

	w.WriteHeader(403)

	return false
}
