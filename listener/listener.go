package listener

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//PortListener - Responsible for exposing an endpoint listener and provide some stock responses
type PortListener interface {
	ListenAndProvideStockResponses()
}

//DefaultPortListener - default implementation of PortListener
type DefaultPortListener struct {
	IP     string
	Port   int
	Server *http.Server
}

func NewPortListener(IP string, Port int) (PortListener, error) {
	return &DefaultPortListener{IP: IP, Port: Port}, nil
}

func (portListener *DefaultPortListener) ListenAndProvideStockResponses() {
	r := mux.NewRouter()
	r.HandleFunc("/", portListener.basicHandler)

	portListener.Server = &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", portListener.IP, portListener.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go portListener.Server.ListenAndServe()
}

func (portListener *DefaultPortListener) basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte(fmt.Sprintf("Hello from : %s:%d", portListener.IP, portListener.Port)))
}
