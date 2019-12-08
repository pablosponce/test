package main
import (
        "fmt"
        "log"
        "net/http"
        "os"
)
func main() {
        port := "6969"
        if fromEnv := os.Getenv("PORT"); fromEnv != "" {
                port = fromEnv
        }
        server := http.NewServeMux()
        server.HandleFunc("/", hello)
        log.Printf("Server listening on port %s", port)
        err := http.ListenAndServe(":"+port, server)
        log.Fatal(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
        log.Printf("GET %s", r.URL.Path)
        fmt.Fprintf(w, "Hello world, it's Mariusz v1\n")
}
