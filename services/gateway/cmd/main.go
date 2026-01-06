package main

import (
	"gateway/internal/gateway"
	MW "gateway/internal/middleware"
	"gateway/internal/router"
	"net/http"
	"time"
)

/*
Steps to create the api gatway:
1. Get gateway Deps
2. Create gateway
3. Get proxies
3. Get router
4. Inject router while buidling gatewat
*/

func main() {
	// all controllable from the main function
	gatewayDeps := &gateway.GatewayDeps{
		JWT:              MW.NewJWT(),
		CORS:             MW.NewCORS(),
		SecurityHeaders:  MW.NewSecurityHeaders(),
		Logger:           MW.NewLogger(),
		RateLimiter:      MW.NewMemoryRateLimiter(1000, time.Minute),
		HeadersInjection: MW.NewHeadersInjection(),
	}
	gateway := gateway.NewGateway(gatewayDeps) // dereferencing gatewayDeps
	proxies := router.NewProxies()
	router := router.NewChiRouter(proxies)

	finalGateway := gateway.BuildGateway(router)

	// finalGateway is a handler with all the middlewares applied
	err := http.ListenAndServe(":8080", finalGateway)
	if err != nil {
		panic(err)
	}
}
