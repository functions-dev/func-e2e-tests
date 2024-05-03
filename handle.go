package function

import (
	"fmt"
	"net/http"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	// This exact string is matched by the E2E test to confirm the funciton
	// was loaded from this repository and successfully deployed.
	fmt.Fprintln(w, "func-e2e-test-deploy-source")
}
