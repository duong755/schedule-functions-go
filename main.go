// Vercel Serverless does not use this file

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "schedule.functions/api/v1"
	v2 "schedule.functions/api/v2"
)

func main() {
	rootRouter := mux.NewRouter()

	apiV1Router := rootRouter.PathPrefix("/api/v1").Subrouter()
	apiV1Router.HandleFunc("/classmembers", v1.ClassMembersHandler).Methods(http.MethodGet)
	apiV1Router.HandleFunc("/schedules", v1.SchedulesHandler).Methods(http.MethodGet)

	apiV2Router := rootRouter.PathPrefix("/api/v2").Subrouter()
	apiV2Router.HandleFunc("/classmembers", v2.ClassMembersHandler).Methods(http.MethodGet)
	apiV2Router.HandleFunc("/schedules", v2.SchedulesHandler).Methods(http.MethodGet)

	http.ListenAndServe(":5000", rootRouter)
}
