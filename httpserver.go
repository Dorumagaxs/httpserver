package main

import (
    "fmt"
    "os"
    "log"
    "time"
    "net/http"
    "io/ioutil"
)

type RequestResolver struct {}

func (resolver RequestResolver) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    // fmt.Println("nada")
    fileName := fmt.Sprint(request.URL)[1:]

    fileContents, err := ioutil.ReadFile(fileName)

    if err == nil {
        response.Write(fileContents)
    } else {
        response.Write([]byte("No such file\n"))
    }
}

// httpserver directory port [certificate private_key]
func main() {
    if len(os.Args) < 3 {
        fmt.Printf("Not enough parameters!\nhttpserver directory port [certificate private_key]\n")
        os.Exit(0)
    }

    directory := os.Args[1]
    // directory := "/var/www/html/"
    port := ":"+os.Args[2]

    resolver := RequestResolver{}

    server := &http.Server{
        Addr: port,
        Handler: resolver,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    if len(os.Args) >= 4 {
        certificate := os.Args[3]
        privateKey := os.Args[4]

        fmt.Printf("Serving directory %s on https://localhost%s\n", directory, port)
        log.Fatal(server.ListenAndServeTLS(certificate, privateKey))
    } else {
        fmt.Printf("Serving directory %s on http://localhost%s\n", directory, port)
        log.Fatal(server.ListenAndServe())
    }
}
