package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"git.icysoft.fr/cedric/industry-go-server/api"

	aws "github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/gin-gonic/gin"
)

// Handler is the exported handler called by AWS Lambda.
var Handler apigatewayproxy.Handler

func init() {
	Handler = NewHandler()
}

func NewHandler() apigatewayproxy.Handler {
	ln := aws.Listen()

	// In local development environment, open 8080 port
	if os.Getenv("GIN_MODE") == "" {
		ln, _ = net.Listen("tcp", "0.0.0.0:8080")
	}

	// Amazon API Gateway Binary support out of the box.
	handle := apigatewayproxy.New(ln, nil).Handle

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, Goa, etc.
	go http.Serve(ln, setUpGin())

	log.Println(ln.Addr().String())

	return handle
}

func setUpGin() *gin.Engine {
	r := gin.Default()

	// Player
	player := r.Group("/player")
	{
		player.POST("/", api.NewPlayer)
		player.PUT("/:id", api.UpdatePlayer)
		player.GET("/:id", api.GetPlayer)
	}

	return r
}

func main() {
	for {
		time.Sleep(time.Hour)
	}
}
