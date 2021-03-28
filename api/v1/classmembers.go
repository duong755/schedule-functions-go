package v1

import (
	"fmt"
	"net/http"
)

// ClassMembersHandler returns the function name
func ClassMembersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "ClassMembersHandler")
}
