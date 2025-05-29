package function

import (
	"fmt"
	"net/http"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	// The E2E test checks for this exact string in the HTTP response to
	// confirm the function was both initialized using this template and
	// successfully deployed.
	fmt.Fprintln(w, "func-e2e-test-core-deploy-template")
}
