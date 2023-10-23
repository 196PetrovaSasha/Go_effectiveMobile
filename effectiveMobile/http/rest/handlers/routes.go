package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// adding logger middleware
	r.Use(handler.MiddlewareLogger())
	//r.HandleFunc("/healthz", handler.Health())
	r.HandleFunc("/fio", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/fio/{id}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/fio/{id}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/fio/{id}", handler.Delete()).Methods(http.MethodDelete)
}
