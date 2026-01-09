package router

// defining proxies
import (
	"gateway/internal/proxy"
	"net/http"
)

type Proxies struct {
	Auth     http.Handler
	Upload   http.Handler
	Metadata http.Handler
	Sharing  http.Handler
	Preview  http.Handler
}

func NewProxies() *Proxies {
	return &Proxies{
		Auth:     proxy.NewReverseProxy("http://localhost:9001"),
		Upload:   proxy.NewReverseProxy("http://localhost:9002"),
		Metadata: proxy.NewReverseProxy("http://localhost:9003"),
		Sharing:  proxy.NewReverseProxy("http://localhost:9004"),
		Preview:  proxy.NewReverseProxy("http://localhost:9005"),
	}
}
