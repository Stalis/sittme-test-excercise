package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"./server"
)

var (
	host          string
	port          int
	readTimeout   time.Duration
	writeTimeout  time.Duration
	useTLS        bool
	tlsCert       string
	tlsKey        string
	bindPath      string
	streamTimeout time.Duration
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
	flag.DurationVar(&streamTimeout, "stream-timeout", 5*time.Second, "Timeout for finish stream after interrupted")
}

func main() {
	flag.Parse()

	app := server.App{}
	app.Initialize(server.Config{
		Host:          host,
		Port:          port,
		BindPath:      bindPath,
		ReadTimeout:   readTimeout,
		WriteTimeout:  writeTimeout,
		UseTLS:        useTLS,
		CertPath:      tlsCert,
		KeyPath:       tlsKey,
		StreamTimeout: streamTimeout,
	})

	app.RunAsync()

	// ожидание команды прерывания
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	app.Shutdown()
}
