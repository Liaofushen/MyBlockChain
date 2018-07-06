package main

import (
	"net/http"
	"../core"
	"encoding/json"
	"io"
	"fmt"
)

var blockChain *core.BlockChain


func main() {
	blockChain = core.NewBlockChain()
	http.HandleFunc(
		"/get",
		func(w http.ResponseWriter, r *http.Request) {
			bytes, err := json.Marshal(blockChain)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			blockChain.Read()
			io.WriteString(w, string(bytes))})
	http.HandleFunc(
		"/write",
		func(w http.ResponseWriter, r *http.Request) {
			blockData := r.URL.Query().Get("data")
			blockChain.Write(blockData)

			fmt.Println(blockData)
			bytes, err := json.Marshal(blockChain)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			io.WriteString(w, string(bytes))
		})
	http.ListenAndServe("localhost:8080", nil)
}
