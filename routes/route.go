// route/router.go
package route

import (
    "net/http"

    "assessment_cld/controller"
)

// SetupRouter sets up the HTTP routes
func SetupRouter() {
    http.HandleFunc("/submit", controller.HandleRequest)
}
