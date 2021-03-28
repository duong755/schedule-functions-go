package main

import (
	"net/http"

	apiv1 "schedule.functions/api/v1"
)

func main() {
	http.HandleFunc("/api/v1/schedules", apiv1.SchedulesHandler)
	http.HandleFunc("/api/v1/classmembers", apiv1.ClassMembersHandler)
}
