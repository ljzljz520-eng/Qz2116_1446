package main

import (
	"log"
	"net/http"
	"os"
	"streetlight40/dp"
	"streetlight40/httpapi"
	"streetlight40/model"
	"streetlight40/query"
	"streetlight40/register"
	"streetlight40/storage"
	"streetlight40/workflow"
)

func main() {
	path := os.Getenv("STREETLIGHT_DB")
	if path == "" {
		path = "streetlight.db"
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	u := model.User{ID: 1, Name: "operator", Role: "admin", Active: true}
	_ = s.SaveUser(u)
	r := register.New(s, dp.NewEngine(1000, 40))
	q := query.New(s)
	w := workflow.New(r, q)
	srv := httpapi.New(w)
	log.Println("streetlight40 listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", srv.Handler()))
}
