package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type request struct {
	Formula string `json:"formula"`
}

type response struct {
	Result float64 `json:"result"`
}

func bodmasHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req request
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		log.Printf("error unmarshaling request: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// tokenise
	// execute
	res := response{
		Result: 69,
	}
	log.Printf("%s = %v", req.Formula, res.Result)
	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	n, err := io.Copy(w, bytes.NewBuffer(resBytes))
	if err != nil {
		log.Printf("error writing response: %d/%d bytes: %+v", n, len(resBytes), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
