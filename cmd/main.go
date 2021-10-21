package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Cyb3r-Jak3/common/v2"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
	db  *sql.DB
)

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
	if err := http.ListenAndServe("127.0.0.1:5000", r); err != nil {
		log.WithError(err).Fatal("Error when starting http server")
	}
}
