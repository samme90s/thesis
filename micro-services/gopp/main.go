package main
import (
    "fmt"
    "io"
    "log"
    "net/http"
    "gopp/datacleaner"
)
func dataCleanerHandler(w http.ResponseWriter, r *http.Request) {
    // Read request body
    rawBody, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "unable to read request", http.StatusBadRequest)
        return
    }
    // Clean the JSON data using the datacleaner module.
    cleanedJSON, err := datacleaner.CleanJSON(rawBody)
    if err != nil {
        http.Error(w, "error cleaning JSON data", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Cleaned JSON: %s", cleanedJSON)
}
func main() {
    http.HandleFunc("/clean", dataCleanerHandler)
    log.Println("Data Cleaner service running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}