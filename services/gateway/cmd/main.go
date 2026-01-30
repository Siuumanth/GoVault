package main

import (
	"context"
	"fmt"
	"gateway/internal/gateway"
	MW "gateway/internal/middleware"
	"gateway/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

/*
Steps to create the api gatway:
1. Get gateway Deps
2. Create gateway
3. Get proxies
3. Get router
4. Inject router while buidling gateway
*/

func main() {
	godotenv.Load()

	rl := MW.NewBasicRateLimiter(1000, time.Minute)
	authz := MW.NewAuthZ()
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
	r := router.NewChiRouter()
	r.ConfigureRoutes(proxies, authz)

	finalGateway := gw.BuildGateway(r)

	// finalGateway is a handler with all the middlewares applied
	server := &http.Server{
		Addr:    ":9000",
		Handler: finalGateway,
	}

	// 1. create a channel to listen for OS signals (Interrupt/Kill), eg:- Ctrl + C is a kill
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM) // set stop channel as a listener

	// 2. run the server in a goroutine so it doesn't block
	go func() {
		fmt.Println("Gateway starting on :9000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	// 3. wait for any signal by OS (usere)
	<-stop

	// 4. shut down gracefully with some timeout
	fmt.Println("\nShutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
