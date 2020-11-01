package decorator

import "net/http"

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuth(r) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("forbidden"))
			return
		}
		h.ServeHTTP(w, r)
	})
}

func isAuth(_ *http.Request) bool {
	return true
}
