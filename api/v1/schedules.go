package v1

import (
	"fmt"
	"net/http"
)

// SchedulesHandler returns the function name
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "SchedulesHandler")
}
