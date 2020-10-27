package server

import (
	"encoding/json"
	"net/http"

	"github.com/kpango/glg"

	"github.com/google/uuid"

	"github.com/gorilla/mux"

	"../stream"
)

type App struct {
	BindPath string
	router   *mux.Router
	server   http.Server
	streams  stream.Repository
}

func (app *App) Initialize() {
	app.streams = stream.NewMapRepository()

	app.router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(server http.Server) {
	server.Handler = app.router
	glg.Fatalln(server.ListenAndServe())
}

func (app *App) RunTLS(server http.Server, certPath, keyPath string) {
	server.Handler = app.router
	glg.Fatalln(server.ListenAndServeTLS(certPath, keyPath))
}

func (app *App) initializeRoutes() {
	app.router.HandleFunc(app.BindPath, app.getStreamInfo).Methods(http.MethodGet)
	app.router.HandleFunc(app.BindPath, app.createStream).Methods(http.MethodPost)
	app.router.HandleFunc(app.BindPath, app.changeStreamState).Methods(http.MethodPut)
	app.router.HandleFunc(app.BindPath, app.deleteStream).Methods(http.MethodDelete)
}

func (app *App) getStreamInfo(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	id, err := uuid.Parse(args.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stream, err := app.streams.GetInfo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(StreamInfo{
		Type:       "data",
		ID:         id.String(),
		Attributes: stream,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *App) createStream(w http.ResponseWriter, r *http.Request) {
	id, err := app.streams.CreateStream()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	stream, err := app.streams.GetInfo(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := json.Marshal(StreamInfo{
		Type:       "data",
		ID:         id.String(),
		Attributes: stream,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *App) changeStreamState(w http.ResponseWriter, r *http.Request) {

}

func (app *App) deleteStream(w http.ResponseWriter, r *http.Request) {

}
