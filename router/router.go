package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pallat/micro/order"
)

type Router struct {
	*gin.Engine
}

func New() *Router {
	return &Router{gin.Default()}
}

type HandlerFunc func(order.Context)

func (r *Router) GET(relativePath string, handler HandlerFunc) {
	r.Engine.GET(relativePath, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func (r *Router) POST(relativePath string, handler HandlerFunc) {
	r.Engine.POST(relativePath, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func (r *Router) ListenAndServe() func() {
	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		<-ctx.Done()
		stop()
		fmt.Println("shutting down gracefully, press Ctrl+C again to force")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(timeoutCtx); err != nil {
			fmt.Println(err)
		}
	}
}

type Context struct {
	*gin.Context
}

func (c *Context) Order() (o order.Order, err error) {
	err = c.ShouldBindJSON(&o)
	return
}

func (c *Context) JSON(code int, v interface{}) {
	c.Context.JSON(code, v)
}

func (c *Context) Status(code int) {
	c.Context.Status(code)
}
