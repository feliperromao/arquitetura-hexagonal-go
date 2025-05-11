package server

import (
	"arquitetura-hexagonal-go/adapters/web/handler"
	"arquitetura-hexagonal-go/application"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeWebServer() *WebServer{
	return &WebServer{}
}

func (w WebServer) Serve() {
	router := mux.NewRouter()
	negr := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(*router, *negr, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		Addr: ":9000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
