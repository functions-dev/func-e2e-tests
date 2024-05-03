package function

import (
	"fmt"
	"net/http"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "func-e2e-test-deploy-source")
}
