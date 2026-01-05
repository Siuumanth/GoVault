package router

/*
Middlewares order:
Logging
SecurityHeaders
CORS
JWT (only for protected routes)
RateLimit
HeaderInjection (only before proxy)
↓
Handler / Proxy

*/

var ReverseProxies *Proxies = NewProxies()

// we will implement our MW chain using an explicit pipeline
// we could do by wrapping MW functions, but that would be harder to understand

type Gateway struct {
	jwt JWTMW
}
