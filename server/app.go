package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kpango/glg"

	"github.com/google/uuid"

	"github.com/gorilla/mux"

	"../stream"
)

// Config структура конфигурации сервера
type Config struct {
	Host         string
	Port         int
	BindPath     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	UseTLS       bool
	CertPath     string
	KeyPath      string
}

// App приложение сервера трансляций
type App struct {
	config  Config
	server  *http.Server
	router  *mux.Router
	streams stream.Repository
}

// Initialize инициализирует роутер, хранилище данных и сохраняет конфигурацию
func (app *App) Initialize(conf Config) {
	app.config = conf
	app.streams = stream.NewMapRepository()

	app.router = mux.NewRouter()
	app.initializeRoutes()
}

// Run запускает сервер синхронно
func (app *App) Run() {
	app.server = &http.Server{
		Addr:         app.config.Host + ":" + fmt.Sprint(app.config.Port),
		Handler:      app.router,
		ReadTimeout:  app.config.ReadTimeout,
		WriteTimeout: app.config.WriteTimeout,
	}

	if app.config.UseTLS {
		glg.Info("Run server at https://%v", app.server.Addr)
		glg.Fatalln(app.server.ListenAndServeTLS(app.config.CertPath, app.config.KeyPath))
	} else {
		glg.Infof("Run server at http://%v", app.server.Addr)
		glg.Fatalln(app.server.ListenAndServe())
	}
}

// RunAsync запускает сервис асинхронно
func (app *App) RunAsync() {
	go app.Run()
}

// Shutdown пытается корректно завершить работу сервера
func (app *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.server.Shutdown(ctx)
}

func (app *App) initializeRoutes() {
	app.router.HandleFunc(app.config.BindPath, app.getStreamInfo).Methods(http.MethodGet)
	app.router.HandleFunc(app.config.BindPath, app.createStream).Methods(http.MethodPost)
	app.router.HandleFunc(app.config.BindPath, app.changeStreamState).Methods(http.MethodPut)
	app.router.HandleFunc(app.config.BindPath, app.deleteStream).Methods(http.MethodDelete)
}

func (app *App) getStreamInfo(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	id, err := uuid.Parse(args.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorData(err))
		return
	}

	stream, err := app.streams.GetInfo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorData(err))
		return
	}

	res, err := json.Marshal(StreamInfo{
		Type:       "stream",
		ID:         id.String(),
		Attributes: stream,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorData(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *App) createStream(w http.ResponseWriter, r *http.Request) {
	id, err := app.streams.CreateStream()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatErrorData(err))
		return
	}

	stream, err := app.streams.GetInfo(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatErrorData(err))
		return
	}

	res, err := json.Marshal(StreamInfo{
		Type:       "data",
		ID:         id.String(),
		Attributes: stream,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorData(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *App) changeStreamState(w http.ResponseWriter, r *http.Request) {

}

func (app *App) deleteStream(w http.ResponseWriter, r *http.Request) {

}

func formatErrorData(e error) []byte {
	res, err := json.Marshal(ErrorInfo{
		Type:  "error",
		Error: e.Error(),
	})
	if err != nil {
		return []byte("Error with marshalling error")
	}

	return res
}
