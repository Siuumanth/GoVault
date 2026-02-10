package router

// defining proxies
import (
	"gateway/internal/proxy"
	"net/http"
	"os"
)

type Proxies struct {
	Auth   http.Handler
	Upload http.Handler
	Files  http.Handler
}

func NewProxies() *Proxies {
	return &Proxies{
		Auth:   proxy.NewReverseProxy(os.Getenv("GOVAULT_AUTH_SERVICE_URL")),
		Upload: proxy.NewReverseProxy(os.Getenv("GOVAULT_UPLOAD_SERVICE_URL")),
		Files:  proxy.NewReverseProxy(os.Getenv("GOVAULT_FILES_SERVICE_URL")),
	}
}
