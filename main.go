// main.go

package main

import (
    "log"
    "net/http"

    "assessment_cld/routes" // Import the routes package
)

func main() {
    // Setup HTTP routes
    route.SetupRouter() // Call the SetupRouter function from the routes package

    port := ":8080"
    log.Printf("Server is running at http://localhost%s\n", port) // Log a message indicating the server is running
    if err := http.ListenAndServe(port, nil); err != nil { // Start the HTTP server
        log.Fatal("Server error:", err) // Log any server errors and exit
    }
}
