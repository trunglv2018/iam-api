package main

import (
	// 1. load first

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gitlab.com/trunglen/iam-api/config"
	"gitlab.com/trunglen/iam-api/helpers"
	"gitlab.com/trunglen/iam-api/middlewares"
)

func main() {
	e, err := casbin.NewEnforcer("acl-model/model.conf", "acl-model/policy.csv")
	e.AddFunction("keyContain", helpers.KeyContainFunc)
	e.AddFunction("customRegexMatch", helpers.CustomRegexMatchFunc)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtMiddleware(e), middlewares.SkywalkerMiddleware(router, config.SkyTracer))
	router.Any("/", handleAuthorize)
	// v1 := router.Group("api/v1")
	// {
	// 	admin.NewAdminApi(v1.Group("admin"))
	// 	customer.NewCustomerApi(v1.Group("customer"))
	// }
	srv := &http.Server{
		Addr:    ":" + config.GetConfig().GetString("server.port"),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

func handleAuthorize(c *gin.Context) {
	c.String(200, "Forwarding")
}
