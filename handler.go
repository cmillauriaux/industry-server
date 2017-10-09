package main

import (
	"log"
	"net/http"
	"time"

	"git.icysoft.fr/cedric/industry-go-server/api"

	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/gorilla/mux"
)

// Handler is the exported handler called by AWS Lambda.
var Handler apigatewayproxy.Handler

func init() {
	Handler = NewHandler()
}

func NewHandler() apigatewayproxy.Handler {
	ln := net.Listen()

	// Amazon API Gateway Binary support out of the box.
	handle := apigatewayproxy.New(ln, nil).Handle

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, Goa, etc.
	go http.Serve(ln, setUpMux())

	log.Println(ln.Addr().String())

	return handle
}

func setUpMux() *mux.Router {
	r := mux.NewRouter()

	// Player
	r.HandleFunc("/player", api.NewPlayer).Methods(http.MethodPost)
	r.HandleFunc("/player/{id}", api.GetPlayer).Methods(http.MethodGet)

	return r
}

func main() {
	for {
		time.Sleep(time.Second)
	}
}
