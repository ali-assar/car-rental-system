package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	listenAddr := flag.String("http address", ":6000", "the listen address of HTTP server")
	flag.Parse()

	http.HandleFunc("/invoice", handleGetInvoice)
	logrus.Info("gateway http server is running on port 6000")
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

}

func handleGetInvoice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("everything is fine"))
}
