package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/Ali-Assar/car-rental-system/aggregator/client"
	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("HTTP listenAddr", ":6000", "the listen address of HTTP server")
	aggregatorServiceAddr := flag.String("aggServiceAddr", "http://localhost:4000", "the listen address of aggregator service")
	flag.Parse()
	var (
		client     = client.NewHTTPClient(*aggregatorServiceAddr)
		invHandler = NewInvocieHandler(client)
	)
	http.HandleFunc("/invoice", makeApiFunc(invHandler.handleGetInvoice))
	logrus.Infof("gateway http server is running on port%s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

}

type InvoiceHandler struct {
	client client.Client
}

func NewInvocieHandler(c client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: c,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {

	invoice, err := h.client.GetInvoice(context.Background(), 1659034394)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, invoice)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func makeApiFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri":  r.RequestURI,
			}).Info("req:: ")
		}(time.Now())

		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
