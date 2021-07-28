package middlewares

import (
	"github.com/noguchidaisuke/go-mysql-docker/api/utils"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Debug(struct {
			Method string `json:"method"`
			Url    string `json:"url"`
			Proto  string `json:"proto"`
		}{
			r.Method,
			r.Host + r.RequestURI,
			r.Proto,
		})

		next(w, r)
	}
}
