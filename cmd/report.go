package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Cyb3r-Jak3/common/v2"
)

// Report404 represents a 404 report from Cloudflare Workers
type Report404 struct {
	URL    string `json:"url,omitempty"`
	IP     string `json:"ip,omitempty"`
	Method string `json:"method,omitempty"`
	Time int `json:"time,omitempty"`
	Host string `json:"host,omitempty"`
}

func report404(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1024)
	if req.Body == http.NoBody || req.ContentLength == 0 {
		log.Info("JSON body was not sent")
		http.Error(w, "JSON body required", http.StatusBadRequest)
		return
	}
	out, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithError(err).Error("Error reading request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data Report404
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.WithError(err).Error("Error unmarshaling JSON")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = insert("INSERT INTO report404(url, ip, method, time, host) VALUES ($1, $2, $3, $4, $5);", data.URL, data.IP, data.Method, data.Time, data.Host)
	if err != nil {
		log.WithError(err).Error("Error inserting into database")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	common.StringResponse(w, "Received")
}
