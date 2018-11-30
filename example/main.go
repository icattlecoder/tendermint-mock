package main

import (
	"github.com/icattlecoder/tendermint-mock"
	"net/http"
)

func main() {
	nw := tendermint_mock.NewNetwork(4)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		nw.AddNewProcess()
	})
	http.ListenAndServe(":3303", mux)
}
