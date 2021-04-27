package main

import (
	"net/http"

	v1 "schedule.functions/api/v1"
	v2 "schedule.functions/api/v2"
)

func main() {
	http.HandleFunc("/api/v1/classmembers", v1.ClassMembersHandler)
	http.HandleFunc("/api/v1/schedules", v1.SchedulesHandler)

	http.HandleFunc("/api/v2/classmembers", v2.ClassMembersHandler)
	http.HandleFunc("/api/v2/schedules", v1.SchedulesHandler)

	http.ListenAndServe(":5000", nil)
}
