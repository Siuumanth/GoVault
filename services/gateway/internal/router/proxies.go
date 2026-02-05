package router

// defining proxies
import (
	"gateway/internal/proxy"
	"net/http"
)

type Proxies struct {
	Auth   http.Handler
	Upload http.Handler
	Files  http.Handler
}

func NewProxies() *Proxies {
	return &Proxies{
		Auth:   proxy.NewReverseProxy("http://localhost:9001"),
		Upload: proxy.NewReverseProxy("http://localhost:9002"),
		Files:  proxy.NewReverseProxy("http://localhost:9003"),
	}
}
