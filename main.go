package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/kpango/glg"

	"./server"
)

var (
	host         string
	port         int
	readTimeout  time.Duration
	writeTimeout time.Duration
	useTLS       bool
	tlsCert      string
	tlsKey       string
	bindPath     string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Host IP address or DNS name")
	flag.IntVar(&port, "port", 8080, "Http server port")
	flag.DurationVar(&readTimeout, "read-timeout", 0, "Timeout for reading data")
	flag.DurationVar(&writeTimeout, "write-timeout", 0, "Timeout for writing data")
	flag.BoolVar(&useTLS, "use-tls", false, "Enables HTTPS by TLS certificate")
	flag.StringVar(&tlsCert, "tls-cert", "cert.pem", "Path to TLS certificate")
	flag.StringVar(&tlsKey, "tls-key", "key.pem", "Path to TLS key file")
	flag.StringVar(&bindPath, "path", "/", "Path for binding service")
}

func main() {
	flag.Parse()

	glg.Infof("Starting server at %v:%v", host, port)

	app := server.App{
		BindPath: bindPath,
	}
	app.Initialize()
	s := http.Server{
		Addr:         fmt.Sprintf("%v:%v", host, port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	if useTLS {
		app.Run(s)
	} else {
		app.RunTLS(s, tlsCert, tlsKey)
	}
}
