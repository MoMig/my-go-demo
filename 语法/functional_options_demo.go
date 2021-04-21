package main

import (
	"crypto/tls"
	"fmt"
	"time"
)
/**
函数式编程
 */

type server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

//定义函数类型
type Option func(server2 *server)

func Protocol(p string) Option {
	return func(s *server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxcount int) Option {
	return func(s *server){
		s.MaxConns = maxcount
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(server2 *server)) (*server,error) {
	svr := server{
		Addr: addr,
		Port: port,
	}
	for _, option := range options {
		option(&svr)
	}

	return &svr, nil
}

func main() {
	s, _ := NewServer("localhost", 8081)
	fmt.Printf("ip:%s port:%d\n", s.Addr, s.Port)

	s2, _ := NewServer("localhost", 8081,Timeout(10))
	fmt.Printf("ip:%s port:%d timeout:%s\n", s2.Addr, s2.Port, s2.Timeout)

	s3, _ := NewServer("localhost", 8081,Timeout(10),Protocol("http"))
	fmt.Printf("ip:%s port:%d timeout:%s protocol:%s\n", s3.Addr, s3.Port, s3.Timeout, s3.Protocol)
}