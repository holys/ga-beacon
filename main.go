package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/holys/ga-beacon/ga-beacon"
	"github.com/ngaut/log"

	"flag"
)

var (
	port = flag.Int64("p", 9091, "server port for listening")
)

func main() {
	flag.Parse()

	go func() {
		addr := fmt.Sprintf("0.0.0.0:%d", *port)
		log.Infof("server listening on %s", addr)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err.Error())
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sc
	log.Infof("Got signal [%d] to exit.", sig)
	os.Exit(0)
}
