package main

import (
    "net/http"
    "os"
    "syscall"

    "github.com/yourusername/yourproject/pkg/api"
)

func main() {
    err := api.StartServer()
    if err!= nil {
        // handle error
        fmt.Println("Error starting server:", err)
        os.Exit(1)
    }

    // wait for a signal to shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit

    // shutdown the server gracefully
    api.ShutdownServer()
}