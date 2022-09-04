package main

import (
    "strings"
    "log"
    "os"
    "os/exec"
    "net"
    "net/http"
    "fmt"
    "github.com/gorilla/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
    remote_host := strings.Split(r.RemoteAddr, ":")[0]
    if net.ParseIP(remote_host) == nil {
        http.Error(w, "huh", http.StatusForbidden)
    } else {
        exec_command := exec.Command("sudo", "firewall-cmd", "--zone=work", "--add-source=" + remote_host + "/32" )

        err := exec_command.Run()
        if err != nil {
            http.Error(w, "forbidden", http.StatusForbidden)
        }
	fmt.Fprintf(w, "Welcome to the world!!")
    }
}

func main() {
    route := os.Getenv("HANDLER_ROUTE")
    route = "/" + route
    http.HandleFunc(route, handler)
    log.Fatal(http.ListenAndServe(":25818", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
