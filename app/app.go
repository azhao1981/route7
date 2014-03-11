// About route7
package app

import (
	"fmt"
	"github.com/azhao1981/route7/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	APP_NAME = "route7"
	VERSION  = "0.9.0"
)

type Server struct {
	config *config.Config
}

func Run(c *config.Config) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	server := Server{}
	go func() {
		s := &http.Server{
			Addr:           fmt.Sprintf("%s:%d", c.Listen.Host, c.Listen.Port),
			Handler:        http.HandlerFunc(server.handler_func),
			ReadTimeout:    120 * time.Second,
			WriteTimeout:   120 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())
	}()

	<-sigchan
}

// @handler_func
// handler function for a request:
// 1. Find a route match it and transit to target server
// 2. Found not route match , send it to default targe server
// 3. Keep Get/Post or Order method
func (s *Server) handler_func(res http.ResponseWriter, req *http.Request) {
	s.preSend()
	defer s.afterSend()
}

// @preSend
func (s *Server) preSend() {
	fmt.Println("req:")
}

// @afterSend
func (s *Server) afterSend() {
	fmt.Println("req:")
}
