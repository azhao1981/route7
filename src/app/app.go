// About route7
package app

import (
	"config"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"route"
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

	server := Server{config: c}
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
	s.beforeSend()
	defer s.afterSend()

	route := s.findRoute(req)
	s.transit(route, req)

}

// Find match route for request
func (s *Server) findRoute(req *http.Request) (route *route.Route) {
	route = s.config.FindRoute(req)
	fmt.Println("find route id : ", route.Id)
	return
}

// Transit request to target url
func (s *Server) transit(route *route.Route, req *http.Request) (res *http.ResponseWriter) {
	return
}

const (
	CLR_N = "\x1b[0m"
	/* you use codes 30+i to specify foreground color, 40+i to specify background color */
	BLACK   = 0
	RED     = 1
	GREEN   = 2
	YELLO   = 3
	BLUE    = 4
	MAGENTA = 5
	CYAN    = 6
	WHITE   = 7
)

// 返回ANSI 控制台颜色格式的字符串
//bc 背景颜色
//fc 前景(文字)颜色
func ansi_color(bc int, fc int, s string) string {
	return fmt.Sprintf("\x1b[%d;%dm%s%s", 40+bc, 30+fc, s, CLR_N)
}

// @beforeSend
func (s *Server) beforeSend() {
	fmt.Println(ansi_color(BLACK, RED, "req:1111"))
}

// @afterSend
func (s *Server) afterSend() {
	fmt.Println("req:222")
}
