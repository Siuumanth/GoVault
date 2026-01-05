package router

// defining proxies
import (
	"gateway/internal/proxy"
	"net/http"
)

type Proxies struct {
	AuthProxy     http.Handler
	UploadProxy   http.Handler
	MetadataProxy http.Handler
	SharingProxy  http.Handler
	PreviewProxy  http.Handler
}

func NewProxies() *Proxies {
	return &Proxies{
		AuthProxy:     proxy.NewReverseProxy("http://localhost:9001"),
		UploadProxy:   proxy.NewReverseProxy("http://localhost:9002"),
		MetadataProxy: proxy.NewReverseProxy("http://localhost:9003"),
		SharingProxy:  proxy.NewReverseProxy("http://localhost:9004"),
		PreviewProxy:  proxy.NewReverseProxy("http://localhost:9005"),
	}
}
