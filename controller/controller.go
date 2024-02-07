package controller

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "assessment_cld/model" // Import the model package
    "assessment_cld/worker"
)

// HandleRequest handles incoming HTTP requests
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    // Decode the JSON request body into the ContactFormSubmission struct
    var requestBody model.ContactFormSubmission
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Process the request as needed
    fmt.Println("Received request:", requestBody)

    // Create a channel to pass the requestBody to the worker
    channel := make(chan model.ContactFormSubmission)

    // Start the worker to process messages from the channel
    go worker.ProcessMessages(channel)

    // Send the requestBody to the worker through the channel
    channel <- requestBody

    // Include the input data in the response
    responseData := map[string]interface{}{
        "message": "Received your request",
        "data":    requestBody,
    }

    // Send a response with the input data included
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(responseData); err != nil {
        log.Println("Error encoding response:", err)
    }
}
