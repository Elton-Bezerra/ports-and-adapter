package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Elton-Bezerra/ports-and-adapter/adapters/web/handler"
	"github.com/Elton-Bezerra/ports-and-adapter/app"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service app.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
