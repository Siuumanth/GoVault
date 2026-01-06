package main

import (
	"gateway/internal/gateway"
	MW "gateway/internal/middleware"
	"gateway/internal/router"
	"net/http"
	"time"

	"github.com/joho/godotenv"
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
	godotenv.Load()

	rl := MW.NewBasicRateLimiter(1000, time.Minute)
	// all controllable from the main function - DI
	gatewayDeps := &gateway.GatewayDeps{
		JWT:              MW.NewJWT(),
		CORS:             MW.NewCORS(),
		SecurityHeaders:  MW.NewSecurityHeaders(),
		Logger:           MW.NewLogger(),
		RateLimiter:      MW.NewRateLimiter(rl),
		HeadersInjection: MW.NewHeadersInjection(),
	}
	gw := gateway.NewGateway(gatewayDeps)
	proxies := router.NewProxies()
	r := router.NewChiRouter(proxies)

	finalGateway := gw.BuildGateway(r)

	// finalGateway is a handler with all the middlewares applied
	err := http.ListenAndServe(":8080", finalGateway)
	if err != nil {
		panic(err)
	}
}
