package main

import (
	"net/http"
	"os"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	DbURL = os.Getenv("DATABASE_URL")
	if DbURL == "" {
		log.Fatal("There is not database URL configured")
	}
	createTables()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/report/404", common.AllowedMethod(report404, "POST"))
	log.Info("Starting Server")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		log.WithError(err).Fatal("Error when starting http server")
	}
}
