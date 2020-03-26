package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    // Create HTTP client with timeout
    client := &http.Client{
        Timeout: 30 * time.Second,
    }

    // Make request
    response, err := client.Get("https://api.darksky.net/forecast/68a391b503f11aa6fa13d405bfefdaba/37.8267,-122.4233")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Copy data from the response to standard output
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Number of bytes copied to STDOUT:", n)
}
