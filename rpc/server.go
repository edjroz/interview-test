// This package is for the RPC/REST API
package rpc

import (
	"log"
	"net/http"
)

// "startRPC" starts the client RPC/REST API server at a specific port.
func StartRPC(port string) {
	log.Fatal(http.ListenAndServe(":"+port, Router(Routes())))
}
