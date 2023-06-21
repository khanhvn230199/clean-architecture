package middleware

import (
	"net/http"

	httppkg "github.com/example-golang-projects/clean-architecture/packages/http"
)

func APILoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := httppkg.NewLoggingResponseWriter(w)
		//span := zipkin.SpanFromContext(r.Context())
		defer func() {
			//log.Info(fmt.Sprintf("API infomation: %v [%v]", r.RequestURI, r.Method), span, map[string]interface{}{
			//	"method": r.Method,
			//	"path":   r.RequestURI,
			//	"status": lrw.StatusCode,
			//})
		}()
		next.ServeHTTP(lrw, r)
	})
}
