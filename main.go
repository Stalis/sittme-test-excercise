package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/kpango/glg"

	"github.com/gorilla/mux"

	"./api"
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

	server := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", host, port),
		Handler:      NewRouter(),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	var err error
	if useTLS {
		err = server.ListenAndServeTLS(tlsCert, tlsKey)
	} else {
		err = server.ListenAndServe()
	}

	if err != nil {
		glg.Fatal(err)
	}
}

// NewRouter Создает экземпляр роутера с обработчиком запросов для сервера
func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(bindPath, Handler)
	return router
}

// Handler функция обработки запроса
func Handler(w http.ResponseWriter, r *http.Request) {
	api.ProcessRequest(r.Method, r.URL.Query())
	err := r.Write(w)
	if err != nil {
		glg.Error(err)
	}
}
