package main

import (
	"fmt"
	_const "github.com/andrewnguyen22/pocket-interview-test/const"
	"github.com/andrewnguyen22/pocket-interview-test/rpc"
	"os"
	"os/signal"
)

// "init" is a built in function that is automatically called before main.
func init() {

}

// "main" is the starting function of the client.
func main() {
	go rpc.StartRPC(_const.RPCPort)
	// Catches OS system interrupt signal and calls unregister
	fmt.Println("Started Pockethereum")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	select {
	case sig := <-c:
		fmt.Println(sig.String() + " signal received")
		os.Exit(0)
	}
}
