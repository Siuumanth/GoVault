package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) http.Handler {
	u, err := url.Parse(target)
	if err != nil {
		panic(err) // ok at startup
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request already has headers added by gateway middleware
		proxy.ServeHTTP(w, r)
	})
}
