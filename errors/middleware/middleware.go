package middleware

import (
	"net/http"

	"github.com/Denvos/core/errors"
	"github.com/Denvos/core/errors/codes"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(*errors.Error)
				if !ok {
					err = errors.New(codes.Internal, "panic recovered")
				}
				w.WriteHeader(codes.DefaultHTTPStatus[err.Code()])
				resp := ErrorResponse{
					Code:    err.Code(),
					Message: err.Message(),
					Details: err.Fields(),
				}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
